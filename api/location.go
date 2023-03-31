package api

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/shoriwe/routes-service/models"
)

var LocationUpgrader = websocket.Upgrader{
	WriteBufferSize: 10240,
	ReadBufferSize:  10240,
}

type LocationHandshake struct {
	Token       string           `json:"token,omitempty"`
	VehicleUUID string           `json:"vehicleUUID,omitempty"`
	Location    *models.Location `json:"location,omitempty"`
	Status      *Status          `json:"status,omitempty"`
}

func (h *Handler) LocationProducer(ctx *gin.Context) {
	conn, upgradeErr := LocationUpgrader.Upgrade(ctx.Writer, ctx.Request, ctx.Writer.Header())
	if upgradeErr != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, Status{Succeed: false, Error: upgradeErr.Error()})
		return
	}
	defer conn.Close()
	var lHS LocationHandshake
	hErr := conn.ReadJSON(&lHS)
	if hErr != nil {
		return
	}
	ak, aErr := h.Controller.AuthorizeAPIKey(lHS.Token)
	if aErr != nil {
		conn.WriteJSON(LocationHandshake{Status: &Status{Succeed: false, Error: aErr.Error()}})
		return
	}
	w, lErr := h.Controller.LocationProducer(ak.VehicleUUID.String())
	if lErr != nil {
		conn.WriteJSON(LocationHandshake{Status: &Status{Succeed: false, Error: lErr.Error()}})
		return
	}
	conn.WriteJSON(LocationHandshake{Status: &StatusOK})
	for {
		rErr := conn.ReadJSON(&lHS)
		if rErr != nil {
			return
		}
		sErr := w.Send(lHS.Location)
		if sErr != nil {
			return
		}
	}
}

func (h *Handler) LocationListener(ctx *gin.Context) {
	conn, upgradeErr := LocationUpgrader.Upgrade(ctx.Writer, ctx.Request, ctx.Writer.Header())
	if upgradeErr != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, Status{Succeed: false, Error: upgradeErr.Error()})
		return
	}
	defer conn.Close()
	var lHS LocationHandshake
	hErr := conn.ReadJSON(&lHS)
	if hErr != nil {
		return
	}
	_, aErr := h.Controller.AuthorizeUser(lHS.Token)
	if aErr != nil {
		conn.WriteJSON(LocationHandshake{Status: &Status{Succeed: false, Error: aErr.Error()}})
		return
	}
	conn.WriteJSON(LocationHandshake{Status: &StatusOK})
	var channels []chan *models.Location
	var chMutex sync.Mutex
	go func() {
		for {
			rErr := conn.ReadJSON(&lHS)
			if rErr != nil {
				return
			}
			ch, lErr := h.Controller.LocationListener(lHS.VehicleUUID)
			if lErr != nil {
				continue
			}
			chMutex.Lock()
			channels = append(channels, ch)
			chMutex.Unlock()
			wErr := conn.WriteJSON(LocationHandshake{Status: &StatusOK})
			if wErr != nil {
				return
			}
		}
	}()
	tick := time.NewTicker(time.Millisecond)
	defer tick.Stop()
	for {
		if func() bool {
			chMutex.Lock()
			defer chMutex.Unlock()
			for _, ch := range channels {
				select {
				case l := <-ch:
					wErr := conn.WriteJSON(LocationHandshake{Location: l})
					if wErr != nil {
						return true
					}
				default:
					<-tick.C
				}
			}
			return false
		}() {
			break
		}
	}
}

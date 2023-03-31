package controller

import (
	"fmt"
	"sync"
	"sync/atomic"

	uuid "github.com/satori/go.uuid"
	"github.com/shoriwe/routes-service/models"
)

type Worker struct {
	Controller      *Controller
	VehicleUUID     uuid.UUID
	CurrentListener atomic.Int64
	Listeners       sync.Map
}

func (w *Worker) closeListener(l chan *models.Location) {
	defer func() {
		recover()
	}()
	close(l)
}

func (w *Worker) Close() {
	w.Listeners.Range(func(key, value any) bool {
		ch := value.(chan *models.Location)
		w.closeListener(ch)
		return true
	})
	w.Controller.Workers.Delete(w.VehicleUUID)
}

func (w *Worker) sendListener(lId int64, l chan *models.Location, location *models.Location) {
	defer func() {
		r := recover()
		if _, ok := r.(error); ok {
			w.Listeners.Delete(lId)
		}
	}()
	l <- location
}

func (w *Worker) Send(location *models.Location) error {
	location.VehicleUUID = w.VehicleUUID
	cErr := w.Controller.DB.Create(location).Error
	if cErr != nil {
		return cErr
	}
	w.Listeners.Range(func(key, value any) bool {
		w.sendListener(key.(int64), value.(chan *models.Location), location)
		return true
	})
	return nil
}

func (w *Worker) Listen() (chan *models.Location, error) {
	listenerId := w.CurrentListener.Add(1)
	ch := make(chan *models.Location, 1024)
	w.Listeners.Store(listenerId, ch)
	return ch, nil
}

func (c *Controller) LocationProducer(vehicleUUID string) (*Worker, error) {
	var vehicle models.Vehicle
	fErr := c.DB.Where("uuid = ?", vehicleUUID).First(&vehicle).Error
	if fErr != nil {
		return nil, fErr
	}
	worker, found := c.Workers.LoadOrStore(vehicle.UUID, &Worker{
		Controller:  c,
		VehicleUUID: vehicle.UUID,
	})
	if found {
		return nil, fmt.Errorf("producer already running")
	}
	return worker.(*Worker), nil
}

func (c *Controller) LocationListener(vehicleUUID string) (chan *models.Location, error) {
	var vehicle models.Vehicle
	fErr := c.DB.Where("uuid = ?", vehicleUUID).First(&vehicle).Error
	if fErr != nil {
		return nil, fErr
	}
	worker, found := c.Workers.Load(vehicle.UUID)
	if !found {
		return nil, fmt.Errorf("no producer found")
	}
	return worker.(*Worker).Listen()
}

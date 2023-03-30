package api

import (
	"net/http"
	"testing"
)

func TestHandler_Map(t *testing.T) {
	c, expect, serverClose := NewTest(t)
	defer serverClose()
	obj := expect.GET(MapRoute).Expect().
		Status(http.StatusOK).JSON().Object()
	obj.Value("nodes").Array().Value(0).Object().Value("latitude").Number().
		IsEqual(c.Map.Nodes[0].Latitude)
}

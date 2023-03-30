package controller

import "fmt"

const DefaultPageSize = 20

var ErrorUnauthorized = fmt.Errorf("unauthorized")

type Result[T any] struct {
	Page       int64 `json:"page"`
	TotalPages int64 `json:"totalPages"`
	Results    []T
}

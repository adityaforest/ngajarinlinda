package handlers

import (
	"context"
	"ngajarinlinda"
	"ngajarinlinda/gen/models"
	"ngajarinlinda/gen/restapi/operations/mantan"
)

type (
	Handler interface {
		MantanHandler
	}

	MantanHandler interface {
		Mantan(ctx context.Context, rt *ngajarinlinda.Runtime, params mantan.GetAllMantanParams) (mantanList []*models.Mantan, err error)
	}
)

func NewHandler() Handler {
	return &handler{}
}

type handler struct{}

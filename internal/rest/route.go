package rest

import (
	"context"
	"ngajarinlinda"
	"ngajarinlinda/gen/models"
	"ngajarinlinda/gen/restapi/operations"
	"ngajarinlinda/gen/restapi/operations/mantan"
	"ngajarinlinda/internal/handlers"

	"github.com/go-openapi/runtime/middleware"
)

func Route(rt *ngajarinlinda.Runtime, api *operations.NgajarinLindaAPI, apiHandler handlers.Handler) {

	api.MantanGetAllMantanHandler = mantan.GetAllMantanHandlerFunc(func(params mantan.GetAllMantanParams) middleware.Responder {
		var mantanLIst []*models.Mantan

		mantanLIst, err := apiHandler.Mantan(context.Background(), rt, params)

		if err != nil {
			errResponse := rt.GetError(err)
			return mantan.NewGetAllMantanInternalServerError().WithPayload(&mantan.GetAllMantanInternalServerErrorBody{
				Message: errResponse.Error(),
			})
		}

		return mantan.NewGetAllMantanOK().WithPayload(mantanLIst)
	})

}

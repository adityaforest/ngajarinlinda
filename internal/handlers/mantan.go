package handlers

import (
	"context"
	"ngajarinlinda"
	"ngajarinlinda/gen/models"
	"ngajarinlinda/gen/restapi/operations/mantan"

	"net/http"

	"gorm.io/gorm"
)

func (h *handler) Mantan(ctx context.Context, rt *ngajarinlinda.Runtime, params mantan.GetAllMantanParams) ([]*models.Mantan, error) {

	var mantanList []*models.Mantan

	rt.Db.Raw("SELECT * FROM mantans").Scan(&mantanList)

	err := rt.Db.Raw("SELECT * FROM mantans").Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = rt.SetError(http.StatusNotFound, "No record found in the database")
		}
	}

	return mantanList, err
}

package handlers

import (
	"errors"
	"net/http"

	"github.com/alamgir-ahosain/go-gorm-postgres-crud-rest-api/cmd/MyApp/internal/db"
	"github.com/alamgir-ahosain/go-gorm-postgres-crud-rest-api/cmd/MyApp/internal/models"
	"github.com/alamgir-ahosain/go-gorm-postgres-crud-rest-api/cmd/MyApp/internal/services"
	"gorm.io/gorm"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := services.GetID(w, r) //get id from the url
	if id == 0 {
		return
	}

	var user models.Users
	row := db.DB.First(&user, id)

	//Handle error
	if row.Error != nil {
		if errors.Is(row.Error, gorm.ErrRecordNotFound) {
			services.HandleHTTPError(w, row.Error, http.StatusNotFound)
		} else {
			services.HandleHTTPError(w, row.Error, http.StatusInternalServerError)
		}
		return
	}

	//delete query
	row = db.DB.Delete(&user, id)
	services.HandleHTTPError(w, row.Error, http.StatusInternalServerError)
	if row.RowsAffected == 0 {
		services.HandleHTTPError(w, row.Error, http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent) //204,no content

}

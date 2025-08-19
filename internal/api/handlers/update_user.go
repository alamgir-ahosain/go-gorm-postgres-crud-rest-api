package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/alamgir-ahosain/go-gorm-postgres-crud-rest-api/cmd/MyApp/internal/db"
	"github.com/alamgir-ahosain/go-gorm-postgres-crud-rest-api/cmd/MyApp/internal/models"
	"github.com/alamgir-ahosain/go-gorm-postgres-crud-rest-api/cmd/MyApp/internal/services"
	"gorm.io/gorm"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {

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

	//decode request body into user struct
	var updateUser struct {
		SID  string  `gorm:"column:sid" json:"sid"`
		Name string  `json:"name"`
		CGPA float32 `json:"cgpa"`
	}

	err := json.NewDecoder(r.Body).Decode(&updateUser)
	services.HandleHTTPError(w, err, http.StatusBadRequest)

	// update Query
	db.DB.Model(&user).Updates(models.Users{
		SID:  updateUser.SID,
		Name: updateUser.Name,
		CGPA: updateUser.CGPA,
	})
	// Reload updated user
	db.DB.First(&user, id)
	services.MakeJSONFormatFunc(w, user, 200)

}

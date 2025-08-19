package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/alamgir-ahosain/go-gorm-postgres-crud-rest-api/cmd/MyApp/internal/db"
	"github.com/alamgir-ahosain/go-gorm-postgres-crud-rest-api/cmd/MyApp/internal/models"
	"github.com/alamgir-ahosain/go-gorm-postgres-crud-rest-api/cmd/MyApp/internal/services"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	//Decode json body into users
	var user models.Users
	err := json.NewDecoder(r.Body).Decode(&user)
	services.HandleHTTPError(w, err, http.StatusBadRequest)

	// Execute Query
	row := db.DB.Create(&user) //insert into users(sid,name,cgpa) values(...)
	services.HandleHTTPError(w, row.Error, http.StatusInternalServerError)
	services.MakeJSONFormatFunc(w, user, 201) // Send JSON response to client

}

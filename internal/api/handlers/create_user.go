package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/alamgir-ahosain/go-gorm-postgres-crud-rest-api/cmd/MyApp/internal/models"
	"github.com/alamgir-ahosain/go-gorm-postgres-crud-rest-api/cmd/MyApp/internal/services"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	//Decode json body into users
	var user models.Users
	err := json.NewDecoder(r.Body).Decode(&user)
	services.HandleHTTPError(w, err, http.StatusBadRequest)

	//insert user and get id
	//var id int
	// err = db.DB.QueryRow("insert into users(sid,name,cgpa) values($1,$2,$3) returning id", user.SID, user.Name, user.CGPA).Scan(&id)
	// services.HandleHTTPError(w, err, http.StatusInternalServerError)
	// user.ID = id
	// services.MakeJSONFormatFunc(w, user, 201) // Send JSON response to client

}

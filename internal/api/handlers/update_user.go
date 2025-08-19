package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/alamgir-ahosain/go-gorm-postgres-crud-rest-api/cmd/MyApp/internal/models"
	"github.com/alamgir-ahosain/go-gorm-postgres-crud-rest-api/cmd/MyApp/internal/services"
	// "github.com/alamgir-ahosain/go-gorm-postgres-crud-rest-api/internal/db"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	id := services.GetID(w, r) //get id from the url
	if id == 0 {
		return
	}
	//decode request body into user struct
	var user models.Users
	err := json.NewDecoder(r.Body).Decode(&user)
	services.HandleHTTPError(w, err, http.StatusBadRequest)

	//Run update query
	// _, err = db.DB.Exec("update users set sid=$1,name=$2,cgpa=$3 where id=$4", user.SID, user.Name, user.CGPA, id)
	// services.HandleHTTPError(w, err, http.StatusInternalServerError)
	// user.ID = id
	// services.MakeJSONFormatFunc(w, user, 200)

}

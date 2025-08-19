package getusers

import (
	"net/http"

	"github.com/alamgir-ahosain/go-gorm-postgres-crud-rest-api/cmd/MyApp/internal/db"
	"github.com/alamgir-ahosain/go-gorm-postgres-crud-rest-api/cmd/MyApp/internal/models"
	"github.com/alamgir-ahosain/go-gorm-postgres-crud-rest-api/cmd/MyApp/internal/services"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.Users
	row := db.DB.Find(&users) //SELECT * FROM users;
	services.HandleHTTPError(w, row.Error, http.StatusInternalServerError)
	services.MakeJSONFormatFunc(w, users, 200) // Convert the slice of users into JSON and send as HTTP response

}

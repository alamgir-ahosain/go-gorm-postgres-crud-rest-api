package getusers

import (
	"errors"
	"net/http"

	"github.com/alamgir-ahosain/go-gorm-postgres-crud-rest-api/cmd/MyApp/internal/db"
	"github.com/alamgir-ahosain/go-gorm-postgres-crud-rest-api/cmd/MyApp/internal/models"
	"github.com/alamgir-ahosain/go-gorm-postgres-crud-rest-api/cmd/MyApp/internal/services"
	"gorm.io/gorm"
)

func GetUsersByID(w http.ResponseWriter, r *http.Request) {
	id := services.GetID(w, r) //get id from the url
	if id == 0 {
		return
	}
	var user models.Users
	row := db.DB.First(&user, id) //select * from users where id=? order by id limit 1

	//Handle error
	if row.Error != nil {
		if errors.Is(row.Error, gorm.ErrRecordNotFound) {
			services.HandleHTTPError(w, row.Error, http.StatusNotFound)
		} else {
			services.HandleHTTPError(w, row.Error, http.StatusInternalServerError)
		}
		return
	}
	services.HandleHTTPError(w, row.Error, http.StatusInternalServerError)
	services.MakeJSONFormatFunc(w, user, 200)
}

package handlers

import (
	"net/http"

	"github.com/alamgir-ahosain/go-gorm-postgres-crud-rest-api/cmd/MyApp/internal/services"
	// "github.com/alamgir-ahosain/go-gorm-postgres-crud-rest-api/internal/db"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := services.GetID(w, r) //get id from the url
	if id == 0 {
		return
	}

	//delete query
	// row, err := db.DB.Exec("delete from users where id=$1", id)
	// services.HandleHTTPError(w, err, http.StatusInternalServerError)

	// //check if any row was deleted
	// rowsAffected, err := row.RowsAffected()
	// services.HandleHTTPError(w, err, http.StatusInternalServerError)
	// if rowsAffected == 0 {
	// 	http.Error(w, "User not found", http.StatusNotFound)
	// 	return
	// }
	// w.WriteHeader(http.StatusNoContent) //204

}

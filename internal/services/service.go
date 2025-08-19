package services

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func HandleHTTPError(w http.ResponseWriter, err error, code int) {
	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}
}

func MakeJSONFormatFunc(w http.ResponseWriter, data interface{}, code int) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}

func GetID(w http.ResponseWriter, r *http.Request) int {
	ids := mux.Vars(r)
	idStr := ids["id"] // map[string]string{"id":"5"}
	id, err := strconv.Atoi(idStr)
	HandleHTTPError(w, err, http.StatusBadRequest)
	return id
}

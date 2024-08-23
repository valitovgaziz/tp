package admin

import (
	"api/src/models"
	"api/src/storages/psql"
	"encoding/json"
	"net/http"
)

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	qr := psql.PSQL_GORM_DB.Find(&users)
	if qr.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	jsData, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	w.Write([]byte(jsData))

}
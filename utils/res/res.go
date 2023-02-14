package res

import (
	"encoding/json"
	"net/http"
)

func JSON(rw http.ResponseWriter, status int, data interface{}) error {
	rw.WriteHeader(status)
	rw.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(rw).Encode(data)
}


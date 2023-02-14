package utils

import (
	"encoding/json"
	"net/http"
)

func Decode(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}
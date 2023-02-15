package res

import (
	"encoding/json"
	"github.com/sant470/neobank/utils/errors"
	"net/http"
)

func JSON(rw http.ResponseWriter, status int, data interface{}) *errors.AppError{
	rw.WriteHeader(status)
	rw.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(rw).Encode(data); err != nil {
		return errors.InternalServerError(err.Error())
	}
	return nil 
}


package v1

import (
	"log"
	"net/http"
	"github.com/sant470/neobank/models"
	"github.com/sant470/neobank/utils"
	"github.com/sant470/neobank/utils/errors"
	"github.com/sant470/neobank/utils/res"
)

type AccountHandler struct {
	lgr *log.Logger
}

func NewAccountHandler(l *log.Logger) *AccountHandler {
	return &AccountHandler{lgr: l}
}

func (ah *AccountHandler) CreateAccount(rw http.ResponseWriter, r *http.Request) *errors.AppError{
	var account models.Account
	if err := utils.Decode(r, &account); err != nil {
		return errors.BadRequest(err.Error())
	}
	return res.JSON(rw, http.StatusCreated, account)
}


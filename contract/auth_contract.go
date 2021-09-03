package contract

import (
	"encoding/json"
	"net/http"

	"rakamin4/common/errors"
	"rakamin4/common/util"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator"
	log "github.com/sirupsen/logrus"
)

type JWTMapClaim struct {
	Authorized bool   `json:"authorized"`
	RequestID  string `json:"requestID"`
	jwt.StandardClaims
}

type GetTokenResponseContract struct {
	Token string `json:"token"`
}

type ValidateTokenRequestContract struct {
	Token string `json:"token"`
}

func NewValidateTokenRequest(r *http.Request) (*ValidateTokenRequestContract, error) {
	validateTokenContract := new(ValidateTokenRequestContract)
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(validateTokenContract); err != nil {
		log.Warning(err)
		return nil, errors.NewBadRequestError(err)
	}

	validate := validator.New()
	util.UseJsonFieldValidation(validate)

	if err := validate.Struct(validate); err != nil {
		log.Error(err)
		return nil, errors.NewValidationError(errors.ValidateErrToMapString(err.(validator.ValidationErrors)))
	}

	return validateTokenContract, nil
}

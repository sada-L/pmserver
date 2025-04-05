package utils

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/sada-L/pmserver/internal/model"
)

// M is a generic map
type M map[string]interface{}

func ValidationError(w http.ResponseWriter, _err error) {
	resp := model.ErrorM{}

	switch err := _err.(type) {
	case validator.ValidationErrors:
		for _, e := range err {
			field := e.Field()
			msg := CheckTagRules(e)
			resp[field] = append(resp[field], msg)
		}
	default:
		resp["non_field_error"] = append(resp["non_field_error"], err.Error())
	}
	ErrorResponse(w, http.StatusUnprocessableEntity, resp)
}

func BadRequestError(w http.ResponseWriter) {
	msg := "unable to process request"
	ErrorResponse(w, http.StatusUnprocessableEntity, msg)
}

func InvalidUserCredentialsError(w http.ResponseWriter) {
	msg := "invalid authentication credentials"
	ErrorResponse(w, http.StatusUnauthorized, msg)
}

func InvalidAuthTokenError(w http.ResponseWriter) {
	w.Header().Set("WWW-Authenticate", "Token")
	msg := "invalid or missing authentication token"
	ErrorResponse(w, http.StatusUnauthorized, msg)
}

func NotFoundError(w http.ResponseWriter, err model.ErrorM) {
	ErrorResponse(w, http.StatusNotFound, err)
}

func ServerError(w http.ResponseWriter, err error) {
	log.Println(err)
	ErrorResponse(w, http.StatusInternalServerError, "internal error")
}

func ErrorResponse(w http.ResponseWriter, code int, errs interface{}) {
	WriteJSON(w, code, M{"errors": errs})
}

func CheckTagRules(e validator.FieldError) (errMsg string) {
	tag, field, param, value := e.ActualTag(), e.Field(), e.Param(), e.Value()

	if tag == "required" {
		errMsg = "this field is required"
	}

	if tag == "email" {
		errMsg = fmt.Sprintf("%q is not a valid email", value)
	}

	if tag == "min" {
		errMsg = fmt.Sprintf("%s must be greater than %v", field, param)
	}

	if tag == "max" {
		errMsg = fmt.Sprintf("%s must be less than %v", field, param)
	}
	return
}

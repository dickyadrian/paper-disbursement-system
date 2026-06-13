package server

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var validate = validator.New()

func bindAndValidateRequest[T any](c echo.Context) (T, error) {
	var req T
	if err := c.Bind(&req); err != nil {
		return req, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := validate.Struct(req); err != nil {
		return req, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return req, nil
}

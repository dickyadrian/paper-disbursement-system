package server

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

type CreateUserRequest struct {
	Name    string `json:"name" validate:"required,max=255"`
	Balance int64  `json:"balance" validate:"min=0"`
}

func (h *Handler) CreateUser(c echo.Context) error {
	ctx := c.Request().Context()

	req, err := bindAndValidateRequest[CreateUserRequest](c)
	if err != nil {
		return err
	}

	user, err := h.App.UserRepository.Create(ctx, req.Name, req.Balance)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, user)
}

func (h *Handler) GetUser(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}

	user, err := h.App.UserRepository.Get(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return echo.NewHTTPError(http.StatusNotFound, "user not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

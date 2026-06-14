package server

import (
	"errors"
	"net/http"

	"github.com/dickyadrian/paper-disbursement-system/internal/task"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

type CreateDisbursementRequest struct {
	UserID int64 `json:"user_id" validate:"required"`
	Amount int64 `json:"amount" validate:"required,gt=0"`
}

func (h *Handler) CreateDisbursement(c echo.Context) error {
	ctx := c.Request().Context()

	req, err := bindAndValidateRequest[CreateDisbursementRequest](c)
	if err != nil {
		return err
	}

	user, err := h.App.UserRepository.Get(ctx, req.UserID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return echo.NewHTTPError(http.StatusNotFound, "user not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if user.Balance < req.Amount {
		return echo.NewHTTPError(http.StatusBadRequest, "insufficient balance")
	}

	disbursement, err := h.App.DisbursementRepository.Create(ctx, req.UserID, req.Amount)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	t, err := task.NewDisbursementProcessTask(disbursement.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if _, err := h.App.AsynqClient.Enqueue(t); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, disbursement)
}

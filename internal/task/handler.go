package task

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/dickyadrian/paper-disbursement-system/application"
	"github.com/hibiken/asynq"
)

type Handler struct {
	*application.App
}

func NewHandler(app *application.App) *Handler {
	return &Handler{app}
}

func (h *Handler) ProcessDisbursement(ctx context.Context, t *asynq.Task) error {
	var payload DisbursementProcessPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return fmt.Errorf("unmarshal payload: %w", err)
	}

	return h.App.DisbursementRepository.Process(ctx, payload.DisbursementID)
}

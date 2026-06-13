package task

import (
	"encoding/json"

	"github.com/hibiken/asynq"
)

const TypeDisbursementProcess = "disbursement:process"

type DisbursementProcessPayload struct {
	DisbursementID int64 `json:"disbursement_id"`
}

func NewDisbursementProcessTask(disbursementID int64) (*asynq.Task, error) {
	payload, err := json.Marshal(DisbursementProcessPayload{DisbursementID: disbursementID})
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TypeDisbursementProcess, payload), nil
}

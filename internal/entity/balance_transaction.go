package entity

import "time"

type BalanceTransaction struct {
	ID             int64     `json:"id"`
	UserID         int64     `json:"user_id"`
	DisbursementID int64     `json:"disbursement_id"`
	Amount         int64     `json:"amount"`
	BalanceAfter   int64     `json:"balance_after"`
	CreatedAt      time.Time `json:"created_at"`
}

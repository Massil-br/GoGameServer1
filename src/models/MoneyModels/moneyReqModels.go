package moneymodels

import "github.com/google/uuid"

type MoneyRequest struct {
	Amount float64 `json:"amount"`
	UserId *uuid.UUID `json:"userId"`
}

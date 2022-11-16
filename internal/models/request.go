package models

type UserRequest struct {
	UserId        string  `json:"id" validate:"required"`
	AmountOfMoney float64 `json:"amount_of_money" validate:"required"`
	OperationCode string  `json:"operation_code" validate:"required"`
}

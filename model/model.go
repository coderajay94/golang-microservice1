package model

type UserRequest struct {
	Email string `json:"email"`
}

type UserResponse struct {
	Email         string  `json:"email"`
	Name          string  `json:"name"`
	Balance       float64 `json:"balance"`
	AccountNumber string  `json:"account_number"`
}
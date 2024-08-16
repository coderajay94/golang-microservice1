package model

// type UserRequest struct {
// 	Email string `json:"email"`
// }

// type UserResponse struct {
// 	Email         string  `json:"email"`
// 	Name          string  `json:"name"`
// 	Balance       float64 `json:"balance"`
// 	AccountNumber string  `json:"accountNumber"`
// }

type UserRequestDB struct {
	Email string `bson:"_id,omitempty" json:"email"`
}

type UserResponseDB struct {
	Email         string  `bson:"_id,omitempty" json:"email"`
	Name          string  `bson:"name" json:"name"`
	Balance       float64 `bson:"balance" json:"balance"`
	AccountNumber string  `bson:"accountNumber" json:"accountNumber"`
}
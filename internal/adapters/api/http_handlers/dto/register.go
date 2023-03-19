package dto

type RegisterInput struct {
	Login       string `json:"login"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

type RegisterPayload struct {
	Id          int    `json:"id"`
	Login       string `json:"login"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Token       string `json:"token"`
}

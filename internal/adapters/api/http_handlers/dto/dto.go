package dto

type RegisterInput struct {
	Login       string `json:"login"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

type RegisterPayload struct {
}

type LoginInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginPayload struct {
}

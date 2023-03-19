package dto

type LoginInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginPayload struct {
	Token string `json:"token"`
}

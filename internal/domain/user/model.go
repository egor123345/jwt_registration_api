package user

type User struct {
	Id          int    `json:"id,omitempty"`
	Login       string `json:"login"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

package dto

type RespCustomer struct {
	ID          int    `json:"id"`
	FullName    string `json:"full_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type RespCustomerLogin struct {
	ID          int    `json:"id"`
	FullName    string `json:"full_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Token       string `json:"token"`
}

type ReqCustomer struct {
	FullName    string `json:"full_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

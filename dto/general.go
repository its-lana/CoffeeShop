package dto

type LoginRequest struct {
	Email   string `json:"email"`
	Passwod string `json:"password"`
}

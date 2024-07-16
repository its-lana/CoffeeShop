package dto

type RespMerchant struct {
	ID           int    `json:"id"`
	MerchantName string `json:"merchant_name"`
	Address      string `json:"address"`
	PICName      string `json:"pic_name"`
	Email        string `json:"email"`
	PhoneNumber  string `json:"phone_number"`
	IsOpen       bool   `json:"is_open"`
}

type RespMerchantLogin struct {
	ID           int    `json:"id"`
	MerchantName string `json:"merchant_name"`
	Address      string `json:"address"`
	PICName      string `json:"pic_name"`
	Email        string `json:"email"`
	PhoneNumber  string `json:"phone_number"`
	IsOpen       bool   `json:"is_open"`
	Token        string `json:"token"`
}

type ReqMerchant struct {
	MerchantName string `json:"merchant_name"`
	Address      string `json:"address"`
	PICName      string `json:"pic_name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	PhoneNumber  string `json:"phone_number"`
	IsOpen       bool   `json:"is_open"`
}

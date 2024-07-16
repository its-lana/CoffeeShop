package dto

type ReqCategory struct {
	CategoryName string `json:"category_name"`
	MerchantID   int    `json:"merchant_id"`
}

type RespCategory struct {
	ID           int        `json:"id"`
	CategoryName string     `json:"category_name"`
	MerchantID   int        `json:"merchant_id"`
	Menus        []RespMenu `json:"menus"`
}

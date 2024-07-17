package dto

type ReqMenu struct {
	ProductName        string `json:"product_name"`
	Price              int    `json:"price"`
	Description        string `json:"description"`
	ProductCode        string `json:"product_code"`
	ProductImage       string `json:"product_image"`
	CategoryID         int    `json:"category_id"`
	AvailabilityStatus bool   `json:"availability_status"`
}

type RespMenu struct {
	ID                 int    `json:"id"`
	ProductName        string `json:"product_name"`
	Price              int    `json:"price"`
	Description        string `json:"description"`
	ProductCode        string `json:"product_code"`
	ProductImage       string `json:"product_image"`
	CategoryID         int    `json:"category_id"`
	AvailabilityStatus bool   `json:"availability_status"`
}

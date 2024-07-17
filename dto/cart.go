package dto

type ReqCart struct {
	CustomerID int `json:"customer_id"`
	MerchantID int `json:"merchant_id"`
}

type RespCart struct {
	ID         int             `json:"id"`
	CustomerID int             `json:"customer_id"`
	MerchantID int             `json:"merchant_id"`
	OrderItems []RespOrderItem `json:"order_items"`
}

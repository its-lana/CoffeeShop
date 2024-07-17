package dto

type ReqOrderItem struct {
	MenuID    int    `json:"menu_id"`
	Quantity  int    `json:"quantity"`
	OwnerID   int    `json:"owner_id"`
	OwnerType string `json:"owner_type"`
}

type RespOrderItem struct {
	ID           int       `json:"id"`
	MenuID       int       `json:"menu_id"`
	Quantity     int       `json:"quantity"`
	OwnerID      int       `json:"owner_id"`
	OwnerType    string    `json:"owner_type"`
	MerchantMenu *RespMenu `json:"merchant_menu"`
}

type Owner struct {
	OwnerID   int    `json:"owner_id"`
	OwnerType string `json:"owner_type"`
}

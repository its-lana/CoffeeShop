package dto

import (
	"time"
)

type ReqOrder struct {
	OrderUID      string    `json:"order_uid"`
	CustomerID    int       `json:"customer_id"`
	MerchantID    int       `json:"merchant_id"`
	FinalAmount   int       `json:"final_amount"`
	OrderType     string    `json:"order_type"`
	OrderNotes    string    `json:"order_notes"`
	OrderStatus   string    `json:"order_status"`
	NoteStatus    string    `json:"note_status"`
	OrderCode     string    `json:"order_code"`
	PaymentMethod string    `json:"payment_method"`
	OrderDate     time.Time `json:"order_date"`
}

type RespOrder struct {
	ID          int             `json:"id"`
	OrderUID    string          `json:"order_uid"`
	CustomerID  int             `json:"customer_id"`
	MerchantID  int             `json:"merchant_id"`
	FinalAmount int             `json:"final_amount"`
	OrderType   string          `json:"order_type"`
	OrderNotes  string          `json:"order_notes,omitempty"`
	OrderStatus string          `json:"order_status"`
	NoteStatus  string          `json:"note_status,omitempty"`
	OrderCode   string          `json:"order_code"`
	OrderDate   time.Time       `json:"order_date"`
	OrderItems  []RespOrderItem `json:"order_items"`
	Payment     RespPayment     `json:"payment"`
}

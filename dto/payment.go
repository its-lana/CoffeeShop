package dto

import "time"

type ReqPayment struct {
	PaymentMethod string `json:"payment_method"`
	OrderUID      string `json:"order_uid"`
	CustomerID    int    `json:"customer_id"`
	FinalAmount   int    `json:"final_amount"`
	PaymentURL    string `json:"payment_url"`
}

type RespPayment struct {
	ID            int       `json:"id"`
	CustomerID    int       `json:"customer_id"`
	OrderUID      string    `json:"order_uid"`
	PaymentAmount int       `json:"payment_amount"`
	Status        string    `json:"status"`
	PaymentMethod string    `json:"payment_method"`
	PaymentURL    string    `json:"payment_url"`
	PaidDate      time.Time `json:"paid_date"`
}

type ReqTransactionDetail struct {
	OrderUID      string `json:"order_uid"`
	GrossAmount   int    `json:"gross_amount"`
	FullName      string `json:"full_name"`
	Email         string `json:"email"`
	PaymentMethod string `json:"payment_method"`
}

type ReqTransactionNotification struct {
	TransactionStatus string `json:"transaction_status"`
	OrderID           string `json:"order_id"`
	PaymentType       string `json:"payment_type"`
	FraudStatus       string `json:"fraud_status"`
	EstimateTime      string `json:"estimate_time"`
}

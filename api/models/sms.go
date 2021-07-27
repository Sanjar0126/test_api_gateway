package models

type SendSmsPaymentsModel struct {
	Phone       string `json:"phone"`
	OrderID     string `json:"order_id"`
	PaymentType string `json:"payment_type"`
}

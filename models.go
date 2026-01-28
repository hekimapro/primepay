package primepay

import (
	"encoding/json"
	"time"
)

type CollectionRequestPayload struct {
	Action               string  `json:"action"`
	Amount               float64 `json:"amount"`
	CustomerPhoneNumber  string  `json:"msisdn"`
	Reference            string  `json:"reference"`
	CustomerName         string  `json:"buyer_name"`
	CustomerEmailAddress string  `json:"buyer_email"`
	Currency             string  `json:"currency"`
	CallbackURL          string  `json:"callback_url"`
}

type CollectionResponsePayload struct {
	Status           string          `json:"status"`
	TransactionID    string          `json:"transaction_id"`
	PaymentStatus    string          `json:"payment_status"`
	ProviderResponse json.RawMessage `json:"provider_response"`
}

type DisbursementRequestPayload struct {
	Action              string  `json:"action"`
	Amount              float64 `json:"amount"`
	CustomerPhoneNumber string  `json:"msisdn"`
	CallbackURL         string  `json:"callback_url"`
	Channel             string  `json:"channel"`
	Reference           string  `json:"reference"`
}

type DisbursementResponsePayload struct {
	Status           string          `json:"status"`
	TransactionID    string          `json:"transaction_id"`
	PaymentStatus    string          `json:"payment_status"`
	ProviderResponse json.RawMessage `json:"provider_response"`
}

type statusRequestPayload struct {
	Reference string `json:"reference"`
}

type StatusResponsePayload struct {
	Currency      string    `json:"currency"`
	Amount        float64   `json:"amount"`
	Status        string    `json:"status"`
	Reference     string    `json:"reference"`
	PaymentStatus string    `json:"payment_status"`
	CreatedTime   time.Time `json:"created_time"`
}

type CallbackResponsePayload struct {
	TranasctionID       string  `json:"transaction_id"`
	Reference           string  `json:"external_ref"`
	Currency            string  `json:"currency"`
	Amount              float64 `json:"amount"`
	CustomerPhoneNumber string  `json:"msisdn"`
	Status              string  `json:"status"`
	ProviderOrderID     string  `json:"provider_order_id"`
	ResultCode          string  `json:"result_code"`
}

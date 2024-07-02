package models

type Transaction struct {
    ID        int     `json:"id"`
    AccountID string  `json:"account_id"`
    Amount    float64 `json:"amount"`
    Status    string  `json:"status,omitempty"` // omitempty ensures status is excluded if empty
}

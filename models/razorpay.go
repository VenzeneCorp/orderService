package models

type RazorpayStatus int

const (
	Started RazorpayStatus = iota
	Success
	Failed
)

type RazorpayOrder struct {
	ID             string         `json:"id"`
	Amount         int            `json:"amount"`
	Receipt        string         `json:"receipt"`
	OrderID        uint64         `json:"order_id"`
	RazorpayStatus RazorpayStatus `json:"Razorpay_status"`
	CreatedAt      int64          `json:"created_at"`
	UpdatedAt      int64          `json:"updated_at"`
}

type RazorpayPaymentRespone struct {
	RazorpayOrderID string `json:"razorpay_order_id"`
	RAZORPAY_KEY_ID string `json:"razorpay_key_id"`
	Amount          int    `json:"amount"`
}

type RazorpaySuccessRequest struct {
	RazorpayOrderID   string         `json:"razorpay_order_id"`
	RazorpayPaymentID string         `json:"razorpay_payment_id"`
	RazorpaySignature string         `json:"razorpay_signature"`
	Status            RazorpayStatus `json:"status"`
}

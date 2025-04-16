package models

type Subscribed struct {
	ID          int    `json:"id"` // incremental id
	OrderID     string `json:"order_id"`
	UserID      int    `json:"user_id"`
	VendorID    int    `json:"vendor_id"`
	Days        int    `json:"days"`     // number of days
	RollOver    int    `json:"rollover"` // number of days
	BreakfastID string `json:"breakfast_id"`
	LunchID     string `json:"lunch_id"`
	DinnerID    string `json:"dinner_id"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

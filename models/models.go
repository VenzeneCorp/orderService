package models

type OrderType int

const (
	LiveOrder OrderType = iota
	SubscriptionOrder
)

type Status int

const (
	Pending Status = iota + 1
	Active
	Completed
	Cancelled
	Failed
	Refunded
	PartialRefunded
)

type Orders struct {
	ID          string    `json:"id"`
	UserID      int       `json:"user_id"`
	VendorID    int       `json:"vendor_id"`
	VendorName  string    `json:"vendor_name"`
	Amount      int       `json:"amount"`
	Discount    int       `json:"discount"`
	FinalAmount int       `json:"final_amount"`
	OrderType   OrderType `json:"order_type"`
	Status      string    `json:"status"`
	CreatedAt   string    `json:"created_at"`
}

// live and schedule orders here
type ItemHistory struct {
	ID          int    `json:"id"` // incremental id
	OrderID     string `json:"order_id"`
	MealID      int    `json:"meal_id"`
	MealName    string `json:"meal_name"`
	Quantity    int    `json:"quantity"`
	Veg         bool   `json:"veg"`
	Price       int    `json:"price"`
	DeliveredAt int64  `json:"delivered_at"`
}

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
	ID          uint64    `json:"id"`
	UserID      int       `json:"user_id"`
	VendorID    int       `json:"vendor_id"`
	VendorName  string    `json:"vendor_name"`
	Amount      int       `json:"amount"`
	Discount    int       `json:"discount"`
	FinalAmount int       `json:"final_amount"`
	OrderType   OrderType `json:"order_type"`
	Status      Status    `json:"status"`
	CreatedAt   int64     `json:"created_at"`
}

// live and schedule orders here
type ItemOrdered struct {
	ID          uint64 `json:"id"` // incremental id
	OrderID     uint64 `json:"order_id"`
	MealID      string `json:"meal_id"`
	MealName    string `json:"meal_name"`
	Quantity    int    `json:"quantity"`
	Veg         bool   `json:"veg"`
	Price       int    `json:"price"`
	DeliveredAt int64  `json:"delivered_at"`
}

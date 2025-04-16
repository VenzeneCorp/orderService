package models

type OrderType int

const (
	LiveOrder OrderType = iota
	SubscriptionOrder
)

type Status int

const (
	Pending Status = iota
	Active
	Completed
)

type Orders struct {
	ID          string    `json:"id"`
	UserID      int       `json:"user_id"`
	VendorID    int       `json:"vendor_id"`
	VendorName  string    `json:"vendor_name"`
	Amount      int       `json:"amount"`
	OrderType   OrderType `json:"order_type"`
	Status      string    `json:"status"`
	DeliveredAt int64     `json:"delivered_at"`
	CreatedAt   string    `json:"created_at"`
}

// live and schedule orders here
type ItemHistory struct {
	ID       int    `json:"id"`
	OrderID  string `json:"order_id"`
	MealID   int    `json:"meal_id"`
	MealName string `json:"meal_name"`
	Quantity int    `json:"quantity"`
}

package models

type OrderType int

const (
	LiveOrder OrderType = iota
	SubscriptionOrder
)

// store in NoSQL, just fetch and show the data
type Order struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	Items     []Item `json:"items"`
	Status    string `json:"status"`
	CreatedAt int64  `json:"created_at"`
}

type Item struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

type CompletedOrders struct {
	ID         int         `json:"id"`
	OrderID    int         `json:"order_id"`
	VendorID   int         `json:"vendor_id"`
	UserID     int         `json:"user_id"`
	Status     string      `json:"status"`
	OrderType  OrderType   `json:"order_type"`
	Items      []OrderItem `json:"items"`
	TotalPrice float64     `json:"total_price"`
	CreatedAt  string      `json:"created_at"`
}

type OrderItem struct {
	ID     int `json:"id"`
	MealID int `json:"meal_id"`
}

type OngoingSubscription struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	VendorID  int    `json:"vendor_id"`
	StartDate string `json:"start_date"`
	RollOver  string `json:"roll_over"`
}

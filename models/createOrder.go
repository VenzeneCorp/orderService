package models

type CreateOrder struct {
	UserID      int       `json:"user_id"`
	VendorID    int       `json:"vendor_id"`
	VendorName  string    `json:"vendor_name"`
	Amount      int       `json:"amount"`
	Discount    int       `json:"discount"`
	FinalAmount int       `json:"final_amount"`
	OrderType   OrderType `json:"order_type"`
}

type CreateLiveOrder struct {
	MealID   string `json:"meal_id"`
	MealName string `json:"meal_name"`
	Quantity int    `json:"quantity"`
	Veg      bool   `json:"veg"`
	Price    int    `json:"price"`
}

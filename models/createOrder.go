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

type CreateSubscription struct {
	MealCount                int    `json:"meal_count"`
	RemainingMealCount       int    `json:"remaining_meal_count"`
	RollOverCount            int    `json:"roll_over_count"` // number of meals that can be rolled over
	BreakfastID              string `json:"breakfast_id"`
	LunchID                  string `json:"lunch_id"`
	DinnerID                 string `json:"dinner_id"`
	BreakfastDeliveryAddress string `json:"breakfast_delivery_address"`
	LunchDeliveryAddress     string `json:"lunch_delivery_address"`
	DinnerDeliveryAddress    string `json:"dinner_delivery_address"`
}

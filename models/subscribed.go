package models

type Subscription struct {
	ID                       uint64 `json:"id"`
	OrderID                  uint64 `json:"order_id"`
	Status                   Status `json:"status"`
	MealCount                int    `json:"meal_count"`
	RemainingMealCount       int    `json:"remaining_meal_count"`
	RollOverCount            int    `json:"roll_over_count"`
	BreakfastID              string `json:"breakfast_id"`
	LunchID                  string `json:"lunch_id"`
	DinnerID                 string `json:"dinner_id"`
	BreakfastDeliveryAddress string `json:"breakfast_delivery_address"`
	LunchDeliveryAddress     string `json:"lunch_delivery_address"`
	DinnerDeliveryAddress    string `json:"dinner_delivery_address"`
	CreatedAt                int64  `json:"created_at"`
	UpdatedAt                int64  `json:"updated_at"`
}

type ScheduledOrder struct {}

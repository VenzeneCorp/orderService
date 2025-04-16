package models

type Subscribed struct {
	ID                       int    `json:"id"` // incremental id
	OrderID                  string `json:"order_id"`
	MealCount                int    `json:"meal_count"`      // number of meals
	RollOverCount            int    `json:"roll_over_count"` // number of meals that can be rolled over
	BreakfastID              string `json:"breakfast_id"`
	LunchID                  string `json:"lunch_id"`
	DinnerID                 string `json:"dinner_id"`
	BreakfastDeliveryAddress string `json:"breakfast_delivery_address"`
	LunchDeliveryAddress     string `json:"lunch_delivery_address"`
	DinnerDeliveryAddress    string `json:"dinner_delivery_address"`
	CreatedAt                int64  `json:"created_at"`
	UpdatedAt                int64  `json:"updated_at"`
}

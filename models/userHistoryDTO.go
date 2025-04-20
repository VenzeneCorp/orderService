package models

type OrderHistory struct {
	OrderID    string  `json:"order_id"`
	Meals      []Meals `json:"meals"`
	Amount     int     `json:"amount"`
	Status     string  `json:"status"`
	OrderType  string  `json:"order_type"`
	Vendor_id  int     `json:"vendor_id"`
	VendorName string  `json:"vendor_name"`
	OrderTime  string  `json:"order_time"`
}

type Meals struct {
	MealName string `json:"meal_name"`
	Quantity int    `json:"quantity"`
}

type SubscriptionHistory struct {
	ID                       int           `json:"id"`
	RemainingMeals           int           `json:"remaining_meals"`
	RollOverCount            int           `json:"roll_over_count"`
	BreakfastID              MealPlanInfo  `json:"breakfast_id"`
	LunchID                  MealPlanInfo  `json:"lunch_id"`
	DinnerID                 MealPlanInfo  `json:"dinner_id"`
	BreakfastDeliveryAddress UserAddresses `json:"breakfast_delivery_address"`
	LunchDeliveryAddress     UserAddresses `json:"lunch_delivery_address"`
	DinnerDeliveryAddress    UserAddresses `json:"dinner_delivery_address"`
}

// get this from mealService
type MealPlanInfo struct {
	MealPlanID     int    `json:"meal_plan_id"`
	MealPlanName   string `json:"meal_plan_name"`
	RestaurantName string `json:"restaurant_name"`
	Veg            bool   `json:"veg"`
	Description    string `json:"description"`
	ImageURL       string `json:"image_url"`
}

// get this from loginSignupService
type UserAddresses struct {
	AddressID   int    `json:"address_id"`
	AddressName string `json:"address_name"`
}

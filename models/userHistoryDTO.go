package models

import (
	loginSignupModels "github.com/venzene/loginSignup/models"
	mealSvcModels "github.com/venzene/mealService/models"
)

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
	ID                       int                       `json:"id"`
	RemainingMeals           int                       `json:"remaining_meals"`
	RollOverCount            int                       `json:"roll_over_count"`
	BreakfastID              mealSvcModels.MealPlan    `json:"breakfast_id"`
	LunchID                  mealSvcModels.MealPlan    `json:"lunch_id"`
	DinnerID                 mealSvcModels.MealPlan    `json:"dinner_id"`
	BreakfastDeliveryAddress loginSignupModels.Address `json:"breakfast_delivery_address"`
	LunchDeliveryAddress     loginSignupModels.Address `json:"lunch_delivery_address"`
	DinnerDeliveryAddress    loginSignupModels.Address `json:"dinner_delivery_address"`
}

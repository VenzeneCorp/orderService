package models

import (
	loginSigupModels "github.com/VenzeneCorp/loginSignup/models"
	mealsSvcModels "github.com/VenzeneCorp/mealService/models"
)

type VendorDashboard struct {
	TotalMealsToDeliver int                     `json:"total_meals_to_deliver"`
	MealDetails         []MealsToDeliver        `json:"meal_details"`
	MealType            mealsSvcModels.MealType `json:"meal_type"`
	DeliveryDetails     []DeliveryDetails       `json:"delivery_details"`
}

type DeliveryDetails struct {
	Address   loginSigupModels.Address `json:"address"`
	MealName  string                   `json:"meal_name"`
	UserPhone string                   `json:"user_phone"`
}

type MealsToDeliver struct {
	MealID   string `json:"meal_id"`
	MealName string `json:"meal_name"`
	Quantity int    `json:"quantity"`
}

type SubscriptionAnalytics struct {
	TotalActiveSubscriptions int               `json:"total_active_subscriptions"`
	MostPopularPlan          []PopularMealPlan `json:"most_popular_plan"`
}

type PopularMealPlan struct {
	MealPlanID         string  `json:"plan_id"`
	MealPlanName       string  `json:"plan_name"`
	TotalSubscriptions int     `json:"total_subscriptions"`
	TotalSales         float64 `json:"total_sales"`
}

type LiveOrderAnalytics struct {
	TotalLiveOrders int `json:"total_live_orders"`
	TotalSales      int `json:"total_sales"`
}

type SalesAnalytics struct {
	TotalSales  float64 `json:"total_sales"`
	TotalOrders int     `json:"total_orders"`

	TotalLiveOrderSales float64 `json:"total_live_order_sales"`
	TotalLiveOrders     int     `json:"total_live_orders"`

	TotalSubscriptionSales float64 `json:"total_subscription_sales"`
	TotalSubscriptions     int     `json:"total_subscriptions"`
}

type ProductAnalytics struct {
	MealID      string  `json:"meal_id"`
	MealName    string  `json:"meal_name"`
	TotalSales  float64 `json:"total_sales"`
	TotalOrders int     `json:"total_orders"`
}

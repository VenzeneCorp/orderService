package mealserviceclient

import (
	"context"

	mealsSvcModels "github.com/VenzeneCorp/mealService/models"
)

type MealServiceClient interface {
	GetMeals(ctx context.Context, mealIDs []string) ([]mealsSvcModels.Meal, error)
	GetMealPlans(ctx context.Context, mealPlanIDs []string) ([]mealsSvcModels.MealPlan, error)
}

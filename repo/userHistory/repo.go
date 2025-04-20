package userhistory

import (
	"context"

	"github.com/VenzeneCorp/orderService/models"
)

type Repository interface {
	PlaceOrder(ctx context.Context, order models.Orders) error
	GetSubscriptionInfo(ctx context.Context, userID string) (models.Subscribed, error)
	GetUserHistory(ctx context.Context, userID string) ([]models.OrderHistory, error)
	GetUserHistoryByOrderID(ctx context.Context, orderID string) ([]models.ItemHistory, error)
}

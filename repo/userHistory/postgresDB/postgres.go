package postgresdb

import (
	"context"

	"github.com/VenzeneCorp/orderService/models"
)

type SQL interface {
	PlaceOrder(ctx context.Context, order models.CreateOrder) error
	CreateSubscription(ctx context.Context, subscription models.Subscribed) error
	GetSubscriptionInfo(ctx context.Context, userID string) (models.Subscribed, error)
	GetUserHistory(ctx context.Context, userID string) ([]models.OrderHistory, error)
	GetUserHistoryByOrderID(ctx context.Context, orderID string) ([]models.ItemOrdered, error)
}

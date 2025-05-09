package postgresdb

import (
	"context"

	"github.com/VenzeneCorp/orderService/models"
)

type SQL interface {
	PlaceLiveOrder(ctx context.Context, order models.CreateOrder, liveOrder []models.CreateLiveOrder) error
	PlaceSubscriptionOrder(ctx context.Context, order models.CreateOrder, subscription models.CreateSubscription) error
	
	CancelOrder(ctx context.Context, orderID string) error

	StartOrder(ctx context.Context, order models.CreateLiveOrder, orderId uint64) error
	DeliverOrder(ctx context.Context, orderId uint64) error

	GetSubscriptionInfo(ctx context.Context, userID string) (models.Subscription, error)
	GetUserHistory(ctx context.Context, userID string) ([]models.ItemOrdered, error)
	GetUserSubscriptionHistory(ctx context.Context, userID string) ([]models.ItemOrdered, error)
}

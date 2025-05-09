package orderhistory

import (
	"context"

	"github.com/VenzeneCorp/orderService/models"
)

type Repository interface {
	PlaceLiveOrder(ctx context.Context, userID string, order models.CreateOrder, liveOrder []models.CreateLiveOrder) error
	PlaceSubscriptionOrder(ctx context.Context, userID string, order models.CreateOrder, subscription models.CreateSubscription) error
	CancelOrder(ctx context.Context, userID string, orderID string) error
	GetSubscriptionInfo(ctx context.Context, userID string) (models.Subscription, error)
	GetUserHistory(ctx context.Context, userID string) ([]models.ItemOrdered, error)
	GetUserSubscriptionHistory(ctx context.Context, userID string) ([]models.ItemOrdered, error)
}

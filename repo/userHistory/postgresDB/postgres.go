package postgresdb

import (
	"context"

	"github.com/VenzeneCorp/orderService/models"
)

type SQL interface {
	PlaceOrder(ctx context.Context, order models.Orders) error
	GetSubscriptionInfo(ctx context.Context, userID int) (models.Subscribed, error)
	GetUserHistory(ctx context.Context, userID int) ([]models.OrderHistory, error)
	GetUserHistoryByOrderID(ctx context.Context, orderID string) ([]models.ItemHistory, error)
}

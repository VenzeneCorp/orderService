package vendormgmt

import (
	"context"

	"github.com/VenzeneCorp/orderService/models"
)

type Service interface {
	GetSubscription(ctx context.Context, vendorID string) ([]models.Subscription, error)
	GetScheduledOrders(ctx context.Context, vendorID string) ([]models.ItemOrdered, error)
	GetLiveOrder(ctx context.Context, vendorID string) ([]models.ItemOrdered, error)
	GetSalesByPeriod(ctx context.Context, vendorID string, period int) ([]models.Orders, error)
}

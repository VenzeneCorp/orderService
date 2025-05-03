package vendorpostgresdb

import (
	"context"

	"github.com/VenzeneCorp/orderService/models"
)

type SQL interface {
	GetSubscriptionAnalytics(ctx context.Context, vendorID string) (models.SubscriptionAnalytics, error)
	GetScheduledOrders(ctx context.Context, vendorID string) ([]models.ScheduledOrder, error)
	GetLiveOrderAnalytics(ctx context.Context, vendorID string) (models.LiveOrderAnalytics, error)
	GetVendorDashboard(ctx context.Context, vendorID string) (models.VendorDashboard, error)
	GetSalesAnalyticsByPeriod(ctx context.Context, vendorID string, period string) (models.SalesAnalytics, error)
}

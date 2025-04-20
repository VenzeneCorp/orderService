package vendorpostgresdb

import (
	"context"

	"github.com/VenzeneCorp/orderService/models"
)

type SQL interface {
	GetSubscriptionAnalytics(ctx context.Context, vendorID string) (models.SubscriptionAnalytics, error)
	GetLiveOrderAnalytics(ctx context.Context, vendorID string) (models.LiveOrderAnalytics, error)
	GetVendorDashboard(ctx context.Context, vendorID string) (models.VendorDashboard, error)
	GetSalesAnalyticsByPeriod(ctx context.Context, vendorID string, period string) (models.SalesAnalytics, error)
	GetTopSellingProducts(ctx context.Context, vendorID string, limit int) ([]models.ProductAnalytics, error)
	GetRevenueByCategory(ctx context.Context, vendorID string) (map[string]float64, error)
	UpdateAnalyticsCache(ctx context.Context, vendorID string) error
	GetInventoryTurnoverRate(ctx context.Context, vendorID string) (float64, error)
}

package vendoranalytics

import (
	"context"

	"github.com/VenzeneCorp/orderService/models"
	vendorpostgresdb "github.com/VenzeneCorp/orderService/repo/vendorAnalytics/vendorPostgresDB"
)

type VendorAnalyticsRepo struct {
	sqlDB vendorpostgresdb.SQL
}

func NewVendorAnalyticsRepo(db *vendorpostgresdb.SQL) Repository {
	return &VendorAnalyticsRepo{
		sqlDB: *db,
	}
}

func (r *VendorAnalyticsRepo) GetSubscription(ctx context.Context, vendorID string) ([]models.Subscription, error) {
	return r.sqlDB.GetSubscription(ctx, vendorID)
}

func (r *VendorAnalyticsRepo) GetScheduledOrders(ctx context.Context, vendorID string) ([]models.ItemOrdered, error) {
	return r.sqlDB.GetScheduledOrders(ctx, vendorID)
}

func (r *VendorAnalyticsRepo) GetLiveOrder(ctx context.Context, vendorID string) ([]models.ItemOrdered, error) {
	return r.sqlDB.GetLiveOrder(ctx, vendorID)
}

func (r *VendorAnalyticsRepo) GetSalesByPeriod(ctx context.Context, vendorID string, period int) ([]models.Orders, error) {
	return r.sqlDB.GetSalesByPeriod(ctx, vendorID, period)
}

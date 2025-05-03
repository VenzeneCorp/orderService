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

func (r *VendorAnalyticsRepo) GetSubscriptionAnalytics(ctx context.Context, vendorID string) (models.SubscriptionAnalytics, error) {
	return r.sqlDB.GetSubscriptionAnalytics(ctx, vendorID)
}

func (r *VendorAnalyticsRepo) GetScheduledOrders(ctx context.Context, vendorID string) ([]models.ScheduledOrder, error) {
	return r.sqlDB.GetScheduledOrders(ctx, vendorID)
}

func (r *VendorAnalyticsRepo) GetLiveOrderAnalytics(ctx context.Context, vendorID string) (models.LiveOrderAnalytics, error) {
	return r.sqlDB.GetLiveOrderAnalytics(ctx, vendorID)
}

func (r *VendorAnalyticsRepo) GetVendorDashboard(ctx context.Context, vendorID string) (models.VendorDashboard, error) {
	return r.sqlDB.GetVendorDashboard(ctx, vendorID)
}

func (r *VendorAnalyticsRepo) GetSalesAnalyticsByPeriod(ctx context.Context, vendorID string, period string) (models.SalesAnalytics, error) {
	return r.sqlDB.GetSalesAnalyticsByPeriod(ctx, vendorID, period)
}

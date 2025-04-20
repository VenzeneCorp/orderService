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

func (r *VendorAnalyticsRepo) GetLiveOrderAnalytics(ctx context.Context, vendorID string) (models.LiveOrderAnalytics, error) {
	return r.sqlDB.GetLiveOrderAnalytics(ctx, vendorID)
}

func (r *VendorAnalyticsRepo) GetVendorDashboard(ctx context.Context, vendorID string) (models.VendorDashboard, error) {
	return r.sqlDB.GetVendorDashboard(ctx, vendorID)
}

func (r *VendorAnalyticsRepo) GetSalesAnalyticsByPeriod(ctx context.Context, vendorID string, period string) (models.SalesAnalytics, error) {
	return r.sqlDB.GetSalesAnalyticsByPeriod(ctx, vendorID, period)
}

func (r *VendorAnalyticsRepo) GetTopSellingProducts(ctx context.Context, vendorID string, limit int) ([]models.ProductAnalytics, error) {
	return r.sqlDB.GetTopSellingProducts(ctx, vendorID, limit)
}

func (r *VendorAnalyticsRepo) GetRevenueByCategory(ctx context.Context, vendorID string) (map[string]float64, error) {
	return r.sqlDB.GetRevenueByCategory(ctx, vendorID)
}

func (r *VendorAnalyticsRepo) UpdateAnalyticsCache(ctx context.Context, vendorID string) error {
	return r.sqlDB.UpdateAnalyticsCache(ctx, vendorID)
}

func (r *VendorAnalyticsRepo) GetInventoryTurnoverRate(ctx context.Context, vendorID string) (float64, error) {
	return r.sqlDB.GetInventoryTurnoverRate(ctx, vendorID)
}

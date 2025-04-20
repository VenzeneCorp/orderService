package vendorpostgresdb

import (
	"context"

	"github.com/VenzeneCorp/orderService/models"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) SQL {
	return &Repository{
		DB: db,
	}
}

func (r *Repository) GetSubscriptionAnalytics(ctx context.Context, vendorID string) (models.SubscriptionAnalytics, error) {
	var analytics models.SubscriptionAnalytics
	err := r.DB.WithContext(ctx).Where("vendor_id = ?", vendorID).First(&analytics).Error
	return analytics, err
}

func (r *Repository) GetLiveOrderAnalytics(ctx context.Context, vendorID string) (models.LiveOrderAnalytics, error) {
	var analytics models.LiveOrderAnalytics
	err := r.DB.WithContext(ctx).Where("vendor_id = ?", vendorID).First(&analytics).Error
	return analytics, err
}

func (r *Repository) GetVendorDashboard(ctx context.Context, vendorID string) (models.VendorDashboard, error) {
	var dashboard models.VendorDashboard
	err := r.DB.WithContext(ctx).Where("vendor_id = ?", vendorID).First(&dashboard).Error
	return dashboard, err
}

func (r *Repository) GetSalesAnalyticsByPeriod(ctx context.Context, vendorID string, period string) (models.SalesAnalytics, error) {
	var analytics models.SalesAnalytics
	err := r.DB.WithContext(ctx).Where("vendor_id = ? AND period = ?", vendorID, period).First(&analytics).Error
	return analytics, err
}

func (r *Repository) GetTopSellingProducts(ctx context.Context, vendorID string, limit int) ([]models.ProductAnalytics, error) {
	var products []models.ProductAnalytics
	err := r.DB.WithContext(ctx).Where("vendor_id = ?", vendorID).Order("sales DESC").Limit(limit).Find(&products).Error
	return products, err
}

func (r *Repository) GetRevenueByCategory(ctx context.Context, vendorID string) (map[string]float64, error) {
	var results []struct {
		Category string
		Revenue  float64
	}
	err := r.DB.WithContext(ctx).Raw("SELECT category, SUM(revenue) as revenue FROM sales WHERE vendor_id = ? GROUP BY category", vendorID).Scan(&results).Error
	if err != nil {
		return nil, err
	}

	revenueByCategory := make(map[string]float64)
	for _, result := range results {
		revenueByCategory[result.Category] = result.Revenue
	}
	return revenueByCategory, nil
}

func (r *Repository) UpdateAnalyticsCache(ctx context.Context, vendorID string) error {
	// Assuming a stored procedure or function exists for updating the cache
	return r.DB.WithContext(ctx).Exec("CALL update_analytics_cache(?)", vendorID).Error
}

func (r *Repository) GetInventoryTurnoverRate(ctx context.Context, vendorID string) (float64, error) {
	var turnoverRate float64
	err := r.DB.WithContext(ctx).Raw("SELECT calculate_inventory_turnover_rate(?)", vendorID).Scan(&turnoverRate).Error
	return turnoverRate, err
}

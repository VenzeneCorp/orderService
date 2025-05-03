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

func (r *Repository) GetScheduledOrders(ctx context.Context, vendorID string) ([]models.ScheduledOrder, error) {
	var orders []models.ScheduledOrder
	err := r.DB.WithContext(ctx).Where("vendor_id = ?", vendorID).Find(&orders).Error
	return orders, err
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
package vendorpostgresdb

import (
	"context"
	"time"

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

func (r *Repository) GetSubscription(ctx context.Context, vendorID string) ([]models.Subscription, error) {
	var subscription []models.Subscription
	err := r.DB.WithContext(ctx).Table("subscriptions").
		Joins("JOIN orders o ON subscriptions.order_id = o.id").
		Where("o.vendor_id = ?", vendorID).
		Find(&subscription).Error
	if err != nil {
		return nil, err
	}
	return subscription, nil
}

func (r *Repository) GetScheduledOrders(ctx context.Context, vendorID string) ([]models.ItemOrdered, error) {
	var items []models.ItemOrdered
	err := r.DB.Table("item_ordered AS io").
		Joins("JOIN orders o ON io.order_id = o.id").
		Where("io.status = ? AND o.vendor_id = ?", models.Pending, vendorID).
		Find(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (r *Repository) GetLiveOrder(ctx context.Context, vendorID string) ([]models.ItemOrdered, error) {
	var items []models.ItemOrdered
	err := r.DB.Table("item_ordered AS io").
		Joins("JOIN orders o ON io.order_id = o.id").
		Where("io.status = ? AND o.vendor_id = ?", models.Pending, vendorID).
		Find(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (r *Repository) GetSalesByPeriod(ctx context.Context, vendorID string, period int) ([]models.Orders, error) {
	var analytics []models.Orders
	startTime := time.Now().AddDate(0, 0, -period).Unix()

	err := r.DB.WithContext(ctx).
		Where("vendor_id = ? AND created_at >= ?", vendorID, startTime).
		Find(&analytics).Error

	return analytics, err
}

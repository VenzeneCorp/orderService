package payments

import (
	"context"
	"time"

	"github.com/VenzeneCorp/orderService/models"
	"gorm.io/gorm"
)

type Repository interface {
	CreateRazorpayOrder(ctx context.Context, orderID uint64, amount int, receipt string, razorpayOrderID string) (string, error)
	UpdateOrderStatus(ctx context.Context, razorpayOrderID string, status int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateRazorpayOrder(ctx context.Context, orderID uint64, amount int, receipt string, razorpayOrderID string) (string, error) {
	timeNow := time.Now().Unix()
	razorpayOrder := models.RazorpayOrder{
		ID:             razorpayOrderID,
		OrderID:        orderID,
		Amount:         amount,
		Receipt:        receipt,
		RazorpayStatus: models.Started,
		CreatedAt:      timeNow,
		UpdatedAt:      timeNow,
	}

	if err := r.db.WithContext(ctx).Create(&razorpayOrder).Error; err != nil {
		return "", err
	}

	return razorpayOrder.ID, nil
}

func (r *repository) UpdateOrderStatus(ctx context.Context, razorpayOrderID string, status int) error {
	timeNow := time.Now().Unix()
	razorpayOrder := models.RazorpayOrder{
		ID:             razorpayOrderID,
		RazorpayStatus: models.RazorpayStatus(status),
		UpdatedAt:      timeNow,
	}
	if err := r.db.WithContext(ctx).Model(&razorpayOrder).Where("id = ?", razorpayOrderID).Updates(razorpayOrder).Error; err != nil {
		return err
	}
	return nil
}

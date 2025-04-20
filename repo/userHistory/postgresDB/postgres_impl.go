package postgresdb

import (
	"context"
	"time"

	"github.com/VenzeneCorp/orderService/models"
	"github.com/VenzeneCorp/orderService/utils"
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

func (p *Repository) PlaceOrder(ctx context.Context, order models.CreateOrder) error {

	orderID, err := utils.GenerateOrderID()
	if err != nil {
		return err
	}

	newOrder := models.Orders{
		ID:          orderID,
		UserID:      order.UserID,
		VendorID:    order.VendorID,
		VendorName:  order.VendorName,
		Amount:      order.Amount,
		Discount:    order.Discount,
		FinalAmount: order.FinalAmount,
		OrderType:   order.OrderType,
		Status:      models.Pending,
		CreatedAt:   time.Now().Unix(),
	}

	err = p.DB.WithContext(ctx).Create(&newOrder).Error
	if err != nil {
		return err
	}

	return nil
}

func (p *Repository) CreateSubscription(ctx context.Context, subscription models.Subscribed) error {
	// Implement the logic to create subscription in the database
	// This is a placeholder implementation
	return nil
}

func (p *Repository) GetSubscriptionInfo(ctx context.Context, userID string) (models.Subscribed, error) {
	// Implement the logic to get subscription info from the database
	// This is a placeholder implementation
	return models.Subscribed{}, nil
}
func (p *Repository) GetUserHistory(ctx context.Context, userID string) ([]models.OrderHistory, error) {
	// Implement the logic to get user history from the database
	// This is a placeholder implementation
	return []models.OrderHistory{}, nil
}
func (p *Repository) GetUserHistoryByOrderID(ctx context.Context, orderID string) ([]models.ItemOrdered, error) {
	// Implement the logic to get user history by order ID from the database
	// This is a placeholder implementation
	return []models.ItemOrdered{}, nil
}

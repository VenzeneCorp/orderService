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

func (p *Repository) PlaceLiveOrder(ctx context.Context, order models.CreateOrder, liveOrder models.CreateLiveOrder) error {

	orderId, err := utils.GenerateID()
	if err != nil {
		return err
	}

	newOrder := models.Orders{
		ID:          orderId,
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

	//TODO: add to ITEM_ORDERED table with txn

	return nil
}

func (p *Repository) PlaceSubscriptionOrder(ctx context.Context, order models.CreateOrder, subscription models.CreateSubscription) error {

	orderId, err := utils.GenerateID()
	if err != nil {
		return err
	}
	newOrder := models.Orders{
		ID:          orderId,
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

	//TODO: add to ITEM_ORDERED table with txn

	return nil
}

func (p *Repository) StartOrder(ctx context.Context, order models.CreateLiveOrder, orderId uint64) error {

	newID, err := utils.GenerateID()
	if err != nil {
		return err
	}

	newOrder := models.ItemOrdered{
		ID:          newID,
		OrderID:     orderId,
		MealID:      order.MealID,
		MealName:    order.MealName,
		Quantity:    order.Quantity,
		Veg:         order.Veg,
		Price:       order.Price,
		DeliveredAt: time.Now().Unix(),
	}

	err = p.DB.WithContext(ctx).Create(&newOrder).Error
	if err != nil {
		return err
	}

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

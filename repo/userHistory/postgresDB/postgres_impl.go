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

	return p.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		orderId, err := utils.GenerateID()
		if err != nil {
			return err
		}

		timeInSec := time.Now().Unix()

		newOrder := models.Orders{
			ID:          orderId,
			UserID:      order.UserID,
			VendorID:    order.VendorID,
			VendorName:  order.VendorName,
			Amount:      order.Amount,
			Discount:    order.Discount,
			FinalAmount: order.FinalAmount,
			OrderType:   order.OrderType,
			CreatedAt:   timeInSec,
		}

		if err := tx.Create(&newOrder).Error; err != nil {
			return err
		}

		newID, err := utils.GenerateID()
		if err != nil {
			return err
		}

		newLiveOrder := models.ItemOrdered{
			ID:          newID,
			OrderID:     orderId,
			OrderType:   models.LiveOrder,
			MealID:      liveOrder.MealID,
			MealName:    liveOrder.MealName,
			Quantity:    liveOrder.Quantity,
			Veg:         liveOrder.Veg,
			Price:       liveOrder.Price,
			DeliveredAt: nil,
		}

		if err := tx.Create(&newLiveOrder).Error; err != nil {
			return err
		}

		return nil
	})
}

func (p *Repository) PlaceSubscriptionOrder(ctx context.Context, order models.CreateOrder, subscription models.CreateSubscription) error {

	return p.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		orderId, err := utils.GenerateID()
		if err != nil {
			return err
		}

		timeInSec := time.Now().Unix()

		newOrder := models.Orders{
			ID:          orderId,
			UserID:      order.UserID,
			VendorID:    order.VendorID,
			VendorName:  order.VendorName,
			Amount:      order.Amount,
			Discount:    order.Discount,
			FinalAmount: order.FinalAmount,
			OrderType:   order.OrderType,
			CreatedAt:   timeInSec,
		}

		if err := tx.Create(&newOrder).Error; err != nil {
			return err
		}

		subscriptionId, err := utils.GenerateID()
		if err != nil {
			return err
		}

		newSubscription := models.Subscribed{
			ID:                       subscriptionId,
			OrderID:                  orderId,
			MealCount:                subscription.MealCount,
			RemainingMealCount:       subscription.MealCount,
			RollOverCount:            subscription.RollOverCount,
			BreakfastID:              subscription.BreakfastID,
			LunchID:                  subscription.LunchID,
			DinnerID:                 subscription.DinnerID,
			BreakfastDeliveryAddress: subscription.BreakfastDeliveryAddress,
			LunchDeliveryAddress:     subscription.LunchDeliveryAddress,
			DinnerDeliveryAddress:    subscription.DinnerDeliveryAddress,
			CreatedAt:                timeInSec,
			UpdatedAt:                timeInSec,
		}

		if err := tx.Create(&newSubscription).Error; err != nil {
			return err
		}

		return nil
	})
}

func (p *Repository) StartOrder(ctx context.Context, order models.CreateLiveOrder, orderId uint64) error {

	newID, err := utils.GenerateID()
	if err != nil {
		return err
	}

	deliveredAt := time.Now().Unix()

	newOrder := models.ItemOrdered{
		ID:          newID,
		OrderID:     orderId,
		OrderType:   models.SubscriptionOrder,
		MealID:      order.MealID,
		MealName:    order.MealName,
		Quantity:    order.Quantity,
		Veg:         order.Veg,
		Price:       order.Price,
		DeliveredAt: &deliveredAt,
	}

	err = p.DB.WithContext(ctx).Create(&newOrder).Error
	if err != nil {
		return err
	}

	return nil
}

func (p *Repository) DeliverOrder(ctx context.Context, orderId uint64) error {
	deliveryTime := time.Now().Unix()

	return p.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		orderItem := models.ItemOrdered{}

		if err := tx.Where("order_id = ?", orderId).First(&orderItem).Error; err != nil {
			return err
		}

		orderItem.DeliveredAt = &deliveryTime
		orderItem.Status = models.Completed

		if err := tx.Save(&orderItem).Error; err != nil {
			return err
		}

		if orderItem.OrderType == models.SubscriptionOrder {
			subscription := models.Subscribed{}
			if err := tx.Where("order_id = ?", orderItem.OrderID).First(&subscription).Error; err != nil {
				return err
			}
			subscription.RemainingMealCount--
			subscription.UpdatedAt = deliveryTime

			if err := tx.Save(&subscription).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func (p *Repository) GetSubscriptionInfo(ctx context.Context, userID string) (models.Subscribed, error) {
	subscription := models.Subscribed{}
	err := p.DB.WithContext(ctx).Where("user_id = ?", userID).First(&subscription).Error
	if err != nil {
		return models.Subscribed{}, err
	}
	return subscription, nil
}

func (p *Repository) GetUserHistory(ctx context.Context, userID string) ([]models.ItemOrdered, error) {
	itemHistory := []models.ItemOrdered{}
	err := p.DB.WithContext(ctx).Where("user_id = ?", userID).Find(&itemHistory).Error
	if err != nil {
		return nil, err
	}
	if len(itemHistory) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return itemHistory, nil
}

func (p *Repository) GetUserSubscriptionHistory(ctx context.Context, userID string) ([]models.ItemOrdered, error) {
	subscriptionItemHistory := []models.ItemOrdered{}
	err := p.DB.WithContext(ctx).Where("user_id = ? AND order_type = ?", userID, models.SubscriptionOrder).Find(&subscriptionItemHistory).Error
	if err != nil {
		return nil, err
	}
	if len(subscriptionItemHistory) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return subscriptionItemHistory, nil
}

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

func (p *Repository) PlaceLiveOrder(ctx context.Context, userID string, order models.CreateOrder, liveOrders []models.CreateLiveOrder) (models.Orders, error) {
	var newOrder models.Orders

	err := p.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		orderId, err := utils.GenerateID()
		if err != nil {
			return err
		}

		timeInSec := time.Now().Unix()

		newOrder = models.Orders{
			ID:          orderId,
			UserID:      userID,
			VendorID:    order.VendorID,
			VendorName:  order.VendorName,
			Amount:      order.Amount,
			Discount:    order.Discount,
			FinalAmount: order.FinalAmount,
			OrderType:   order.OrderType,
			OrderStatus: models.OrderCreated,
			CreatedAt:   timeInSec,
		}

		if err := tx.Create(&newOrder).Error; err != nil {
			return err
		}

		var items []models.ItemOrdered

		for _, lo := range liveOrders {
			itemID, err := utils.GenerateID()
			if err != nil {
				return err
			}

			item := models.ItemOrdered{
				ID:          itemID,
				OrderID:     orderId,
				OrderType:   models.LiveOrder,
				MealID:      lo.MealID,
				MealName:    lo.MealName,
				Quantity:    lo.Quantity,
				Veg:         lo.Veg,
				Price:       lo.Price,
				DeliveredAt: nil,
			}
			items = append(items, item)
		}

		if len(items) > 0 {
			if err := tx.Create(&items).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return models.Orders{}, err
	}

	return newOrder, nil
}

func (p *Repository) PlaceSubscriptionOrder(ctx context.Context, userID string, order models.CreateOrder, subscription models.CreateSubscription) (models.Orders, error) {

	var newOrder models.Orders

	err := p.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		orderId, err := utils.GenerateID()
		if err != nil {
			return err
		}

		timeInSec := time.Now().Unix()

		newOrder = models.Orders{
			ID:          orderId,
			UserID:      userID,
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

		newSubscription := models.Subscription{
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

	if err != nil {
		return models.Orders{}, err
	}

	return newOrder, nil
}

func (p *Repository) CancelOrder(ctx context.Context, userID string, orderID string) error {
	if userID == "-1" {
		err := p.DB.WithContext(ctx).
			Model(&models.Orders{}).
			Where("id = ?", orderID).
			Update("order_status", models.Cancelled).Error
		if err != nil {
			return err
		}
	} else {
		err := p.DB.WithContext(ctx).
			Model(&models.Orders{}).
			Where("id = ? AND user_id = ?", orderID, userID).
			Update("order_status", models.Cancelled).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *Repository) StartOrder(ctx context.Context, order models.CreateLiveOrder, orderId uint64) error {

	newID, err := utils.GenerateID()
	if err != nil {
		return err
	}

	newOrder := models.ItemOrdered{
		ID:        newID,
		OrderID:   orderId,
		OrderType: models.SubscriptionOrder,
		MealID:    order.MealID,
		MealName:  order.MealName,
		Quantity:  order.Quantity,
		Veg:       order.Veg,
		Price:     order.Price,
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
			subscription := models.Subscription{}
			if err := tx.Where("order_id = ?", orderItem.OrderID).First(&subscription).Error; err != nil {
				return err
			}
			subscription.RemainingMealCount--
			subscription.UpdatedAt = deliveryTime

			if subscription.RemainingMealCount == 0 {
				subscription.Status = models.Completed
				if err := tx.Save(&subscription).Error; err != nil {
					return err
				}
				if err := tx.Model(&models.Orders{}).
					Where("id = ?", orderItem.OrderID).
					Update("order_status", models.Completed).Error; err != nil {
					return err
				}
			} else {
				if err := tx.Save(&subscription).Error; err != nil {
					return err
				}
			}
		}
		return nil
	})
}

func (p *Repository) GetSubscriptionInfo(ctx context.Context, userID string) (models.Subscription, error) {
	subscription := models.Subscription{}
	err := p.DB.WithContext(ctx).Where("user_id = ?", userID).First(&subscription).Error
	if err != nil {
		return models.Subscription{}, err
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

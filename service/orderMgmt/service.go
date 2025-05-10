package ordermgmt

import (
	"context"

	"github.com/VenzeneCorp/orderService/models"
)

type Service interface {
	PlaceLiveOrder(ctx context.Context, userID string, order models.CreateOrder, liveOrder []models.CreateLiveOrder) (models.RazorpayPaymentRespone, error)
	UpdateOrderStatus(ctx context.Context, userID, orderID string, status models.RazorpaySuccessRequest) error
	PlaceSubscriptionOrder(ctx context.Context, userID string, order models.CreateOrder, subscription models.CreateSubscription) (models.RazorpayPaymentRespone, error)
	CancelOrder(ctx context.Context, userID string, orderID string) error
	GetSubscriptionInfo(ctx context.Context, userID string) (models.Subscription, error)
	GetUserHistory(ctx context.Context, userID string) ([]models.ItemOrdered, error)
	GetUserSubscriptionHistory(ctx context.Context, userID string) ([]models.ItemOrdered, error)
}

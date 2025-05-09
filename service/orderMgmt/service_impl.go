package ordermgmt

import (
	"context"

	"github.com/VenzeneCorp/orderService/models"
	orderhistory "github.com/VenzeneCorp/orderService/repo/orderHistory"
)

type ServiceImpl struct {
	repo orderhistory.Repository
}

func NewService(repo orderhistory.Repository) Service {
	return &ServiceImpl{
		repo: repo,
	}
}

func (s *ServiceImpl) PlaceLiveOrder(ctx context.Context, userID string, order models.CreateOrder, liveOrder []models.CreateLiveOrder) error {
	return s.repo.PlaceLiveOrder(ctx, userID, order, liveOrder)
}

func (s *ServiceImpl) PlaceSubscriptionOrder(ctx context.Context, userID string, order models.CreateOrder, subscription models.CreateSubscription) error {
	return s.repo.PlaceSubscriptionOrder(ctx, userID, order, subscription)
}

func (s *ServiceImpl) CancelOrder(ctx context.Context, userID string, orderID string) error {
	return s.repo.CancelOrder(ctx, userID, orderID)
}

func (s *ServiceImpl) GetSubscriptionInfo(ctx context.Context, userID string) (models.Subscription, error) {
	return s.repo.GetSubscriptionInfo(ctx, userID)
}

func (s *ServiceImpl) GetUserHistory(ctx context.Context, userID string) ([]models.ItemOrdered, error) {
	return s.repo.GetUserHistory(ctx, userID)
}

func (s *ServiceImpl) GetUserSubscriptionHistory(ctx context.Context, userID string) ([]models.ItemOrdered, error) {
	return s.repo.GetUserSubscriptionHistory(ctx, userID)
}

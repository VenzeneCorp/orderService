package vendormgmt

import (
	"context"

	"github.com/VenzeneCorp/orderService/models"
	vendoranalytics "github.com/VenzeneCorp/orderService/repo/vendorAnalytics"
)

type ServiceImpl struct {
	repo vendoranalytics.Repository
}

func NewService(repo vendoranalytics.Repository) Service {
	return &ServiceImpl{
		repo: repo,
	}
}

func (s *ServiceImpl) GetSubscription(ctx context.Context, vendorID string) ([]models.Subscription, error) {
	return s.repo.GetSubscription(ctx, vendorID)
}

func (s *ServiceImpl) GetScheduledOrders(ctx context.Context, vendorID string) ([]models.ItemOrdered, error) {
	return s.repo.GetScheduledOrders(ctx, vendorID)
}

func (s *ServiceImpl) GetLiveOrder(ctx context.Context, vendorID string) ([]models.ItemOrdered, error) {
	return s.repo.GetLiveOrder(ctx, vendorID)
}

func (s *ServiceImpl) GetSalesByPeriod(ctx context.Context, vendorID string, period int) ([]models.Orders, error) {
	return s.repo.GetSalesByPeriod(ctx, vendorID, period)
}

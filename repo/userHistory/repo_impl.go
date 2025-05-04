package userhistory

import (
	"context"

	"github.com/VenzeneCorp/orderService/models"
	postgresdb "github.com/VenzeneCorp/orderService/repo/userHistory/postgresDB"
)

type UserHistoryRepo struct {
	sqlDB postgresdb.SQL
}

func NewUserHistoryRepo(db *postgresdb.SQL) Repository {
	return &UserHistoryRepo{
		sqlDB: *db,
	}
}

func (r *UserHistoryRepo) PlaceLiveOrder(ctx context.Context, order models.CreateOrder, liveOrder []models.CreateLiveOrder) error {
	return r.sqlDB.PlaceLiveOrder(ctx, order, liveOrder)
}

func (r *UserHistoryRepo) PlaceSubscriptionOrder(ctx context.Context, order models.CreateOrder, subscription models.CreateSubscription) error {
	return r.sqlDB.PlaceSubscriptionOrder(ctx, order, subscription)
}

func (r *UserHistoryRepo) CancelOrder(ctx context.Context, orderID string) error {

	return r.sqlDB.CancelOrder(ctx, orderID)
}

func (r *UserHistoryRepo) GetSubscriptionInfo(ctx context.Context, userID string) (models.Subscription, error) {
	return r.sqlDB.GetSubscriptionInfo(ctx, userID)
}

func (r *UserHistoryRepo) GetUserHistory(ctx context.Context, userID string) ([]models.ItemOrdered, error) {
	return r.sqlDB.GetUserHistory(ctx, userID)
}

func (r *UserHistoryRepo) GetUserSubscriptionHistory(ctx context.Context, userID string) ([]models.ItemOrdered, error) {
	return r.sqlDB.GetUserSubscriptionHistory(ctx, userID)
}

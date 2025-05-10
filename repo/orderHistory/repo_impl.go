package orderhistory

import (
	"context"

	"github.com/VenzeneCorp/orderService/models"
	postgresdb "github.com/VenzeneCorp/orderService/repo/orderHistory/postgresDB"
)

type UserHistoryRepo struct {
	sqlDB postgresdb.SQL
}

func NewUserHistoryRepo(db *postgresdb.SQL) Repository {
	return &UserHistoryRepo{
		sqlDB: *db,
	}
}

func (r *UserHistoryRepo) PlaceLiveOrder(ctx context.Context, userID string, order models.CreateOrder, liveOrder []models.CreateLiveOrder) (models.Orders, error) {
	orders, err := r.sqlDB.PlaceLiveOrder(ctx, userID, order, liveOrder)
	return orders, err
}

func (r *UserHistoryRepo) PlaceSubscriptionOrder(ctx context.Context, userID string, order models.CreateOrder, subscription models.CreateSubscription) (models.Orders, error) {
	orders, err := r.sqlDB.PlaceSubscriptionOrder(ctx, userID, order, subscription)
	return orders, err
}

func (r *UserHistoryRepo) CancelOrder(ctx context.Context, userID string, orderID string) error {

	return r.sqlDB.CancelOrder(ctx, userID, orderID)
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

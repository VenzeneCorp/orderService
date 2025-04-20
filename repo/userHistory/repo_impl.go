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

func (r *UserHistoryRepo) PlaceLiveOrder(ctx context.Context, order models.CreateOrder) error {
	return r.sqlDB.PlaceOrder(ctx, order)
}

func (r *UserHistoryRepo) GetSubscriptionInfo(ctx context.Context, userID string) (models.Subscribed, error) {
	return r.sqlDB.GetSubscriptionInfo(ctx, userID)
}

func (r *UserHistoryRepo) GetUserHistory(ctx context.Context, userID string) ([]models.OrderHistory, error) {
	return r.sqlDB.GetUserHistory(ctx, userID)
}

func (r *UserHistoryRepo) GetUserHistoryByOrderID(ctx context.Context, orderID string) ([]models.ItemOrdered, error) {
	return r.sqlDB.GetUserHistoryByOrderID(ctx, orderID)
}

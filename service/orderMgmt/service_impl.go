package ordermgmt

import (
	"context"
	"os"

	"github.com/VenzeneCorp/orderService/models"
	"github.com/VenzeneCorp/orderService/razorpayPayments"
	orderhistory "github.com/VenzeneCorp/orderService/repo/orderHistory"
	"github.com/VenzeneCorp/orderService/repo/payments"
)

type ServiceImpl struct {
	repo     orderhistory.Repository
	payments payments.Repository
}

func NewService(repo orderhistory.Repository) Service {
	return &ServiceImpl{
		repo: repo,
	}
}

func (s *ServiceImpl) PlaceLiveOrder(ctx context.Context, userID string, order models.CreateOrder, liveOrder []models.CreateLiveOrder) (models.RazorpayPaymentRespone, error) {
	var razporpayResponse models.RazorpayPaymentRespone
	orders, err := s.repo.PlaceLiveOrder(ctx, userID, order, liveOrder)
	if err != nil {
		return razporpayResponse, err
	}

	razorpayOrderID, err := razorpayPayments.CreateRazorpayOrder(order.Amount, "liveOrder")
	if err != nil {
		return razporpayResponse, err
	}

	_, err = s.payments.CreateRazorpayOrder(ctx, orders.ID, order.Amount, "liveOrder", razorpayOrderID)
	if err != nil {
		return razporpayResponse, err
	}

	razporpayResponse = models.RazorpayPaymentRespone{
		RazorpayOrderID: razorpayOrderID,
		Amount:          order.Amount,
		RAZORPAY_KEY_ID: os.Getenv("RAZORPAY_KEY_ID"),
	}

	return razporpayResponse, nil
}

func (s *ServiceImpl) UpdateOrderStatus(ctx context.Context, userID, razorpayOrderID string, razorpayResponse models.RazorpaySuccessRequest) error {
	err := s.payments.UpdateOrderStatus(ctx, razorpayOrderID, int(razorpayResponse.Status))
	if err != nil {
		return err
	}
	if razorpayResponse.Status == models.RazorpayStatus(models.Success) {
		return nil
	} else if razorpayResponse.Status == models.RazorpayStatus(models.Failed) {
		return s.repo.CancelOrder(ctx, userID, razorpayOrderID)
	}
	return nil
}

func (s *ServiceImpl) PlaceSubscriptionOrder(ctx context.Context, userID string, order models.CreateOrder, subscription models.CreateSubscription) (models.RazorpayPaymentRespone, error) {
	var razporpayResponse models.RazorpayPaymentRespone
	razorpayOrderID, err := razorpayPayments.CreateRazorpayOrder(order.Amount, "subscription")
	if err != nil {
		return razporpayResponse, err
	}

	orders, err := s.repo.PlaceSubscriptionOrder(ctx, userID, order, subscription)
	if err != nil {
		return razporpayResponse, err
	}

	_, err = s.payments.CreateRazorpayOrder(ctx, orders.ID, order.Amount, "subscription", razorpayOrderID)
	if err != nil {
		return razporpayResponse, err
	}

	razporpayResponse = models.RazorpayPaymentRespone{
		RazorpayOrderID: razorpayOrderID,
		Amount:          order.Amount,
		RAZORPAY_KEY_ID: os.Getenv("RAZORPAY_KEY_ID"),
	}

	return razporpayResponse, nil
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

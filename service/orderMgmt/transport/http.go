package transport

import (
	"encoding/json"
	"net/http"

	"github.com/VenzeneCorp/orderService/middlewares"
	"github.com/VenzeneCorp/orderService/models"
	ordermgmt "github.com/VenzeneCorp/orderService/service/orderMgmt"
	"github.com/gorilla/mux"
)

var RoleUser = "user"
var RoleVendor = "vendor"
var RoleAdmin = "admin"

type Handler struct {
	service ordermgmt.Service
}

func NewHandler(service ordermgmt.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.Use(middlewares.AuthMiddleware)

	router.HandleFunc("/orders/live", h.PlaceLiveOrder).Methods(http.MethodPost)
	router.HandleFunc("/orders/subscription", h.PlaceSubscriptionOrder).Methods(http.MethodPost)
	router.HandleFunc("/orders/cancel", h.CancelOrder).Methods(http.MethodPost)
	router.HandleFunc("/subscriptions", h.GetSubscriptionInfo).Methods(http.MethodGet)
	router.HandleFunc("/orders/history", h.GetUserHistory).Methods(http.MethodGet)
	router.HandleFunc("/subscriptions/history", h.GetUserSubscriptionHistory).Methods(http.MethodGet)
}
func (h *Handler) PlaceLiveOrder(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("X-ID")
	role := r.Header.Get("X-Role")

	if userID == "" || (role == "" || (role != RoleUser && role != RoleAdmin)) {
		http.Error(w, "Missing userID or role", http.StatusBadRequest)
		return
	}

	var req struct {
		Order     models.CreateOrder       `json:"order"`
		LiveOrder []models.CreateLiveOrder `json:"live_order"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.service.PlaceLiveOrder(r.Context(), userID, req.Order, req.LiveOrder); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) PlaceSubscriptionOrder(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("X-ID")
	role := r.Header.Get("X-Role")

	if userID == "" || (role == "" || (role != RoleUser && role != RoleAdmin)) {
		http.Error(w, "Missing userID or role", http.StatusBadRequest)
		return
	}

	var req struct {
		Order        models.CreateOrder        `json:"order"`
		Subscription models.CreateSubscription `json:"subscription"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.service.PlaceSubscriptionOrder(r.Context(), userID, req.Order, req.Subscription); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) CancelOrder(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("X-ID")
	role := r.Header.Get("X-Role")

	if userID == "" || (role == "" || (role != RoleUser && role != RoleAdmin)) {
		http.Error(w, "Missing userID or role", http.StatusBadRequest)
		return
	}

	var req struct {
		OrderID string `json:"order_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if req.OrderID == "" {
		http.Error(w, "Missing orderID", http.StatusBadRequest)
		return
	}

	if err := h.service.CancelOrder(r.Context(), userID, req.OrderID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetSubscriptionInfo(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("X-ID")
	role := r.Header.Get("X-Role")

	if userID == "" || (role == "" || (role != RoleUser && role != RoleAdmin)) {
		http.Error(w, "Missing userID or role", http.StatusBadRequest)
		return
	}

	subscription, err := h.service.GetSubscriptionInfo(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(subscription)
}

func (h *Handler) GetUserHistory(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("X-ID")
	role := r.Header.Get("X-Role")

	if userID == "" || (role == "" || (role != RoleUser && role != RoleAdmin)) {
		http.Error(w, "Missing userID or role", http.StatusBadRequest)
		return
	}

	history, err := h.service.GetUserHistory(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(history)
}

func (h *Handler) GetUserSubscriptionHistory(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("X-ID")
	role := r.Header.Get("X-Role")

	if userID == "" || (role == "" || (role != RoleUser && role != RoleAdmin)) {
		http.Error(w, "Missing userID or role", http.StatusBadRequest)
		return
	}

	history, err := h.service.GetUserSubscriptionHistory(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(history)
}

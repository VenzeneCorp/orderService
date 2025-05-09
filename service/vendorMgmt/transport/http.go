package transport

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/VenzeneCorp/orderService/middlewares"
	vendormgmt "github.com/VenzeneCorp/orderService/service/vendorMgmt"
	"github.com/gorilla/mux"
)

var RoleUser = "user"
var RoleVendor = "vendor"
var RoleAdmin = "admin"

type Handler struct {
	service vendormgmt.Service
}

func NewHandler(service vendormgmt.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.Use(middlewares.AuthMiddleware)

	router.HandleFunc("/vendor/subscriptions", h.GetSubscriptions).Methods(http.MethodGet)
	router.HandleFunc("/vendor/scheduled-orders", h.GetScheduledOrders).Methods(http.MethodGet)
	router.HandleFunc("/vendor/live-orders", h.GetLiveOrders).Methods(http.MethodGet)
	router.HandleFunc("/vendor/sales", h.GetSalesByPeriod).Methods(http.MethodGet)
}

func (h *Handler) GetSubscriptions(w http.ResponseWriter, r *http.Request) {
	vendorID := r.Header.Get("X-ID")
	role := r.Header.Get("X-Role")

	if vendorID == "" || (role == "" || (role != RoleVendor && role != RoleAdmin)) {
		http.Error(w, "Missing userID or role", http.StatusBadRequest)
		return
	}

	subscriptions, err := h.service.GetSubscription(r.Context(), vendorID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(subscriptions)
}

func (h *Handler) GetScheduledOrders(w http.ResponseWriter, r *http.Request) {
	vendorID := r.Header.Get("X-ID")
	role := r.Header.Get("X-Role")

	if vendorID == "" || (role == "" || (role != RoleVendor && role != RoleAdmin)) {
		http.Error(w, "Missing userID or role", http.StatusBadRequest)
		return
	}

	orders, err := h.service.GetScheduledOrders(r.Context(), vendorID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

func (h *Handler) GetLiveOrders(w http.ResponseWriter, r *http.Request) {
	vendorID := r.Header.Get("X-ID")
	role := r.Header.Get("X-Role")

	if vendorID == "" || (role == "" || (role != RoleVendor && role != RoleAdmin)) {
		http.Error(w, "Missing userID or role", http.StatusBadRequest)
		return
	}

	orders, err := h.service.GetLiveOrder(r.Context(), vendorID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

func (h *Handler) GetSalesByPeriod(w http.ResponseWriter, r *http.Request) {
	vendorID := r.Header.Get("X-ID")
	role := r.Header.Get("X-Role")

	if vendorID == "" || (role == "" || (role != RoleVendor && role != RoleAdmin)) {
		http.Error(w, "Missing userID or role", http.StatusBadRequest)
		return
	}

	periodStr := r.URL.Query().Get("period")
	period, err := strconv.Atoi(periodStr)
	if err != nil {
		http.Error(w, "Invalid period parameter", http.StatusBadRequest)
		return
	}

	sales, err := h.service.GetSalesByPeriod(r.Context(), vendorID, period)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sales)
}

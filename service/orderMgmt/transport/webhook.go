package transport

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/VenzeneCorp/orderService/models"
	"github.com/gorilla/mux"
)

func (h *Handler) RegisterWebhookRoutes(router *mux.Router) {
	webhookRouter := router.PathPrefix("/webhook").Subrouter()
	webhookRouter.HandleFunc("/razorpay", h.HandleWebhook).Methods(http.MethodPost)
}

// TODO: setup razorpay webhook on razorpay dashboard
func (h *Handler) HandleWebhook(w http.ResponseWriter, r *http.Request) {
	// Step 1: Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	// Step 2: Extract the signature from the request header
	signature := r.Header.Get("X-Razorpay-Signature")
	if signature == "" {
		http.Error(w, "Missing signature", http.StatusBadRequest)
		return
	}
	// Step 3: Parse the JSON payload
	var webhookPayload struct {
		Event   string `json:"event"`
		Payload struct {
			Payment struct {
				Entity struct {
					ID      string `json:"id"`
					OrderID string `json:"order_id"`
					Status  string `json:"status"`
				} `json:"entity"`
			} `json:"payment"`
		} `json:"payload"`
	}

	if err := json.Unmarshal(body, &webhookPayload); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	if webhookPayload.Event != "payment.captured" && webhookPayload.Event != "payment.failed" {
		http.Error(w, "Unsupported event type", http.StatusBadRequest)
		return
	}

	// Step 4: Construct RazorpaySuccessRequest
	var successRequest models.RazorpaySuccessRequest
	if webhookPayload.Payload.Payment.Entity.Status != "captured" {
		successRequest = models.RazorpaySuccessRequest{
			RazorpayOrderID:   webhookPayload.Payload.Payment.Entity.OrderID,
			RazorpayPaymentID: webhookPayload.Payload.Payment.Entity.ID,
			RazorpaySignature: signature,
			Status:            1,
		}
	} else {
		successRequest = models.RazorpaySuccessRequest{
			RazorpayOrderID:   webhookPayload.Payload.Payment.Entity.OrderID,
			RazorpayPaymentID: webhookPayload.Payload.Payment.Entity.ID,
			RazorpaySignature: signature,
			Status:            2,
		}
	}

	// Step 6: Update order status in the database
	if err := h.service.UpdateOrderStatus(r.Context(), "-1", successRequest.RazorpayOrderID, successRequest); err != nil {
		http.Error(w, "Failed to update order status", http.StatusInternalServerError)
		return
	}

	// Step 7: Send success response to Razorpay
	w.WriteHeader(http.StatusOK)
}

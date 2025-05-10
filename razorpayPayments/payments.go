package razorpayPayments

import (
	"os"

	"github.com/razorpay/razorpay-go"
)

func CreateRazorpayOrder(amount int, receipt string) (string, error) {
	client := razorpay.NewClient(
		os.Getenv("RAZORPAY_KEY_ID"),
		os.Getenv("RAZORPAY_SECRET"),
	)

	data := map[string]interface{}{
		"amount":   amount, // in paise
		"currency": "INR",
		"receipt":  receipt,
	}

	body, err := client.Order.Create(data, nil)
	if err != nil {
		return "", err
	}

	orderID := body["id"].(string)
	return orderID, nil
}

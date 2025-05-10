package razorpayPayments

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
)

func VerifyPaymentSignature(orderID, paymentID, signature string) error {
	secret := os.Getenv("RAZORPAY_SECRET")
	data := orderID + "|" + paymentID

	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	computed := hex.EncodeToString(h.Sum(nil))

	if computed != signature {
		return fmt.Errorf("invalid signature")
	}
	return nil
}

package loginsignupclient

import (
	"context"

	loginsignupmodels "github.com/venzene/loginSignup/models"
)

type LoginSignupClient interface {
	GetUserAddresses(ctx context.Context, userID int) ([]loginsignupmodels.Address, error)
}

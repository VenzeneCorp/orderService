package loginsignupclient

import (
	"context"

	loginsignupmodels "github.com/VenzeneCorp/loginSignup/models"
)

type LoginSignupClient interface {
	GetUserAddresses(ctx context.Context, userID int) ([]loginsignupmodels.Address, error)
}

package requests

import (
	"github.com/guregu/null"
	"github.com/storiqateam/payments/app/models"
)

type SessionCreateInput struct {
	Email      string
	Password   string
	DeviceType models.DeviceType
	DeviceOs   null.String
	DeviceId   null.String
}

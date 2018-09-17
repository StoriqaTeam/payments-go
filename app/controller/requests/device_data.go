package requests

import (
	"github.com/guregu/null"
	"github.com/storiqateam/payments/app/models"
)

type DeviceData struct {
	DeviceType models.DeviceType `validate:"required"`
	DeviceOs   null.String
	DeviceId   null.String
}

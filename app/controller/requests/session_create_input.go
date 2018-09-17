package requests

import (
	"database/sql"

	"github.com/storiqateam/payments/app/models"
)

type SessionCreateInput struct {
	Email      string
	Password   string
	DeviceType models.DeviceType
	DeviceOs   sql.NullString
	DeviceId   sql.NullString
}

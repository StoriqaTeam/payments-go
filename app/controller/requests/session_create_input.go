package requests

import (
	"database/sql"

	"github.com/storiqateam/payments/app/models"
)

type SessionCreateInput struct {
	email      string
	password   string
	deviceType models.DeviceType
	deviceOs   sql.NullString
	deviceId   sql.NullString
}

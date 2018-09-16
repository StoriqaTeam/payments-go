package services

import (
	"database/sql"

	"github.com/storiqateam/payments/app/models"
)

// AuthService provides all necessary auth methods
type AuthService struct{}

func (as *AuthService) loginByEmail(email string, password string, deviceType models.DeviceType, deviceOs sql.NullString, deviceId sql.NullString) {

}

package requests

import "github.com/storiqateam/app/models"

type SessionCreateInput struct {
	email      string
	password   string
	deviceType models.DeviceType
}

package requests

type SessionCreateInput struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
	DeviceData
}

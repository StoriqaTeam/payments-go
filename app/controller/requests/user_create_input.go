package requests

type UserCreateInput struct {
	Email     string `validate:"required"`
	Password  string `validate:"required"`
	Phone     string `validate:"required"`
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	DeviceData
}

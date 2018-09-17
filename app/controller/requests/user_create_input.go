package requests

type UserCreateInput struct {
	Email     string
	Password  string
	Phone     string
	FirstName string
	LastName  string
	DeviceData
}

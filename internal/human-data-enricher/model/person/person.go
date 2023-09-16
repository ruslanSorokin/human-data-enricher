package person

type Person struct {
	Name       string `validate:"required,alpha"`
	Surname    string `validate:"required,alpha"`
	MiddleName string `validate:"omitempty,alpha"`

	Nationality string `validate:""`
	Sex         string `validate:""`
	Age         uint   `validate:""`
}

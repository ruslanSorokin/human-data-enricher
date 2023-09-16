package person

type Person struct {
	Name       string `validate:"required,alpha"`
	Surname    string `validate:"required,alpha"`
	MiddleName string `validate:"omitempty,alpha"`

	Nationality string `validate:"omitempty"`
	Sex         string `validate:"omitempty"`
	Age         uint   `validate:"omitempty"`
}

func New(n, sn, mn string) *Person{
	return &Person{
		Name:        n,
		Surname:     sn,
		MiddleName:  mn,
	}
}

package model

import (
	"database/sql"
	"time"

	uuid "github.com/gofrs/uuid/v5"
)

type PersonID uuid.UUID

// Person represents a person entity.
//
//nolint:govet // Alignment is pretty much optimized
type Person struct {
	id uuid.UUID

	createdAt time.Time
	updatedAt time.Time
	deletedAt sql.NullTime

	Name       string
	Surname    string
	MiddleName sql.NullString

	Gender      string `validate:"required,alpha"`
	Nationality string `validate:"required,alpha"`
	Age         int    `validate:"required,gte=0,lte=130"`
}

// PersonOptions contains the input data for creating the Person struct.
//
//nolint:govet // Alignment is pretty much optimized
type PersonOptions struct {
	Name       string
	Surname    string
	MiddleName sql.NullString

	Gender      string
	Nationality string
	Age         int
}

type ReinstatedPersonOpts struct {
	ID uuid.UUID

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime

	Name       string
	Surname    string
	MiddleName sql.NullString

	Gender      string
	Nationality string
	Age         int
}

func (o *PersonOptions) toModel() (Person, error) {
	var res Person
	id, err := uuid.NewV7()
	if err != nil {
		return res, err
	}

	t := time.Now().UTC()

	res.id = id
	res.createdAt = t
	res.updatedAt = t
	res.deletedAt = sql.NullTime{Valid: false, Time: t}
	res.Name = o.Name
	res.Surname = o.Surname
	res.MiddleName = o.MiddleName
	res.Gender = o.Gender
	res.Nationality = o.Nationality
	res.Age = o.Age

	return res, nil
}

func (p *Person) ID() uuid.UUID {
	return p.id
}

func (p *Person) CreatedAt() time.Time {
	return p.createdAt
}

func (p *Person) UpdatedAt() time.Time {
	return p.updatedAt
}

func (p *Person) DeletedAt() sql.NullTime {
	return p.deletedAt
}

// NewPerson creates a new Person with given "opts".
func NewPerson(
	opts *PersonOptions,
) (Person, error) {
	return opts.toModel()
}

func ReinstatePerson(
	opts *ReinstatedPersonOpts,
) Person {
	return Person{
		id:          opts.ID,
		Name:        opts.Name,
		Surname:     opts.Surname,
		MiddleName:  opts.MiddleName,
		Nationality: opts.Nationality,
		Gender:      opts.Gender,
		Age:         opts.Age,
		createdAt:   opts.CreatedAt,
		updatedAt:   opts.UpdatedAt,
		deletedAt:   opts.DeletedAt,
	}
}

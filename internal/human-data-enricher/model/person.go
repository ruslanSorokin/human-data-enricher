package model

import (
	"database/sql"
	"time"

	uuid "github.com/gofrs/uuid/v5"
)

type PersonID uuid.UUID

// Person represents a person entity.
//
//nolint:govet // alignment is pretty much optimized
type Person struct {
	ID uuid.UUID

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime

	Name       string  `validate:"required,alpha"`
	Surname    string  `validate:"required,alpha"`
	MiddleName *string `validate:"omitempty,alpha"`

	Gender      string `validate:"required,alpha"`
	Nationality string `validate:"required,alpha"`
	Age         int    `validate:"required,gte=0,lte=130"`
}

// PersonOptions contains the input data for creating the Person struct.
type PersonOptions struct {
	Name       string
	Surname    string
	MiddleName *string

	Gender      string
	Nationality string
	Age         int
}

func (opts *PersonOptions) toModel() Person {
	return Person{
		ID:          uuid.Nil,
		Name:        opts.Name,
		Surname:     opts.Surname,
		MiddleName:  opts.MiddleName,
		Nationality: opts.Nationality,
		Gender:      opts.Gender,
		Age:         opts.Age,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   sql.NullTime{Time: time.Time{}, Valid: false},
	}
}

// NewPerson creates a new Person with given "opts".
func NewPerson(
	opts *PersonOptions,
) Person {
	return opts.toModel()
}

func ReinstatePerson(
	id uuid.UUID,
	name string,
	surname string,
	middleName *string,
	nationality string,
	gender string,
	age int,
	createdAt time.Time,
	updatedAt time.Time,
	deletedAt sql.NullTime,
) Person {
	return Person{
		ID:          id,
		Name:        name,
		Surname:     surname,
		MiddleName:  middleName,
		Nationality: nationality,
		Gender:      gender,
		Age:         age,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
		DeletedAt:   deletedAt,
	}
}

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

func (opts *PersonOptions) toModel() Person {
	return Person{
		id:          uuid.Nil,
		Name:        opts.Name,
		Surname:     opts.Surname,
		MiddleName:  opts.MiddleName,
		Nationality: opts.Nationality,
		Gender:      opts.Gender,
		Age:         opts.Age,
		createdAt:   time.Time{},
		updatedAt:   time.Time{},
		deletedAt:   sql.NullTime{Time: time.Time{}, Valid: false},
	}
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
) Person {
	return opts.toModel()
}

func ReinstatePerson(
	id uuid.UUID,
	name string,
	surname string,
	middleName sql.NullString,
	nationality string,
	gender string,
	age int,
	createdAt time.Time,
	updatedAt time.Time,
	deletedAt sql.NullTime,
) Person {
	return Person{
		id:          id,
		Name:        name,
		Surname:     surname,
		MiddleName:  middleName,
		Nationality: nationality,
		Gender:      gender,
		Age:         age,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
		deletedAt:   deletedAt,
	}
}

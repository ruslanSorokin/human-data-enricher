package validation

import (
	"context"
	"database/sql"

	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/model"
)

type ValidatorI interface {
	Person(ctx context.Context, person *model.Person) error
	ID(ctx context.Context, id model.PersonID) error
	Name(ctx context.Context, name string) error
	Surname(ctx context.Context, surname string) error
	MiddleName(ctx context.Context, middlename sql.NullString) error
	Gender(ctx context.Context, gender string) error
	Nationality(ctx context.Context, nationality string) error
	Age(ctx context.Context, age int) error
}

// Code generated by ent, DO NOT EDIT.

package gen

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/provider/storage/ient/gen/person"
	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/provider/storage/ient/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	personFields := schema.Person{}.Fields()
	_ = personFields
	// personDescCreatedAt is the schema descriptor for created_at field.
	personDescCreatedAt := personFields[1].Descriptor()
	// person.DefaultCreatedAt holds the default value on creation for the created_at field.
	person.DefaultCreatedAt = personDescCreatedAt.Default.(func() time.Time)
	// personDescUpdatedAt is the schema descriptor for updated_at field.
	personDescUpdatedAt := personFields[2].Descriptor()
	// person.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	person.DefaultUpdatedAt = personDescUpdatedAt.Default.(func() time.Time)
	// person.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	person.UpdateDefaultUpdatedAt = personDescUpdatedAt.UpdateDefault.(func() time.Time)
	// personDescDeletedAt is the schema descriptor for deleted_at field.
	personDescDeletedAt := personFields[3].Descriptor()
	// person.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	person.DefaultDeletedAt = personDescDeletedAt.Default.(func() sql.NullTime)
	// personDescName is the schema descriptor for name field.
	personDescName := personFields[4].Descriptor()
	// person.NameValidator is a validator for the "name" field. It is called by the builders before save.
	person.NameValidator = personDescName.Validators[0].(func(string) error)
	// personDescSurname is the schema descriptor for surname field.
	personDescSurname := personFields[5].Descriptor()
	// person.SurnameValidator is a validator for the "surname" field. It is called by the builders before save.
	person.SurnameValidator = personDescSurname.Validators[0].(func(string) error)
	// personDescMiddleName is the schema descriptor for middle_name field.
	personDescMiddleName := personFields[6].Descriptor()
	// person.DefaultMiddleName holds the default value on creation for the middle_name field.
	person.DefaultMiddleName = personDescMiddleName.Default.(func() sql.NullString)
	// personDescGender is the schema descriptor for gender field.
	personDescGender := personFields[8].Descriptor()
	// person.GenderValidator is a validator for the "gender" field. It is called by the builders before save.
	person.GenderValidator = personDescGender.Validators[0].(func(string) error)
	// personDescNationality is the schema descriptor for nationality field.
	personDescNationality := personFields[9].Descriptor()
	// person.NationalityValidator is a validator for the "nationality" field. It is called by the builders before save.
	person.NationalityValidator = personDescNationality.Validators[0].(func(string) error)
}

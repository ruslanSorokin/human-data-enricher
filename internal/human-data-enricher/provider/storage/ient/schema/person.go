package schema

import (
	"database/sql"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	uuid "github.com/gofrs/uuid/v5"
)

// Person holds the schema definition for the Person entity.
type Person struct {
	ent.Schema
}

// Fields of the Person.
func (Person) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Unique().
			Immutable(),

		field.Time("created_at").
			Immutable().
			Optional().
			Default(time.Now),
		field.Time("updated_at").
			Optional().
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Time("deleted_at").
			Optional().
			Default(func() sql.NullTime {
				return sql.NullTime{Valid: false}
			}).
			GoType(sql.NullTime{}),

		field.String("name").
			NotEmpty(),
		field.String("surname").
			NotEmpty(),
		field.String("middle_name").
			Optional().
			DefaultFunc(func() sql.NullString {
				return sql.NullString{Valid: false}
			}).
			GoType(sql.NullString{}),

		field.Int("age"),
		field.String("gender").
			NotEmpty(),
		field.String("nationality").
			NotEmpty(),
	}
}

// Edges of the Person.
func (Person) Edges() []ent.Edge {
	return nil
}

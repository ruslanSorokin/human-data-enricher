package test_person

import (
	"database/sql"

	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/model"
	"github.com/ruslanSorokin/human-data-enricher/internal/pkg/util"
)

func aPerson() model.Person {
	return util.MustXX(model.NewPerson(&model.PersonOptions{
		Name:        "Ivan",
		Surname:     "Ivanov",
		MiddleName:  sql.NullString{Valid: false, String: ""},
		Gender:      "MALE",
		Nationality: "RU",
		Age:         21,
	}))
}

func aListOfPersons() []*model.Person {
	return []*model.Person{
		func() *model.Person {
			m := util.MustXX(model.NewPerson(&model.PersonOptions{
				Name:        "Dmitry",
				Surname:     "Ivanov",
				MiddleName:  sql.NullString{Valid: false, String: ""},
				Gender:      "M",
				Nationality: "RU",
				Age:         21,
			}))
			return &m
		}(),
		func() *model.Person {
			m := util.MustXX(model.NewPerson(&model.PersonOptions{
				Name:        "Ivan",
				Surname:     "Kashin",
				MiddleName:  sql.NullString{String: "Dmitrievich", Valid: true},
				Gender:      "MALE",
				Nationality: "RU",
				Age:         26,
			}))
			return &m
		}(),
		func() *model.Person {
			m := util.MustXX(model.NewPerson(&model.PersonOptions{
				Name:        "Kirill",
				Surname:     "Federov",
				MiddleName:  sql.NullString{String: "Ivanov", Valid: true},
				Gender:      "MALE",
				Nationality: "RU",
				Age:         31,
			}))
			return &m
		}(),
		func() *model.Person {
			m := util.MustXX(model.NewPerson(&model.PersonOptions{
				Name:        "Maxim",
				Surname:     "Korablev",
				MiddleName:  sql.NullString{Valid: false, String: ""},
				Gender:      "M",
				Nationality: "UA",
				Age:         24,
			}))
			return &m
		}(),
		func() *model.Person {
			m := util.MustXX(model.NewPerson(&model.PersonOptions{
				Name:        "Diana",
				Surname:     "Orlova",
				MiddleName:  sql.NullString{Valid: false, String: ""},
				Gender:      "F",
				Nationality: "UA",
				Age:         20,
			}))
			return &m
		}(),
		func() *model.Person {
			m := util.MustXX(model.NewPerson(&model.PersonOptions{
				Name:        "Katya",
				Surname:     "Lebedeva",
				MiddleName:  sql.NullString{Valid: false, String: ""},
				Gender:      "FEMALE",
				Nationality: "RU",
				Age:         19,
			}))
			return &m
		}(),
	}
}

package person_test

import (
	"log/slog"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/require"

	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/model/person"
)

func TestValidatorRequired(t *testing.T) {
	log := slog.Default()
	v := person.NewValidator(
		log,
		validator.New(validator.WithRequiredStructEnabled()),
	)
	require := require.New(t)

	type got struct {
		p *person.Person
	}
	type want struct {
		err error
	}

	testCases := []struct {
		desc string
		got  got
		want want
	}{
		{
			desc: "Missing Name",
			got: got{
				p: &person.Person{
					Name:        "",
					Surname:     "Ivan",
					MiddleName:  "Invanov",
					Age:         22,
					Sex:         "male",
					Nationality: "RU",
				},
			},
			want: want{
				err: person.ErrMissingName,
			},
		},
		{
			desc: "Missing Surname",
			got: got{
				p: &person.Person{
					Name:        "Dmitry",
					Surname:     "",
					MiddleName:  "Ivanovich",
					Age:         21,
					Sex:         "m",
					Nationality: "Russian",
				},
			},
			want: want{
				err: person.ErrMissingSurname,
			},
		},
		{
			desc: "Without MiddleName",
			got: got{
				p: &person.Person{
					Name:        "Dmitry",
					Surname:     "Ivanov",
					MiddleName:  "",
					Age:         21,
					Sex:         "m",
					Nationality: "Russian",
				},
			},
			want: want{
				err: nil,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			err := v(tc.got.p)
			require.ErrorIs(err, tc.want.err)
		})
	}
}

func TestValidatorInvalid(t *testing.T) {
	log := slog.Default()
	v := person.NewValidator(log,
		validator.New(validator.WithRequiredStructEnabled()),
	)
	require := require.New(t)

	type got struct {
		p *person.Person
	}
	type want struct {
		err error
	}

	testCases := []struct {
		desc string
		got  got
		want want
	}{
		{
			desc: "Invalid Name",
			got: got{
				p: &person.Person{
					Name:        "Ivan123",
					Surname:     "Ivanov",
					MiddleName:  "Ivanovich",
					Age:         21,
					Sex:         "m",
					Nationality: "ru",
				},
			},
			want: want{
				err: person.ErrInvalidName,
			},
		},
		{
			desc: "Invalid Surname",
			got: got{
				p: &person.Person{
					Name:        "Dmitry",
					Surname:     "Ivan0v",
					MiddleName:  "Ivanovich",
					Age:         46,
					Sex:         "male",
					Nationality: "RU",
				},
			},
			want: want{
				err: person.ErrInvalidSurname,
			},
		},
		{
			desc: "Invalid MiddleName",
			got: got{
				p: &person.Person{
					Name:        "Dmitry",
					Surname:     "Ivanov",
					MiddleName:  "III123",
					Age:         0,
					Sex:         "",
					Nationality: "",
				},
			},
			want: want{
				err: person.ErrInvalidMiddleName,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			err := v(tc.got.p)
			require.ErrorIs(err, tc.want.err)
		})
	}
}

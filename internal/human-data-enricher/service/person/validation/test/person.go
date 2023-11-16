//nolint:exhaustruct // It's ok here
package validationtest

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/model"
	person_service "github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/service/person"
	"github.com/ruslanSorokin/human-data-enricher/internal/pkg/ierror"
)

func (s VSuite) TestPersonNoViolations() {
	t := s.T()
	require := require.New(t)
	ctx := context.Background()

	type got struct {
		p *model.Person
	}

	testCases := []struct {
		desc string
		got  got
	}{
		{
			desc: "With MiddleName",
			got: got{
				p: &model.Person{
					Name:        "Dmitry",
					Surname:     "Ivanov",
					MiddleName:  sql.NullString{String: "Ivanovich", Valid: true},
					Age:         21,
					Gender:      "MALE",
					Nationality: "RU",
				},
			},
		},
		{
			desc: "Without MiddleName",
			got: got{
				p: &model.Person{
					Name:        "Dmitry",
					Surname:     "Ivanov",
					MiddleName:  sql.NullString{},
					Age:         21,
					Gender:      "M",
					Nationality: "UA",
				},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			tc := tc
			err := s.vtor.Person(ctx, tc.got.p)
			t.Log("got error:", err)

			var wantErr error // nil error
			require.NoErrorf(err,
				"must return %#v, returned %#v",
				wantErr, err)
		})
	}
}

func (s VSuite) TestPersonMissingPropertyError() {
	t := s.T()
	assert := assert.New(t)
	require := require.New(t)
	ctx := context.Background()

	type got struct{ p *model.Person }
	type want struct{ err error }

	testCases := []struct {
		desc string
		got  got
		want want
	}{
		{
			desc: "Missing Name",
			got: got{
				p: &model.Person{
					Name:        "",
					Surname:     "Ivan",
					MiddleName:  sql.NullString{},
					Age:         21,
					Gender:      "M",
					Nationality: "UA",
				},
			},
			want: want{
				err: person_service.ErrMissingName,
			},
		},
		{
			desc: "Missing Surname",
			got: got{
				p: &model.Person{
					Name:        "Dmitry",
					Surname:     "",
					MiddleName:  sql.NullString{},
					Age:         21,
					Gender:      "M",
					Nationality: "UA",
				},
			},
			want: want{
				err: person_service.ErrMissingSurname,
			},
		},
		{
			desc: "Missing Age",
			got: got{
				p: &model.Person{
					Name:        "Dmitry",
					Surname:     "Ivanov",
					MiddleName:  sql.NullString{},
					Gender:      "M",
					Nationality: "UA",
				},
			},
			want: want{
				err: person_service.ErrMissingAge,
			},
		},
		{
			desc: "Missing Gender",
			got: got{
				p: &model.Person{
					Name:        "Dmitry",
					Surname:     "Ivanov",
					MiddleName:  sql.NullString{},
					Age:         21,
					Nationality: "UA",
				},
			},
			want: want{
				err: person_service.ErrMissingGender,
			},
		},
		{
			desc: "Missing Nationality",
			got: got{
				p: &model.Person{
					Name:       "Dmitry",
					Surname:    "Ivanov",
					MiddleName: sql.NullString{},
					Age:        21,
					Gender:     "M",
				},
			},
			want: want{
				err: person_service.ErrMissingNationality,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			tc := tc
			err := s.vtor.Person(ctx, tc.got.p)
			t.Log("want error:", tc.want.err)
			t.Log("got error:", err)

			errType := ierror.NewMissingProperty(
				"", "") // .Error() must include: "missing property: "
			assert.Contains(errType.Error(), "missing property: ")

			require.ErrorContainsf(err, errType.Error(),
				"must be subtype of %#v", errType)

			require.ErrorAsf(err, &tc.want.err,
				"must return %#v, returned %#v",
				tc.want.err, err)
		})
	}
}

func (s VSuite) TestPersonInvalidPropertyError() {
	t := s.T()
	assert := assert.New(t)
	require := require.New(t)
	ctx := context.Background()

	type got struct{ p *model.Person }
	type want struct{ err error }

	testCases := []struct {
		desc string
		got  got
		want want
	}{
		{
			desc: "Invalid Name",
			got: got{
				p: &model.Person{
					Name:        "Ivan123",
					Surname:     "Ivanov",
					MiddleName:  sql.NullString{},
					Age:         21,
					Gender:      "MALE",
					Nationality: "RU",
				},
			},
			want: want{
				err: person_service.ErrInvalidName,
			},
		},
		{
			desc: "Invalid Surname",
			got: got{
				p: &model.Person{
					Name:        "Dmitry",
					Surname:     "Ivan0v",
					MiddleName:  sql.NullString{},
					Age:         21,
					Gender:      "MALE",
					Nationality: "RU",
				},
			},
			want: want{
				err: person_service.ErrInvalidSurname,
			},
		},
		{
			desc: "Invalid MiddleName",
			got: got{
				p: &model.Person{
					Name:        "Dmitry",
					Surname:     "Ivanov",
					MiddleName:  sql.NullString{String: "Ivan0vich", Valid: true},
					Age:         21,
					Gender:      "MALE",
					Nationality: "RU",
				},
			},
			want: want{
				err: person_service.ErrInvalidMiddleName,
			},
		},
		{
			desc: "Invalid Age",
			got: got{
				p: &model.Person{
					Name:        "Dmitry",
					Surname:     "Ivanov",
					MiddleName:  sql.NullString{},
					Age:         -1,
					Gender:      "MALE",
					Nationality: "RU",
				},
			},
			want: want{
				err: person_service.ErrInvalidAge,
			},
		},
		{
			desc: "Invalid Gender",
			got: got{
				p: &model.Person{
					Name:        "Dmitry",
					Surname:     "Ivanov",
					MiddleName:  sql.NullString{},
					Age:         21,
					Gender:      "MALE2",
					Nationality: "RU",
				},
			},
			want: want{
				err: person_service.ErrInvalidGender,
			},
		},
		{
			desc: "Invalid Nationality",
			got: got{
				p: &model.Person{
					Name:        "Dmitry",
					Surname:     "Ivanov",
					MiddleName:  sql.NullString{},
					Age:         21,
					Gender:      "MALE",
					Nationality: "RU52",
				},
			},
			want: want{
				err: person_service.ErrInvalidNationality,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			tc := tc
			err := s.vtor.Person(ctx, tc.got.p)
			t.Log("want error:", tc.want.err)
			t.Log("got error:", err)

			errType := ierror.NewInvalidProperty(
				"", "") // .Error() must include: "invalid property: "
			assert.Contains(errType.Error(), "invalid property: ")

			require.ErrorContainsf(err, errType.Error(),
				"must be subtype of %#v", errType)

			require.ErrorAsf(err, &tc.want.err,
				"must return %#v, returned %#v",
				tc.want.err, err)
		})
	}
}

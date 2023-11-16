package model

import "github.com/ruslanSorokin/human-data-enricher/internal/pkg/ierror"

const (
	_name        = "name"
	_sname       = "surname"
	_mname       = "middle name"
	_age         = "age"
	_gender      = "gender"
	_nationality = "nationality"
)

var (
	// ErrMissingName occurs if the mandatory field "Name" is missing.
	ErrMissingName = ierror.NewMissingProperty(
		_name, "MISSING_NAME")

	// ErrMissingSurname occurs if the mandatory field "Surname" is missing.
	ErrMissingSurname = ierror.NewMissingProperty(
		_sname, "MISSING_SURNAME")

	// ErrMissingAge occurs if the mandatory field "Age" is missing.
	ErrMissingAge = ierror.NewMissingProperty(
		_age, "MISSING_AGE")

	// ErrMissingGender occurs if the mandatory field "Gender" is missing.
	ErrMissingGender = ierror.NewMissingProperty(
		_gender, "MISSING_GENDER")

	// ErrMissingNationality occurs if the mandatory field "Nationality" is missing.
	ErrMissingNationality = ierror.NewMissingProperty(
		_nationality, "MISSING_NATIONALITY")

	// ErrInvalidName occurs if the "Name" field consists of something other than
	// alphabetic characters.
	ErrInvalidName = ierror.NewInvalidProperty(
		_name, "INVALID_NAME")

	// ErrInvalidSurname occurs if the "Surname" field consists of something other
	// than alphabetic characters.
	ErrInvalidSurname = ierror.NewInvalidProperty(
		_sname, "INVALID_SURNAME")

	// ErrInvalidMiddleName occurs if the "MiddleName" field consists of something
	// other than alphabetic characters.
	ErrInvalidMiddleName = ierror.NewInvalidProperty(
		_mname, "INVALID_MIDDLE_NAME")

	// ErrInvalidAge occurs if the "Age" field is less than 0 or greater than 130.
	ErrInvalidAge = ierror.NewInvalidProperty(
		_age, "INVALID_AGE")

	// ErrInvalidGender occurs if the "Gender" field consists of something
	// other than alphabetic characters.
	ErrInvalidGender = ierror.NewInvalidProperty(
		_gender, "INVALID_GENDER")

	// ErrInvalidMiddleName occurs if the "MiddleName" field consists of something
	// other than alphabetic characters.
	ErrInvalidNationality = ierror.NewInvalidProperty(
		_nationality, "INVALID_NATIONALITY")
)

package person

import (
	"github.com/ruslanSorokin/human-data-enricher/internal/pkg/ierror"
)

const (
	_name  = "name"
	_sname = "surname"
	_mname = "middle name"
)

var (
	ErrMissingName    = ierror.NewMissingProperty(_name)
	ErrMissingSurname = ierror.NewMissingProperty(_sname)

	ErrInvalidName       = ierror.NewInvalidProperty(_name)
	ErrInvalidSurname    = ierror.NewInvalidProperty(_sname)
	ErrInvalidMiddleName = ierror.NewInvalidProperty(_mname)
)

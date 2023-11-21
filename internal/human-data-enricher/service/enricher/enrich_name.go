package person_enricher

import (
	"context"
	"database/sql"

	"github.com/sourcegraph/conc/pool"

	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/model"
)

type ByNameEnricherI interface {
	EnrichByName(
		ctx context.Context,
		name, sname string,
		mname sql.NullString,
	) (model.Person, error)
}

func (e *Enricher) EnrichByName(
	ctx context.Context,
	name, sname string,
	mname sql.NullString,
) (model.Person, error) {
	var props presumedProperties

	ageCh := make(chan int, 1)
	genderCh := make(chan string, 1)
	nationalityCh := make(chan string, 1)

	pool := pool.New().
		WithContext(ctx).
		WithCancelOnError()

	pool.Go(func(ctx context.Context) error {
		r, err := e.ageProvider.AgeByName(ctx, name)
		ageCh <- r
		close(ageCh)
		return err
	})

	pool.Go(func(ctx context.Context) error {
		r, err := e.genderProvider.GenderByName(ctx, name)
		genderCh <- r
		close(genderCh)
		return err
	})

	pool.Go(func(ctx context.Context) error {
		r, err := e.nationalityProvider.NationalityByName(ctx, name)
		nationalityCh <- r
		close(nationalityCh)
		return err
	})

	if err := pool.Wait(); err != nil {
		return model.Person{}, err
	}

	props.Nationality = <-nationalityCh
	props.Gender = <-genderCh
	props.Age = <-ageCh

	opts := model.PersonOptions{
		Name:        name,
		Surname:     sname,
		MiddleName:  mname,
		Gender:      props.Gender,
		Nationality: props.Nationality,
		Age:         props.Age,
	}

	return e.personCreator.Create(ctx, &opts)
}

type presumedProperties struct {
	Gender      string
	Nationality string
	Age         int
}

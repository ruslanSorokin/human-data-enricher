package ient

import (
	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/model"
	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/provider/storage/ient/gen"
)

func toModel(p *gen.Person) model.Person {
	if p == nil {
		panic("p can't be nil")
	}

	res := model.ReinstatePerson(&model.ReinstatedPersonOpts{
		ID:          p.ID,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
		DeletedAt:   p.DeletedAt,
		Name:        p.Name,
		Surname:     p.Surname,
		MiddleName:  p.MiddleName,
		Gender:      p.Gender,
		Nationality: p.Nationality,
		Age:         p.Age,
	})

	return res
}

func toModelP(p *gen.Person) *model.Person {
	res := toModel(p)
	return &res
}

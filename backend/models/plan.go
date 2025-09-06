package models

import "github.com/google/uuid"

type Plan struct {
	ID        uuid.UUID `json:"id" bun:"id,pk"`
	VacancyID uuid.UUID `json:"vacancy-id"`
	Vacancy   *Vacancy  `json:"vacancy" bun:"rel:has-one,join:vacancy_id=id"`
	Plan      string    `json:"plan"`
}

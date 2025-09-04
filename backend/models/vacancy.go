package models

import (
	"github.com/google/uuid"
)

type Vacancy struct {
	ID                    uuid.UUID `json:"id" bun:"id,pk"`
	Status                string    `json:"status"`
	Name                  string    `json:"name"`
	Region                string    `json:"region"`
	City                  string    `json:"city"`
	Address               string    `json:"address"`
	WorkType              string    `json:"work_type"`
	ContractType          string    `json:"contract_type"`
	EmploymentType        string    `json:"employment_type"`
	WorkSchedule          string    `json:"work_schedule"`
	Income                *int      `json:"income,omitempty"`
	SalaryMax             *int      `json:"salary_max,omitempty"`
	SalaryMin             *int      `json:"salary_min,omitempty"`
	Pronounces            string    `json:"pronounces"`
	Responsibilities      string    `json:"responsibilities"`
	Requirements          string    `json:"requirements"`
	Education             string    `json:"education"`
	Experience            string    `json:"experience"`
	SpecialPrograms       *bool     `json:"special_programs,omitempty"`
	ComputerSkills        *bool     `json:"computer_skills,omitempty"`
	ForeignLanguages      *bool     `json:"foreign_languages,omitempty"`
	LanguageLevel         *string   `json:"language_level,omitempty"`
	HasBusinessTrips      *bool     `json:"has_business_trips,omitempty"`
	AdditionalInformation *string   `json:"additional_information,omitempty"`
}

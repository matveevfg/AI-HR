package requests

import (
	"github.com/google/uuid"

	"github.com/matveevfg/AI-HR/backend/models"
)

type Vacancy struct {
	ID                    uuid.UUID `param:"id"`
	Status                string    `json:"status"`
	Name                  string    `json:"name"`
	Region                string    `json:"region"`
	City                  string    `json:"city"`
	Address               string    `json:"address"`
	WorkType              string    `json:"work_type"`
	ContractType          string    `json:"contract_type"`
	EmploymentType        string    `json:"employment_type"`
	WorkSchedule          string    `json:"work_schedule"`
	Income                *int      `json:"income"`
	SalaryMax             *int      `json:"salary_max"`
	SalaryMin             *int      `json:"salary_min"`
	Pronounces            string    `json:"pronounces"`
	Responsibilities      string    `json:"responsibilities"`
	Requirements          string    `json:"requirements"`
	Education             string    `json:"education"`
	Experience            string    `json:"experience"`
	SpecialPrograms       *bool     `json:"special_programs"`
	ComputerSkills        *bool     `json:"computer_skills"`
	ForeignLanguages      *bool     `json:"foreign_languages"`
	LanguageLevel         *string   `json:"language_level"`
	HasBusinessTrips      *bool     `json:"has_business_trips"`
	AdditionalInformation *string   `json:"additional_information"`
}

func (v *Vacancy) ToModel() *models.Vacancy {
	return &models.Vacancy{
		ID:                    v.ID,
		Status:                v.Status,
		Name:                  v.Name,
		Region:                v.Region,
		City:                  v.City,
		Address:               v.Address,
		ContractType:          v.ContractType,
		WorkType:              v.WorkType,
		EmploymentType:        v.EmploymentType,
		WorkSchedule:          v.WorkSchedule,
		Income:                v.Income,
		SalaryMax:             v.SalaryMax,
		SalaryMin:             v.SalaryMin,
		Pronounces:            v.Pronounces,
		Responsibilities:      v.Responsibilities,
		Requirements:          v.Requirements,
		Education:             v.Education,
		Experience:            v.Experience,
		SpecialPrograms:       v.SpecialPrograms,
		ComputerSkills:        v.ComputerSkills,
		ForeignLanguages:      v.ForeignLanguages,
		LanguageLevel:         v.LanguageLevel,
		HasBusinessTrips:      v.HasBusinessTrips,
		AdditionalInformation: v.AdditionalInformation,
	}
}

type VacancyFilter struct {
	Status *string `json:"status"`
	Name   *string `json:"name"`
}

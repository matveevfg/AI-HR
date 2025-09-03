package models

import "github.com/google/uuid"

type Resume struct {
	ID                       uuid.UUID `json:"id" bun:"id,pk"`
	Location                 string    `json:"location"`
	Citizenship              string    `json:"citizenship"`
	RelocationWillingness    string    `json:"relocation_willingness"`
	BusinessTripsWillingness string    `json:"business_trips_willingness"`
	DesiredPosition          string    `json:"desired_position"`
	Specializations          []string  `json:"specializations"`
	EmploymentTypes          []string  `json:"employment_types"`
	WorkSchedules            []string  `json:"work_schedules"`
	MaxCommuteTime           string    `json:"max_commute_time"`
	TotalExperience          string    `json:"total_experience"`
	Education                []string  `json:"education"`
	Courses                  []string  `json:"courses"`
	Languages                []string  `json:"languages"`
	Skills                   []string  `json:"skills"`
	HasCar                   bool      `json:"has_car"`
	LicenseCategory          string    `json:"license_category"`
	About                    string    `json:"about"`

	WorkPlaces []*WorkPlace `json:"work_places" bun:"rel:has-many,join:id=resume_id"`
}

type WorkPlace struct {
	ID               int       `json:"id" bun:"id,pk"`
	ResumeID         uuid.UUID `json:"resume_id"`
	Resume           *Resume   `json:"resume,omitempty" bun:"rel:has-one,join:resume_id=id"`
	Company          string    `json:"company"`
	Position         string    `json:"position"`
	Period           string    `json:"period"`
	Duration         string    `json:"duration"`
	Responsibilities []string  `json:"responsibilities"`
}

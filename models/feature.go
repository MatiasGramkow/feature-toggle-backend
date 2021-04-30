package models

// Feature Model
type Feature struct {
	ID                 uint   `json:"id"`
	Name               string `json:"name" gorm:"unique"`
	IsActive           bool   `json:"is_active" gorm:"type:boolean; column:is_active"`
	Description        string `json:"description"`
	DeactivationReason string `json:"deactivation_reason"`
	StartDate          string `json:"start_date"`
	StopDate           string `json:"stop_date"`
}

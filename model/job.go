package model

import "time"

type Job struct {
	Description  string    `json:"description,omitempty"`
	ContractType string    `json:"contractType,omitempty"`
	Duration     int16     `json:"duration,omitempty"`
	Remote       bool      `json:"remote,omitempty"`
	StudyLevel   string    `json:"studyLevel,omitempty"`
	StartDate    time.Time `json:"startDate"`
}

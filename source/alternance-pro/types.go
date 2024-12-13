package alternancepro

import "time"

type JobResponse struct {
	FoundJobs    bool   `json:"found_jobs"`
	Showing      string `json:"showing"`
	MaxNumPages  int    `json:"max_num_pages"`
	ShowingLinks string `json:"showing_links"`
	HTML         string `json:"html"`
}

type JobDetail struct {
	CompanyName  string
	City         string
	StartDate    time.Time
	ContractType string
	StudyLevel   string
}

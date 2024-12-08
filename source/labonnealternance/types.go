package labonnealternance

type Contact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone,omitempty"`
	Info  string `json:"info,omitempty"`
	IV    string `json:"iv,omitempty"`
}

type Rome struct {
	Code  string `json:"code"`
	Label string `json:"label"`
}

type Naf struct {
	Code  string `json:"code"`
	Label string `json:"label"`
}

type Place struct {
	Distance          float64 `json:"distance"`
	FullAddress       string  `json:"fullAddress,omitempty"`
	Latitude          float64 `json:"latitude,omitempty"`
	Longitude         float64 `json:"longitude,omitempty"`
	City              string  `json:"city,omitempty"`
	Address           string  `json:"address,omitempty"`
	Cedex             string  `json:"cedex,omitempty"`
	ZipCode           string  `json:"zipCode,omitempty"`
	Insee             string  `json:"insee,omitempty"`
	DepartementNumber string  `json:"departementNumber,omitempty"`
	Region            string  `json:"region,omitempty"`
	RemoteOnly        bool    `json:"remoteOnly,omitempty"`
}

type Headquarter struct {
	ID            string `json:"id"`
	UAI           string `json:"uai"`
	Name          string `json:"name"`
	Siret         string `json:"siret"`
	Type          string `json:"type,omitempty"`
	HasConvention string `json:"hasConvention,omitempty"`
	Place         Place  `json:"place,omitempty"`
}

type Company struct {
	ID            string      `json:"id"`
	UAI           string      `json:"uai"`
	Name          string      `json:"name"`
	Siret         string      `json:"siret,omitempty"`
	Size          string      `json:"size,omitempty"`
	Logo          string      `json:"logo,omitempty"`
	Description   string      `json:"description,omitempty"`
	SocialNetwork string      `json:"socialNetwork,omitempty"`
	URL           string      `json:"url,omitempty"`
	Mandataire    bool        `json:"mandataire,omitempty"`
	CreationDate  string      `json:"creationDate,omitempty"`
	Place         Place       `json:"place,omitempty"`
	Headquarter   Headquarter `json:"headquarter,omitempty"`
}

type Job struct {
	ID                   string `json:"id"`
	Description          string `json:"description,omitempty"`
	EmployeurDescription string `json:"employeurDescription,omitempty"`
	CreationDate         string `json:"creationDate,omitempty"`
	ContractType         string `json:"contractType,omitempty"`
	ContractDescription  string `json:"contractDescription,omitempty"`
	Duration             string `json:"duration,omitempty"`
	JobStartDate         string `json:"jobStartDate,omitempty"`
	RythmeAlternance     string `json:"rythmeAlternance,omitempty"`
	ElligibleHandicap    bool   `json:"elligibleHandicap,omitempty"`
	DureeContrat         string `json:"dureeContrat,omitempty"`
	QuantiteContrat      int64  `json:"quantiteContrat,omitempty"`
	Status               string `json:"status,omitempty"`
}

type PeJob struct {
	ID            string  `json:"id"`
	Title         string  `json:"title"`
	IdeaType      string  `json:"ideaType,omitempty"`
	URL           string  `json:"url,omitempty"`
	DetailsLoaded bool    `json:"detailsLoaded,omitempty"`
	Contact       Contact `json:"contact,omitempty"`
	Place         Place   `json:"place,omitempty"`
	Company       Company `json:"company,omitempty"`
	Job           Job     `json:"job,omitempty"`
	Romes         []Rome  `json:"romes,omitempty"`
	Nafs          []Naf   `json:"nafs,omitempty"`
}

type JobFormationsResponse struct {
	Jobs struct {
		PeJobs struct {
			Results []PeJob `json:"results,omitempty"`
		} `json:"peJobs,omitempty"`
	} `json:"jobs,omitempty"`
}

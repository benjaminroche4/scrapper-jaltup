package cities

type City struct {
	Name       string  `json:"name"`
	ZipCode    string  `json:"zipCode"`
	Department string  `json:"department"`
	Region     string  `json:"region"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
}

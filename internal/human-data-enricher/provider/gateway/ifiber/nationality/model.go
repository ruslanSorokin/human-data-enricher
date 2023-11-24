package ifiber

type country struct {
	CountryID   string  `json:"country_id"`
	Probability float64 `json:"probability"`
}

type response struct {
	Name      string    `json:"name"`
	Countries []country `json:"country"`
	Count     int64     `json:"count"`
}

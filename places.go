package main

type (
	Place struct {
		Geometry     geometry
		Icon         string
		Id           string
		Name         string
		OpeningHours hours  `json:"opening_hours"`
		PlaceId      string `json:"place_id"`
		Reference    string
		Scope        string
		Types        []string
		Vicinity     string
	}

	Places []Place

	geometry struct {
		Location location
	}

	location struct {
		Lat float64
		Lng float64
	}

	hours struct {
		OpenNow     bool     `json:"open_now"`
		WeekdayText []string `json:"weekday_text"`
	}

	rawData struct {
		Places Places `json:"results"`
		Meta   map[string]interface{}
	}
)

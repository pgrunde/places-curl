package main

type (
	Place struct {
		Geometry      location
		Icon          string
		Id            string
		Name          string
		Opening_hours hours
		Place_id      string
		Reference     string
		Scope         string
		Types         []string
		Vicinity      string
	}

	Places []Place

	location struct {
		Lat float64
		Lng float64
	}

	hours struct {
		Open_now     bool
		Weekday_text []string
	}

	rawData struct {
		Places Places `json:"results"`
		Meta   map[string]interface{}
	}
)

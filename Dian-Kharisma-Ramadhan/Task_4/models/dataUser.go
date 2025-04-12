package models

type Friend struct {
	Name     string
	Division string
}

type Data struct {
	VehicleChoice   string
	Items          []string
	Recommendations map[string]string
	Friends        []Friend
}
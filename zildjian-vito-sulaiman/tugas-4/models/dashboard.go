package models

type Dashboard struct {
	VehicleChoice   string
	Items           []string
	Recommendations map[string]string
	Friends         []Friend
}

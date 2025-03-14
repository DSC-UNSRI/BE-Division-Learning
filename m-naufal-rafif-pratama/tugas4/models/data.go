package models

type Data struct {
	Vehicle        string
	Items          []string
	Recommendations map[string]string
	Friends        map[string]string
}

func NewData() Data {
	return Data{
		Items:          []string{},
		Recommendations: make(map[string]string),
		Friends:        make(map[string]string),
	}
}

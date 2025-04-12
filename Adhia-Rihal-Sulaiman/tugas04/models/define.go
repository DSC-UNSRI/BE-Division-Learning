package models

type Transportation struct {
	Type string
}

type Item struct {
	Name string
}

type Recommendation struct {
	Category string
	Content  string
}

type Friend struct {
	Name     string
	Division string
}

type Dashboard struct {
	Transportation []Transportation
	Items          []Item
	Recommendations []Recommendation
	Friends        []Friend
}

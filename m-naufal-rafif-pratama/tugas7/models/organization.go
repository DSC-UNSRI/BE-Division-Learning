package models

type Organization struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

var OrganizationQuery = `
	CREATE TABLE IF NOT EXISTS organizations (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(100) NOT NULL UNIQUE,
		type VARCHAR(100) NOT NULL
	);
`

func (o Organization) IsValid() bool {
	return o.Name != "" && o.Type != ""
}
package models

type Student struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Pass  string `json:"password"`
	Major string `json:"major"`
	Year  int    `json:"year"`
	OrgID int    `json:"org_id"`
}

var StudentQuery = `
	CREATE TABLE IF NOT EXISTS students (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(100) NOT NULL,
		password VARCHAR(100) NOT NULL,
		major VARCHAR(100) NOT NULL,
		year INT NOT NULL,
		org_id INT
	);
`

func (s Student) IsValid() bool {
	return s.Name != "" && s.Email != "" && s.Pass != "" && s.Major != "" && s.Year > 0
}

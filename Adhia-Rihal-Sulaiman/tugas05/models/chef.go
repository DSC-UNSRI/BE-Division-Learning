package models  

type Chef struct {  
	ID           int    `json:"id"`  
	Name         string `json:"name"`  
	Speciality   string `json:"speciality"`  
	Experience   int    `json:"experience"`  
	Username     string `json:"username"`  
	PasswordHash string `json:"-"`  
}
package models  

type Menu struct {  
	ID          int     `json:"id"`  
	Name        string  `json:"name"`  
	Description string  `json:"description"`  
	Price       float64 `json:"price"`  
	ChefID      int     `json:"chef_id"`  
	Category    string  `json:"category"`  
}
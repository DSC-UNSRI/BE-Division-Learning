package models

type Country struct {
	ID        	int    	`json:"id"`
	NamaNegara 	string 	`json:"negara_asal"`
	KodeNegara	int 	`json:"kode_negara"`
}

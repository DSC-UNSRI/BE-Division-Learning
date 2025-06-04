package models

type negara struct {
	ID        	int    	`json:"id"`
	NamaNegara 	string 	`json:"negara_asal"`
	KodeNegara	int 	`json:"kode_negara"`
}
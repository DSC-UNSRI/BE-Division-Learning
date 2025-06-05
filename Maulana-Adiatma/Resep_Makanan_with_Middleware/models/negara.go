package models

type Negara struct {
	ID        	int    	`json:"id"`
	NamaNegara 	string 	`json:"negara_asal"`
	KodeNegara	int 	`json:"kode_negara"`
	AuthKey    	string 	`json:"auth_key"`
}
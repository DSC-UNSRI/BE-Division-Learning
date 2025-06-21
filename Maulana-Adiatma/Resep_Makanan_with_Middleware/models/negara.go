package models

type Negara struct {
    ID          int    `json:"id"`
    NamaNegara  string `json:"negara_asal"`
    KodeNegara  string `json:"kode_negara"`
    Email       string `json:"email_users"`
    Password    string `json:"-"`
    Token       string `json:"token_users"`
    Role        string `json:"role_users"`
}

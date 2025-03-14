package controllers

import (
	"backend-iftar-gdgoc/config"
)

func CekAutentikasi() bool {
	nama := config.AmbilVariabel("NAMA")
	email := config.AmbilVariabel("EMAIL")
	password := config.AmbilVariabel("PASSWORD")

	return nama != "" && email != "" && password != ""
}

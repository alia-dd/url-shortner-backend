package shortner

import (
	"url-shortener/api"
	"url-shortener/model"
)

func UrlShorner() {
	model.InitDB()
	api.ApiHandler()
}

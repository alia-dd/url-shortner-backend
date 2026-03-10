package shortner

import (
	"url-shortner-backend/api"
	"url-shortner-backend/model"
)

func UrlShorner() {
	model.InitDB()
	api.ApiHandler()
}

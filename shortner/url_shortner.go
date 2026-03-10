package shortner

import (
	"url-shortner-backend/api"
	"url-shortner-backend/initializer"
)

func UrlShorner() {
	initializer.LoadEnv()
	initializer.InitDB()
	api.ApiHandler()
}

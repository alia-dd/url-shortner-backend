package api

import (
	"log"
	"net/http"
	"os"
	"url-shortner-backend/controller"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type LongUrl struct {
	Url string `json:"url" binding:"required"`
}

func ApiHandler() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	gin.SetMode(gin.ReleaseMode)
	var route = gin.Default()

	route.POST("/shorten", postLong)
	route.GET("/:alias", getAlias)
	log.Println("Server running on port", port)
	route.Run(":" + port)
}

func postLong(c *gin.Context) {
	longUrl := LongUrl{}
	if err := c.ShouldBindBodyWith(&longUrl, binding.JSON); err != nil {
		c.String(400, `provide correct url.`)
		return
	}
	alias, err := controller.PostUrl(longUrl.Url)
	if err != nil {
		c.String(400, `provide correct url.`)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"alias": alias,
	})
}

func getAlias(c *gin.Context) {
	alias := c.Param("alias")
	// if exist,_,_ := controller.Checkifexist("Short_url",alias); exist{
	url, err := controller.GetUrlLongUrl(alias)
	if err != nil {
		c.JSON(404, `page not found.`)
		return
	}
	c.Redirect(http.StatusMovedPermanently, url)
	// }

}

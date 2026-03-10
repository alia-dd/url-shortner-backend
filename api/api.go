package api

import (
	"net/http"
	"url-shortner-backend/controller"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type LongUrl struct {
	Url string `json:"url" binding:"required"`
}

func ApiHandler() {
	gin.SetMode(gin.ReleaseMode)
	var route = gin.Default()
	route.POST("/shorten", postLong)
	route.GET("/:alias", getAlias)
	route.Run(":8000")
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
	c.String(http.StatusOK, alias)
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

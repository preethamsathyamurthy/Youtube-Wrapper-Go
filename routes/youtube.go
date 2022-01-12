package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var youtubeGroup *echo.Group

func ActivateYoutubeIndex(g *echo.Group) {
	youtubeGroup = g
	UseRoute(youtubeGroup, YoutubeRoutes)
}

func YoutubeRoutes(group *echo.Group) {
	group.GET("", sayHelloYoutube)
	group.GET("/", sayHelloYoutube)
	// group.GET("/healthCheck", healthCheck)
	// group.POST("/user", addUser)
}

func sayHelloYoutube(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, API/Youtube")
}

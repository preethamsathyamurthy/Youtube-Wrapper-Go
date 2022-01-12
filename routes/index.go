package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Inspired from a Stackoverflow answer
// https://stackoverflow.com/questions/58186129/how-to-import-sub-routes-from-other-packages

// This function passes the Group object to the Route parsing functions
func UseRoute(group *echo.Group, routes func(group *echo.Group)) {
	routes(group)
}

var mainGroup *echo.Group

// Initializing the API index
func ActivateIndex(g *echo.Group) {
	mainGroup = g
	UseRoute(mainGroup, IndexRoutes) // If this is not needed, we need to do maingroup.GET()
	// Developer Preference

	// sub routes - Keep Adding Sub Routes here
	youtubeGroup := mainGroup.Group("/youtube")
	ActivateYoutubeIndex(youtubeGroup)
}

// Main Group Routes
func IndexRoutes(group *echo.Group) {
	group.GET("", apiReached)
	group.GET("/", apiReached)
	group.GET("/healthCheck", healthCheck)
	// group.POST("/user", addUser)
}

// Main Group Functions
func apiReached(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, You have reached the /api. Call the sub functions!")
}

func healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "The Youtube Wrapper service is working fine")
}

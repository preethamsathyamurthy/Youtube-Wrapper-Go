package main

import (
	"fmt"

	logrus_stack "github.com/Gurpartap/logrus-stack"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/preethamsathyamurthy/Youtube-Wrapper-Go/framework"
	"github.com/preethamsathyamurthy/Youtube-Wrapper-Go/routes"
	"github.com/sirupsen/logrus"
	echologrus "github.com/spirosoik/echo-logrus"
)

func main() {
	e := echo.New()

	// Logrus Logger
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	//JSON Format // can set fluentD formats also
	logger.SetFormatter(&logrus.JSONFormatter{})

	// Add the stack hook.
	// Custom setting caller level to panic, fatal and error
	logger.AddHook(logrus_stack.LogrusStackHook{
		CallerLevels: []logrus.Level{logrus.PanicLevel, logrus.FatalLevel, logrus.ErrorLevel},
		StackLevels:  []logrus.Level{logrus.PanicLevel, logrus.FatalLevel, logrus.ErrorLevel},
	})

	// Writelogs function is a wrapper for calling Echo framework's logger
	// For centralized logging
	var writeLogs framework.EchoLogger
	writeLogs.Initialize(e)

	// Usage of logger middleware
	mw := echologrus.NewLoggerMiddleware(logger)
	e.Logger = mw

	e.Use(
		middleware.Recover(),   // Recover from all panics to always have your server up
		middleware.Logger(),    // Log everything to stdout
		middleware.RequestID(), // Generate a request id on the HTTP response headers for identification
	)

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		// Take required information from error and context and send it to a service like New Relic
		// or any logging device
		fmt.Println(c.Path(), c.QueryParams(), err.Error())

		// Call the default handler to return the HTTP response
		e.DefaultHTTPErrorHandler(err, c)
	}

	apiGroup := e.Group("/api")
	routes.ActivateIndex(apiGroup)

	e.Logger.Fatal(e.Start(":1323"))
}

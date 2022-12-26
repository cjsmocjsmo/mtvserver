package main

import (
	// "log"
	"net/http"
	// "os"

	"github.com/labstack/echo/v4"
)

// func mtv_logging() {
// 	logfile := os.Getenv("MTV_LOG_BASE_PATH") + "/MTV.log"
// 	// If the file doesn't exist, create it or append to the file
// 	file, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.SetOutput(file)
// 	log.Println("MTV logging started")
// }

func main() {
	// mtv_logging()
	MOVSetup()
	// TVSetUp()
	

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/IntActionHandler", IntActionHandler)
	e.GET("/IntCartoons", IntCartoonsHandler)
	e.Logger.Fatal(e.Start(":8888"))
}

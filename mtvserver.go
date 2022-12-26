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
	TVSetUp()
	

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/Action", IntActionHandler)
	e.GET("/Cartoons", IntCartoonsHandler)
	e.GET("/Comedy", IntComedyHandler)
	e.GET("/Drama", IntDramaHandler)
	e.GET("/Godzilla", IntGodzillaHandler)
	e.GET("/IndianaJones", IntIndianaJonesHandler)
	e.GET("/JohnWayne", IntJohnWayneHandler)
	e.GET("/JurassicPark", IntJurassicParkHandler)
	e.GET("/KingsMan", IntKingsManHandler)
	e.GET("/HarryPotter", IntHarryPotterHandler)
	e.GET("/MenInBlack", IntMenInBlackHandler)
	e.GET("/Misc", IntMiscHandler)
	e.GET("/SciFi", IntSciFiHandler)
	e.GET("/StarTrek", IntStarTrekHandler)
	e.GET("/StarWars", IntStarWarsHandler)
	e.GET("/SuperHeros", IntSuperHerosHandler)
	e.GET("/Tremors", IntTremorsHandler)
	e.GET("/JohnWick", IntJohnWickHandler)
	e.GET("/Pirates", IntPiratesHandler)
	e.GET("/BruceWillis", IntBruceWillisHandler)
	e.GET("/Fantasy", IntFantasyHandler)
	e.GET("/Riddick", IntRiddickHandler)
	e.GET("/TomCruize", IntTomCruizeHandler)
	e.GET("/XMen", IntXMenHandler)
	e.GET("/Documentary", IntDocumentaryHandler)
	e.GET("/TheRock", IntTheRockHandler)
	e.GET("/NicolasCage", IntNicolasCageHandler)
	e.GET("/JamesBond", IntJamesBondHandler)
	e.GET("/Transformers", IntTransformersHandler)

	e.GET("/STTV", IntSTTVHandler)
	
	e.Logger.Fatal(e.Start(":8888"))
}

package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"os"
)

func mtv_logging() {
	logfile := os.Getenv("MTV_LOG_BASE_PATH") + "/MTV.log"
	// If the file doesn't exist, create it or append to the file
	file, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	log.Println("MTV logging started")
}

func main() {

	mtv_logging()
	MOVSetup()
	TVSetUp()

	e := echo.New()
	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, mtv is up and running")
	})

	e.GET("/Action", IntActionHandler)
	e.GET("/Arnold", IntArnoldHandler)

	e.GET("/Cartoons", IntCartoonsHandler)
	e.GET("/TinkerBell", IntTinkerBell)
	e.GET("/CharlieBrown", IntCharlieBrown)
	e.GET("/Minions", IntMinions)
	e.GET("/Buzz", IntBuzz)
	e.GET("/Oldies", IntOldies)

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
	e.GET("/TNG", IntTNGHandler)
	e.GET("/Enterprise", IntEnterpriseHandler)
	e.GET("/Discovery", IntDiscoveryHandler)
	e.GET("/Voyager", IntVoyagerHandler)
	e.GET("/Picard", IntPicardHandler)
	e.GET("/Prodigy", IntProdigyHandler)
	e.GET("/StrangeNewWorlds", IntStrangeNewWorldsHandler)
	e.GET("/LowerDecks", IntLowerDecksHandler)

	e.GET("/Orville", IntOrvilleHandler)
	e.GET("/LostInSpace", IntLostInSpaceHandler)
	e.GET("/AlteredCarbon", IntAlteredCarbonHandler)
	e.GET("/RaisedByWolves", IntRaisedByWolvesHandler)
	e.GET("/ForAllManKind", IntForAllManKindHandler)
	e.GET("/AlienWorlds", IntAlienWorldsHandler)
	e.GET("/YTheLastMan", IntYTheLastManHandler)
	e.GET("/Foundation", IntFoundationHandler)
	e.GET("/MastersOfTheUniverse", IntMastersOfTheUniverseHandler)
	e.GET("/WheelOfTime", IntWheelOfTimeHandler)
	e.GET("/CowboyBebop", IntCowboyBebopHandler)
	e.GET("/Reacher", IntReacherHandler)
	e.GET("/Halo", IntHaloHandler)
	e.GET("/HouseOfTheDragon", IntHouseOfTheDragonHandler)
	e.GET("/TheLordOfTheRingsTheRingsOfPower", IntTheLordOfTheRingsTheRingsOfPowerHandler)
	e.GET("/NightSky", IntNightSkyHandler)
	e.GET("/PrehistoricPlanet", IntPrehistoricPlanetHandler)

	e.GET("/TheLastOfUs", IntTheLastOfUsHandler)
	e.GET("/1923", IntHFord1923Handler)

	e.GET("/BlackKight", IntBlackKnightHandler)

	e.GET("/Mandalorian", IntMandalorianHandler)
	e.GET("/BookOfBobaFett", IntBookOfBobaFettHandler)
	e.GET("/TheBadBatch", IntTheBadBatchHandler)
	e.GET("/ObiWanKenobi", IntObiWanKenobiHandler)
	e.GET("/Andor", IntAndorHandler)
	e.GET("/TalesOfTheJedi", IntTalesOfTheJediHandler)
	e.GET("/Visions", IntVisionsHandler)

	e.GET("/WandaVision", IntWandaVisionHandler)
	e.GET("/FalconWinterSoldier", IntFalconWinterSoldierHandler)
	e.GET("/Loki", IntLokiHandler)
	e.GET("/WhatIf", IntWhatIfHandler)
	e.GET("/Hawkeye", IntHawkeyeHandler)
	e.GET("/MoonKnight", IntMoonKnightHandler)
	e.GET("/MSMarvel", IntMSMarvelHandler)
	e.GET("/IAmGroot", IntIAmGrootHandler)
	e.GET("/SheHulk", IntSheHulkHandler)

	e.GET("/FooBar", IntFooBarHandler)

	e.Static("/static", "/root/static")
	e.Logger.Fatal(e.Start(":8888"))
}

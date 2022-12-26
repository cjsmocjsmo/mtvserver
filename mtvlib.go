///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
// LICENSE: GNU General Public License, version 2 (GPLv2)
// Copyright 2016, Charlie J. Smotherman
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of the GNU General Public License v2
// as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program; if not, write to the Free Software
// Foundation, Inc., 59 Temple Place - Suite 330, Boston, MA  02111-1307, USA.
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
package main

import (
	// "encoding/json"
	"fmt"
	// "io/ioutil"
	"log"
	"net/http"

	// "net/url"
	"os"

	"github.com/labstack/echo/v4"

	// "github.com/cjsmocjsmo/movgo"
	// "github.com/cjsmocjsmo/tvgo"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	// "github.com/gorilla/handlers"
	// "github.com/gorilla/mux"
)

// DBcon is exported because I want it
func DBcon() *mgo.Session {
	s, err := mgo.Dial(os.Getenv("MTV_MONGODB_ADDRESS"))
	if err != nil {
		log.Println("Session creation dial error")
		log.Println(err)
	}
	log.Println("Session Connection to db established")
	return s
}

func IntActionHandler(c echo.Context) error {
	log.Println("IntActionHandler start")
	ses := DBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var ActionMedia []map[string]string
	b1 := bson.M{"catagory": "Action"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&ActionMedia)
	if err != nil {
		log.Println("ActionHandler db call error")
		log.Println(err)
	}
	log.Println(ActionMedia)
	return c.JSON(http.StatusOK, ActionMedia)
}

func IntCartoonsHandler(c echo.Context) error {
	log.Println("IntCartoonsHandler started")
	ses := DBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	b1 := bson.M{"catagory": "Cartoons"}
	b2 := bson.M{"_id": 0}
	var CartoonMedia []map[string]string
	err := MTc.Find(b1).Select(b2).All(&CartoonMedia)
	if err != nil {
		log.Println("cartoonHandler db call error")
		log.Println(err)
	}
	return c.JSON(http.StatusOK, CartoonMedia)
}

func IntComedyHandler(c echo.Context) error {
	log.Println("IntComedyHandler started")
	ses := DBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var ComedyMedia []map[string]string
	b1 := bson.M{"catagory": "Comedy"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&ComedyMedia)
	if err != nil {
		log.Println("comedyHandler db call error")
		log.Println(err)
	}
	return c.JSON(http.StatusOK, ComedyMedia)
}

func IntDramaHandler(c echo.Context) error {
	log.Println("IntDramaHandler started")
	ses := DBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var DramaMedia []map[string]string
	b1 := bson.M{"catagory": "Drama"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&DramaMedia)
	if err != nil {
		log.Println("dramaHandler db call error")
		log.Println(err)
	}
	return c.JSON(http.StatusOK, DramaMedia)
}

func IntGodzillaHandler(c echo.Context) error {
	log.Println("initGodzillaHandler started")
	ses := DBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var GodzillaMedia []map[string]string
	b1 := bson.M{"catagory": "Godzilla"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&GodzillaMedia)
	if err != nil {
		log.Println("godzillaHandler db call error")
		log.Println(err)
	}
	return c.JSON(http.StatusOK, GodzillaMedia)
}

func IntIndianaJonesHandler(c echo.Context) error {
	log.Println("IntIndianaJonesHandler started")
	ses := DBcon()
	defer ses.Close()
	MTrc := ses.DB("moviegobs").C("moviegobs")
	var IndianaJonesMedia []map[string]string
	b1 := bson.M{"catagory": "IndianaJones"}
	b2 := bson.M{"_id": 0}
	err := MTrc.Find(b1).Select(b2).All(&IndianaJonesMedia)
	if err != nil {
		log.Println("indianaJonesHandler db call error")
		log.Println(err)
	}
	return c.JSON(http.StatusOK, IndianaJonesMedia)
}

func IntJohnWayneHandler(c echo.Context) error {
	log.Println("johnWayneHandler started")
	ses := DBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var JohnWayneMedia []map[string]string
	b1 := bson.M{"catagory": "JohnWayne"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&JohnWayneMedia)
	if err != nil {
		log.Println("johnwayneHandler db call error")
		log.Println(err)
	}
	return c.JSON(http.StatusOK, JohnWayneMedia)
}

func IntJurassicParkHandler(c echo.Context) error {
	log.Println("JurassicParkHandler started")
	ses := DBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var JurassicParkMedia []map[string]string
	b1 := bson.M{"catagory": "JurassicPark"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&JurassicParkMedia)
	if err != nil {
		log.Println("JurassicParkHandler db call error")
		log.Println(err)
	}
	return c.JSON(http.StatusOK, JurassicParkMedia)
}

func IntKingsManHandler(c echo.Context) error {
	log.Println("KingsManHandler started")
	ses := DBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var KingsmanMedia []map[string]string
	b1 := bson.M{"catagory": "Kingsman"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&KingsmanMedia)
	if err != nil {
		log.Println("Kingsman db call error")
		log.Println(err)
	}
	return c.JSON(http.StatusOK, KingsmanMedia)
}

func IntHarryPotterHandler(c echo.Context) error {
	log.Println("HarryPotterHandler started")
	ses := DBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var HarryPotterMedia []map[string]string
	b1 := bson.M{"catagory": "HarryPotter"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&HarryPotterMedia)
	if err != nil {
		log.Println(err)
	}
	return c.JSON(http.StatusOK, HarryPotterMedia)
}

func IntMenInBlackHandler(c echo.Context) error {
	log.Println("MenInBlackHandler started")
	ses := DBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var MenInBlackMedia []map[string]string
	b1 := bson.M{"catagory": "MenInBlack"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&MenInBlackMedia)
	if err != nil {
		log.Println(err)
	}
	return c.JSON(http.StatusOK, MenInBlackMedia)
}

func IntMiscHandler(c echo.Context) error {
	log.Println("MiscHandler started")
	ses := DBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var MiscMedia []map[string]string
	b1 := bson.M{"catagory": "Misc"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&MiscMedia)
	if err != nil {
		log.Println(err)
	}
	return c.JSON(http.StatusOK, MiscMedia)
}

func IntSciFiHandler(c echo.Context) error {
	log.Println("SciFiHandler started")
	ses := DBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var SciFiMedia []map[string]string
	b1 := bson.M{"catagory": "SciFi"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&SciFiMedia)
	if err != nil {
		log.Println("InitSciFiHandler db call error")
		log.Println(err)
	}
	return c.JSON(http.StatusOK, SciFiMedia)
}

func IntStarTrekHandler(c echo.Context) error {
	log.Println("StarTrekHandler started")
	ses := DBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var StarTrekMedia []map[string]string
	b1 := bson.M{"catagory": "StarTrek"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&StarTrekMedia)
	if err != nil {
		log.Println("StarTrekHandler db call error")
		log.Println(err)
	}
	return c.JSON(http.StatusOK, StarTrekMedia)
}

func IntStarWarsHandler(c echo.Context) error {
	log.Println("StarWarsHandler started")
	ses := DBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var StarWarsMedia []map[string]string
	b1 := bson.M{"catagory": "StarWars"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&StarWarsMedia)
	if err != nil {
		log.Printf("StarWarsHandler db call error")
		log.Println(err)
	}
	return c.JSON(http.StatusOK, StarWarsMedia)
}

func IntSuperHerosHandler(c echo.Context) error {
	log.Println("SuperHerosHandler started")
	ses := DBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var SuperHerosMedia []map[string]string
	b1 := bson.M{"catagory": "SuperHeros"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&SuperHerosMedia)
	if err != nil {
		log.Printf("SuperHerosHandler db call error")
		log.Println(err)
	}
	return c.JSON(http.StatusOK, SuperHerosMedia)
}

func IntTremorsHandler(c echo.Context) error {
	log.Println("TremorsHandler started")
	ses := DBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var TremorsMedia []map[string]string
	b1 := bson.M{"catagory": "Tremors"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&TremorsMedia)
	if err != nil {
		log.Printf("TremorsHandler db call error")
		log.Println(err)
	}
	return c.JSON(http.StatusOK, TremorsMedia)
}

func IntJohnWickHandler(c echo.Context) error {
	log.Println("JohnWickHandler started")
	ses := DBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var JohnWickMedia []map[string]string
	b1 := bson.M{"catagory": "JohnWick"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&JohnWickMedia)
	if err != nil {
		log.Printf("JohnWickHandler db call error")
		log.Println(err)
	}
	return c.JSON(http.StatusOK, JohnWickMedia)
}

func IntPiratesHandler(c echo.Context) error {
	log.Println("PiratesHandler started")
	ses := DBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var PiratesMedia []map[string]string
	b1 := bson.M{"catagory": "Pirates"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&PiratesMedia)
	if err != nil {
		log.Printf("PiratesHandler db call error")
		log.Println(err)
	}
	return c.JSON(http.StatusOK, PiratesMedia)
}

func IntBruceWillisHandler(c echo.Context) error {
	log.Println("BruceWillisHandler started")
	ses := DBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var DieHardMedia []map[string]string
	b1 := bson.M{"catagory": "BruceWillis"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&DieHardMedia)
	if err != nil {
		log.Printf("BruceWillisHandler db call error")
		log.Println(err)
	}
	return c.JSON(http.StatusOK, DieHardMedia)
}

func IntFantasyHandler(c echo.Context) error {
	log.Println("FantasyHandler started")
	ses := DBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var FantasyMedia []map[string]string
	b1 := bson.M{"catagory": "Fantasy"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&FantasyMedia)
	if err != nil {
		log.Printf("FantasyHandler db call error")
		log.Println(err)
	}
	return c.JSON(http.StatusOK, FantasyMedia)
}

func IntRiddickHandler(c echo.Context) error {
	log.Println("RiddickHandler started")
	ses := DBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var RiddickMedia []map[string]string
	b1 := bson.M{"catagory": "Riddick"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&RiddickMedia)
	if err != nil {
		log.Printf("RiddickHandler db call error")
		log.Println(err)
	}
	return c.JSON(http.StatusOK, RiddickMedia)
}

func IntTomCruizeHandler(c echo.Context) error {
	log.Println("TomCruizeHandler started")
	ses := DBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var TCMedia []map[string]string
	b1 := bson.M{"catagory": "TomCruize"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&TCMedia)
	if err != nil {
		log.Printf("TomCruizeHandler db call error")
		log.Println(err)
	}
	return c.JSON(http.StatusOK, TCMedia)
}

func IntXMenHandler(c echo.Context) error {
	log.Println("XMenHandler started")
	ses := DBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var XMenMedia []map[string]string
	b1 := bson.M{"catagory": "XMen"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&XMenMedia)
	if err != nil {
		log.Printf("XMenHandler db call error")
		log.Println(err)
	}
	return c.JSON(http.StatusOK, XMenMedia)
}

func IntDocumentaryHandler(c echo.Context) error {
	log.Println("DocumentaryHandler started")
	ses := DBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var DocumentaryMedia []map[string]string
	b1 := bson.M{"catagory": "Documentary"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&DocumentaryMedia)
	if err != nil {
		log.Println("DocumentaryHandler db call error")
		log.Println(err)
	}
	return c.JSON(http.StatusOK, DocumentaryMedia)
}

func IntTheRockHandler(c echo.Context) error {
	log.Println("TheRockHandler started")
	ses := DBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var TheRockMedia []map[string]string
	b1 := bson.M{"catagory": "TheRock"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&TheRockMedia)
	if err != nil {
		log.Println("TheRockHandler db call error")
		log.Println(err)
	}
	return c.JSON(http.StatusOK, TheRockMedia)
}

func IntNicolasCageHandler(c echo.Context) error {
	log.Println("NicolasCage started")
	ses := DBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var NicolasCageMedia []map[string]string
	b1 := bson.M{"catagory": "NicolasCage"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&NicolasCageMedia)
	if err != nil {
		log.Println("NicolasCage db call error")
		log.Println(err)
	}
	return c.JSON(http.StatusOK, NicolasCageMedia)
}

func IntJamesBondHandler(c echo.Context) error {
	log.Println("IntJamesBondHandler start")
	ses := DBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var JamesBondMedia []map[string]string
	b1 := bson.M{"catagory": "JamesBond"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&JamesBondMedia)
	if err != nil {
		log.Println("JamesBondHandler db call error")
		log.Println(err)
	}
	log.Println(JamesBondMedia)
	return c.JSON(http.StatusOK, JamesBondMedia)
}

func IntTransformersHandler(c echo.Context) error {
	log.Println("IntTransformersHandler start")
	ses := DBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var TransformersMedia []map[string]string
	b1 := bson.M{"catagory": "Transformers"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&TransformersMedia)
	if err != nil {
		log.Println("TransformersHandler db call error")
		log.Println(err)
	}
	log.Println(TransformersMedia)
	return c.JSON(http.StatusOK, TransformersMedia)
}

// func playMediaReactHandler(c echo.Context) error {
// 	log.Println(" started")
// 	u, err := url.Parse(r.URL.String())
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	m, _ := url.ParseQuery(u.RawQuery)
// 	mf := m["movie"]
// 	omxAddr := os.Getenv("MTV_OMXPLAYER_ADDRESS_REACT")
// 	u, _ = url.Parse(omxAddr)
// 	q, _ := url.ParseQuery(u.RawQuery)
// 	q.Add("medPath", mf[0])
// 	u.RawQuery = q.Encode()

// 	log.Println("this is u.String()")
// 	log.Println(u.String())

// 	resp, err := http.Get(u.String())
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	defer resp.Body.Close()
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	Abody := string(body)
// 	log.Println(Abody)
// 	return c.JSON(http.StatusOK, Abody)
// }

func IntSTTVHandler(c echo.Context) error {
	log.Println("STTVHandler started")
	// u, err := url.Parse(r.URL.String())
	// if err != nil {
	// 	log.Println("STTVHandler url parse error")
	// 	log.Println(err)
	// }
	// m, eff := url.ParseQuery(u.RawQuery)
	// if eff != nil {
	// 	log.Println("usrl parsequery error")
	// 	log.Println(eff)
	// }
	// s1 := m["season"][0]
	s1 := c.QueryParam("season")
	ses := DBcon()
	defer ses.Close()
	MTyc := ses.DB("tvgobs").C("tvgobs")
	var STTVMedia []map[string]string
	b1 := bson.M{"catagory": "STTV", "season": s1}
	b2 := bson.M{"_id": 0}
	fmt.Printf("this is b1 %s", b1)
	errG := MTyc.Find(b1).Select(b2).Sort("episode").All(&STTVMedia)
	if errG != nil {
		log.Println("STTVHandler db call error")
		log.Println(errG)
	}
	return c.JSON(http.StatusOK, STTVMedia)
}

// func IntTNGHandler(c echo.Context) error {
// 	log.Println("TNGHandler started")
// 	u, err := url.Parse(r.URL.String())
// 	if err != nil {
// 		log.Println("TNGHandler url parse error")
// 		log.Println(err)
// 	}
// 	m, eff := url.ParseQuery(u.RawQuery)
// 	if eff != nil {
// 		log.Println("usrl parsequery error")
// 		log.Println(eff)
// 	}
// 	s1 := m["season"][0]
// 	ses := DBcon()
// 	defer ses.Close()
// 	MTyc := ses.DB("tvgobs").C("tvgobs")
// 	var nextGenerationMedia []map[string]string
// 	b1 := bson.M{"catagory": "TNG", "season": s1}
// 	b2 := bson.M{"_id": 0}
// 	errG := MTyc.Find(b1).Select(b2).Sort("episode").All(&nextGenerationMedia)
// 	if errG != nil {
// 		log.Println("TNGHandler db call error")
// 		log.Println(errG)
// 	}
// 	return c.JSON(http.StatusOK, nextGenerationMedia)
// }

// func IntEnterpriseHandler(c echo.Context) error {
// 	log.Println(" started")
// 	u, err := url.Parse(r.URL.String())
// 	if err != nil {
// 		log.Println(" url parse error")
// 		log.Println(err)
// 	}
// 	m, eff := url.ParseQuery(u.RawQuery)
// 	if eff != nil {
// 		log.Println("usrl parsequery error")
// 		log.Println(eff)
// 	}
// 	s1 := m["season"][0]
// 	ses := DBcon()
// 	defer ses.Close()
// 	MTyc := ses.DB("tvgobs").C("tvgobs")
// 	var enterpriseMedia []map[string]string
// 	b1 := bson.M{"catagory": "Enterprise", "season": s1}
// 	b2 := bson.M{"_id": 0}
// 	errG := MTyc.Find(b1).Select(b2).Sort("episode").All(&enterpriseMedia)
// 	if errG != nil {
// 		log.Println(errG)
// 	}
// 	return c.JSON(http.StatusOK, enterpriseMedia)
// }

// func IntDiscoveryHandler(c echo.Context) error {
// 	log.Println(" started")
// 	u, err := url.Parse(r.URL.String())
// 	if err != nil {
// 		log.Println(" url parse error")
// 		log.Println(err)
// 	}
// 	m, eff := url.ParseQuery(u.RawQuery)
// 	if eff != nil {
// 		log.Println("usrl parsequery error")
// 		log.Println(eff)
// 	}
// 	s1 := m["season"][0]
// 	ses := DBcon()
// 	defer ses.Close()
// 	MTyc := ses.DB("tvgobs").C("tvgobs")
// 	var discoveryMedia []map[string]string
// 	b1 := bson.M{"catagory": "Discovery", "season": s1}
// 	b2 := bson.M{"_id": 0}
// 	errG := MTyc.Find(b1).Select(b2).Sort("episode").All(&discoveryMedia)
// 	if errG != nil {
// 		log.Println(errG)
// 	}
// 	return c.JSON(http.StatusOK, discoveryMedia)
// }

// func IntVoyagerHandler(c echo.Context) error {
// 	log.Println(" started")
// 	u, err := url.Parse(r.URL.String())
// 	if err != nil {
// 		log.Println(" url parse error")
// 		log.Println(err)
// 	}
// 	m, eff := url.ParseQuery(u.RawQuery)
// 	if eff != nil {
// 		log.Println("usrl parsequery error")
// 		log.Println(eff)
// 	}
// 	s1 := m["season"][0]
// 	ses := DBcon()
// 	defer ses.Close()
// 	MTyc := ses.DB("tvgobs").C("tvgobs")
// 	var voyagerMedia []map[string]string
// 	b1 := bson.M{"catagory": "Voyager", "season": s1}
// 	b2 := bson.M{"_id": 0}
// 	errG := MTyc.Find(b1).Select(b2).Sort("episode").All(&voyagerMedia)
// 	if errG != nil {
// 		log.Println(errG)
// 	}
// 	return c.JSON(http.StatusOK, voyagerMedia)

// }

// func IntOrvilleHandler(c echo.Context) error {
// 	log.Println(" started")
// 	u, err := url.Parse(r.URL.String())
// 	if err != nil {
// 		log.Println(" url parse error")
// 		log.Println(err)
// 	}
// 	m, eff := url.ParseQuery(u.RawQuery)
// 	if eff != nil {
// 		log.Println("usrl parsequery error")
// 		log.Println(eff)
// 	}
// 	s1 := m["season"][0]
// 	ses := DBcon()
// 	defer ses.Close()
// 	MTtc := ses.DB("tvgobs").C("tvgobs")
// 	var OrvilleMedia []map[string]string
// 	b1 := bson.M{"catagory": "Orville", "season": s1}
// 	b2 := bson.M{"_id": 0}
// 	errG := MTtc.Find(b1).Select(b2).Sort("episode").All(&OrvilleMedia)
// 	if errG != nil {
// 		log.Println(errG)
// 	}
// 	return c.JSON(http.StatusOK, OrvilleMedia)
// }

// func IntLostInSpaceHandler(c echo.Context) error {
// 	log.Println(" started")
// 	u, err := url.Parse(r.URL.String())
// 	if err != nil {
// 		log.Println(" url parse error")
// 		log.Println(err)
// 	}
// 	m, eff := url.ParseQuery(u.RawQuery)
// 	if eff != nil {
// 		log.Println("usrl parsequery error")
// 		log.Println(eff)
// 	}
// 	s1 := m["season"][0]
// 	ses := DBcon()
// 	defer ses.Close()
// 	MTtc := ses.DB("tvgobs").C("tvgobs")
// 	var LostInSpaceMedia []map[string]string
// 	b1 := bson.M{"catagory": "Lost In Space", "season": s1}
// 	b2 := bson.M{"_id": 0}
// 	errG := MTtc.Find(b1).Select(b2).Sort("episode").All(&LostInSpaceMedia)
// 	if errG != nil {
// 		log.Println(errG)
// 	}
// 	log.Println(&LostInSpaceMedia)
// 	return c.JSON(http.StatusOK, LostInSpaceMedia)
// }

// func IntPicardHandler(c echo.Context) error {
// 	log.Println(" started")
// 	u, err := url.Parse(r.URL.String())
// 	if err != nil {
// 		log.Println(" url parse error")
// 		log.Println(err)
// 	}
// 	m, eff := url.ParseQuery(u.RawQuery)
// 	if eff != nil {
// 		log.Println("usrl parsequery error")
// 		log.Println(eff)
// 	}
// 	s1 := m["season"][0]
// 	ses := DBcon()
// 	defer ses.Close()
// 	MTtc := ses.DB("tvgobs").C("tvgobs")
// 	var PicardMedia []map[string]string
// 	b1 := bson.M{"catagory": "Picard", "season": s1}
// 	b2 := bson.M{"_id": 0}
// 	errG := MTtc.Find(b1).Select(b2).Sort("episode").All(&PicardMedia)
// 	if errG != nil {
// 		log.Println(errG)
// 	}
// 	log.Println(&PicardMedia)
// 	return c.JSON(http.StatusOK, PicardMedia)
// }

// func IntMandalorianHandler(c echo.Context) error {
// 	log.Println(" started")
// 	u, err := url.Parse(r.URL.String())
// 	if err != nil {
// 		log.Println(" url parse error")
// 		log.Println(err)
// 	}
// 	m, eff := url.ParseQuery(u.RawQuery)
// 	if eff != nil {
// 		log.Println("usrl parsequery error")
// 		log.Println(eff)
// 	}
// 	s1 := m["season"][0]
// 	ses := DBcon()
// 	defer ses.Close()
// 	MTtc := ses.DB("tvgobs").C("tvgobs")
// 	var MandalorianMedia []map[string]string
// 	b1 := bson.M{"catagory": "Mandalorian", "season": s1}
// 	b2 := bson.M{"_id": 0}
// 	errG := MTtc.Find(b1).Select(b2).Sort("episode").All(&MandalorianMedia)
// 	if errG != nil {
// 		log.Println(errG)
// 	}
// 	log.Println(&MandalorianMedia)
// 	return c.JSON(http.StatusOK, MandalorianMedia)
// }

// func IntAlteredCarbonHandler(c echo.Context) error {
// 	log.Println(" started")
// 	u, err := url.Parse(r.URL.String())
// 	if err != nil {
// 		log.Println(" url parse error")
// 		log.Println(err)
// 	}
// 	m, eff := url.ParseQuery(u.RawQuery)
// 	if eff != nil {
// 		log.Println("usrl parsequery error")
// 		log.Println(eff)
// 	}
// 	s1 := m["season"][0]
// 	ses := DBcon()
// 	defer ses.Close()
// 	MTyc := ses.DB("tvgobs").C("tvgobs")
// 	var alteredcarbonMedia []map[string]string
// 	b1 := bson.M{"catagory": "AlteredCarbon", "season": s1}
// 	b2 := bson.M{"_id": 0}
// 	errG := MTyc.Find(b1).Select(b2).Sort("episode").All(&alteredcarbonMedia)
// 	if errG != nil {
// 		log.Println(errG)
// 	}
// 	return c.JSON(http.StatusOK, alteredcarbonMedia)
// }

// func IntLowerDecksHandler(c echo.Context) error {
// 	log.Println(" started")
// 	u, err := url.Parse(r.URL.String())
// 	if err != nil {
// 		log.Println(" url parse error")
// 		log.Println(err)
// 	}
// 	m, eff := url.ParseQuery(u.RawQuery)
// 	if eff != nil {
// 		log.Println("usrl parsequery error")
// 		log.Println(eff)
// 	}
// 	s1 := m["season"][0]
// 	ses := DBcon()
// 	defer ses.Close()
// 	MTyc := ses.DB("tvgobs").C("tvgobs")
// 	var lowerdecksMedia []map[string]string
// 	b1 := bson.M{"catagory": "LowerDecks", "season": s1}
// 	b2 := bson.M{"_id": 0}
// 	errG := MTyc.Find(b1).Select(b2).Sort("episode").All(&lowerdecksMedia)
// 	if errG != nil {
// 		log.Println(errG)
// 	}
// 	return c.JSON(http.StatusOK, lowerdecksMedia)
// }

// func IntRaisedByWolvesHandler(c echo.Context) error {
// 	log.Println(" started")
// 	u, err := url.Parse(r.URL.String())
// 	if err != nil {
// 		log.Println(" url parse error")
// 		log.Println(err)
// 	}
// 	m, eff := url.ParseQuery(u.RawQuery)
// 	if eff != nil {
// 		log.Println("usrl parsequery error")
// 		log.Println(eff)
// 	}
// 	s1 := m["season"][0]
// 	ses := DBcon()
// 	defer ses.Close()
// 	MTyc := ses.DB("tvgobs").C("tvgobs")
// 	var raisedbywolvesMedia []map[string]string
// 	b1 := bson.M{"catagory": "RaisedByWolves", "season": s1}
// 	b2 := bson.M{"_id": 0}
// 	errG := MTyc.Find(b1).Select(b2).Sort("episode").All(&raisedbywolvesMedia)
// 	if errG != nil {
// 		log.Println(errG)
// 	}
// 	return c.JSON(http.StatusOK, raisedbywolvesMedia)
// }

// func IntForAllManKindHandler(c echo.Context) error {
// 	log.Println("Starting initForAllManKind")
// 	log.Println(" started")
// 	u, err := url.Parse(r.URL.String())
// 	if err != nil {
// 		log.Println(" url parse error")
// 		log.Println(err)
// 	}
// 	m, eff := url.ParseQuery(u.RawQuery)
// 	if eff != nil {
// 		log.Println("usrl parsequery error")
// 		log.Println(eff)
// 	}
// 	s1 := m["season"][0]
// 	log.Println(s1)
// 	ses := DBcon()
// 	defer ses.Close()
// 	MTyc := ses.DB("tvgobs").C("tvgobs")
// 	var forallmankindMedia []map[string]string
// 	b1 := bson.M{"catagory": "ForAllManKind", "season": s1}
// 	b2 := bson.M{"_id": 0}
// 	errG := MTyc.Find(b1).Select(b2).Sort("episode").All(&forallmankindMedia)
// 	if errG != nil {
// 		log.Println(errG)
// 	}
// 	return c.JSON(http.StatusOK, forallmankindMedia)
// }

// func IntAlienWorldsHandler(c echo.Context) error {
// 	log.Println("Starting initAlienWorlds")
// 	log.Println(" started")
// 	u, err := url.Parse(r.URL.String())
// 	if err != nil {
// 		log.Println(" url parse error")
// 		log.Println(err)
// 	}
// 	m, eff := url.ParseQuery(u.RawQuery)
// 	if eff != nil {
// 		log.Println("usrl parsequery error")
// 		log.Println(eff)
// 	}
// 	s1 := m["season"][0]
// 	log.Println(s1)
// 	ses := DBcon()
// 	defer ses.Close()
// 	MTyc := ses.DB("tvgobs").C("tvgobs")
// 	var alienworldsMedia []map[string]string
// 	b1 := bson.M{"catagory": "AlienWorlds", "season": `01`}
// 	b2 := bson.M{"_id": 0}
// 	errG := MTyc.Find(b1).Select(b2).Sort("episode").All(&alienworldsMedia)
// 	if errG != nil {
// 		log.Println(errG)
// 	}
// 	return c.JSON(http.StatusOK, alienworldsMedia)
// }

// func IntWandaVisionHandler(c echo.Context) error {
// 	log.Println("Starting initWandaVision")
// 	log.Println(" started")
// 	u, err := url.Parse(r.URL.String())
// 	if err != nil {
// 		log.Println(" url parse error")
// 		log.Println(err)
// 	}
// 	m, eff := url.ParseQuery(u.RawQuery)
// 	if eff != nil {
// 		log.Println("usrl parsequery error")
// 		log.Println(eff)
// 	}
// 	s1 := m["season"][0]
// 	log.Println(s1)
// 	ses := DBcon()
// 	defer ses.Close()
// 	MTyc := ses.DB("tvgobs").C("tvgobs")
// 	var wandavisionMedia []map[string]string
// 	b1 := bson.M{"catagory": "WandaVision", "season": `01`}
// 	b2 := bson.M{"_id": 0}
// 	errG := MTyc.Find(b1).Select(b2).Sort("episode").All(&wandavisionMedia)
// 	if errG != nil {
// 		log.Println(errG)
// 	}
// 	return c.JSON(http.StatusOK, wandavisionMedia)
// }

// func IntFalconWInterSoldierHandler(c echo.Context) error {
// 	log.Println("Starting IntFalconWInterSoldier")
// 	log.Println(" started")
// 	u, err := url.Parse(r.URL.String())
// 	if err != nil {
// 		log.Println(" url parse error")
// 		log.Println(err)
// 	}
// 	m, eff := url.ParseQuery(u.RawQuery)
// 	if eff != nil {
// 		log.Println("usrl parsequery error")
// 		log.Println(eff)
// 	}
// 	s1 := m["season"][0]
// 	log.Println(s1)
// 	ses := DBcon()
// 	defer ses.Close()
// 	MTyc := ses.DB("tvgobs").C("tvgobs")
// 	var lokiMedia []map[string]string
// 	b1 := bson.M{"catagory": "FalconWInterSoldier", "season": s1}
// 	b2 := bson.M{"_id": 0}
// 	errG := MTyc.Find(b1).Select(b2).Sort("episode").All(&lokiMedia)
// 	if errG != nil {
// 		log.Println(errG)
// 	}
// 	return c.JSON(http.StatusOK, lokiMedia)
// }

// func IntLokiHandler(c echo.Context) error {
// 	log.Println("Starting IntLoki")
// 	log.Println(" started")
// 	u, err := url.Parse(r.URL.String())
// 	if err != nil {
// 		log.Println(" url parse error")
// 		log.Println(err)
// 	}
// 	m, eff := url.ParseQuery(u.RawQuery)
// 	if eff != nil {
// 		log.Println("usrl parsequery error")
// 		log.Println(eff)
// 	}
// 	s1 := m["season"][0]
// 	log.Println(s1)
// 	ses := DBcon()
// 	defer ses.Close()
// 	MTyc := ses.DB("tvgobs").C("tvgobs")
// 	var lokiMedia []map[string]string
// 	b1 := bson.M{"catagory": "Loki", "season": s1}
// 	b2 := bson.M{"_id": 0}
// 	errG := MTyc.Find(b1).Select(b2).Sort("episode").All(&lokiMedia)
// 	if errG != nil {
// 		log.Println(errG)
// 	}
// 	log.Println(lokiMedia)
// 	return c.JSON(http.StatusOK, lokiMedia)
// }

// func IntWhatIfHandler(c echo.Context) error {
// 	log.Println("Starting IntWhatIf")
// 	log.Println(" started")
// 	u, err := url.Parse(r.URL.String())
// 	if err != nil {
// 		log.Println(" url parse error")
// 		log.Println(err)
// 	}
// 	m, eff := url.ParseQuery(u.RawQuery)
// 	if eff != nil {
// 		log.Println("usrl parsequery error")
// 		log.Println(eff)
// 	}
// 	s1 := m["season"][0]
// 	log.Println(s1)
// 	ses := DBcon()
// 	defer ses.Close()
// 	MTyc := ses.DB("tvgobs").C("tvgobs")
// 	var whatifMedia []map[string]string
// 	b1 := bson.M{"catagory": "WhatIf", "season": s1}
// 	b2 := bson.M{"_id": 0}
// 	errG := MTyc.Find(b1).Select(b2).Sort("episode").All(&whatifMedia)
// 	if errG != nil {
// 		log.Println(errG)
// 	}
// 	log.Println(whatifMedia)
// 	return c.JSON(http.StatusOK, whatifMedia)
// }

// func IntYTheLastManHandler(c echo.Context) error {
// 	log.Println("Starting IntYTheLastMan")
// 	log.Println("YTheLastMan started")
// 	u, err := url.Parse(r.URL.String())
// 	if err != nil {
// 		log.Println("YTheLastManHandler url parse error")
// 		log.Println(err)
// 	}
// 	m, eff := url.ParseQuery(u.RawQuery)
// 	if eff != nil {
// 		log.Println("YTheLastManHandler usrl parsequery error")
// 		log.Println(eff)
// 	}
// 	s1 := m["season"][0]
// 	log.Println(s1)
// 	ses := DBcon()
// 	defer ses.Close()
// 	MTyc := ses.DB("tvgobs").C("tvgobs")
// 	var YTheLastManMedia []map[string]string
// 	b1 := bson.M{"catagory": "YTheLastMan", "season": s1}
// 	b2 := bson.M{"_id": 0}
// 	errG := MTyc.Find(b1).Select(b2).Sort("episode").All(&YTheLastManMedia)
// 	if errG != nil {
// 		log.Println(errG)
// 	}
// 	log.Println(YTheLastManMedia)
// 	return c.JSON(http.StatusOK, YTheLastManMedia)
// }

// func IntFoundationHandler(c echo.Context) error {
// 	log.Println("Starting Foundation")
// 	log.Println("Foundation started")
// 	u, err := url.Parse(r.URL.String())
// 	if err != nil {
// 		log.Println("Foundation url parse error")
// 		log.Println(err)
// 	}
// 	m, eff := url.ParseQuery(u.RawQuery)
// 	if eff != nil {
// 		log.Println("Foundation usrl parsequery error")
// 		log.Println(eff)
// 	}
// 	s1 := m["season"][0]
// 	log.Println(s1)
// 	ses := DBcon()
// 	defer ses.Close()
// 	MTyc := ses.DB("tvgobs").C("tvgobs")
// 	var FoundationMedia []map[string]string
// 	b1 := bson.M{"catagory": "Foundation", "season": s1}
// 	b2 := bson.M{"_id": 0}
// 	errG := MTyc.Find(b1).Select(b2).Sort("episode").All(&FoundationMedia)
// 	if errG != nil {
// 		log.Println(errG)
// 	}
// 	log.Println(FoundationMedia)
// 	return c.JSON(http.StatusOK, FoundationMedia)
// }

// func IntVisionsHandler(c echo.Context) error {
// 	log.Println("Starting Visions")
// 	log.Println("Visions started")
// 	u, err := url.Parse(r.URL.String())
// 	if err != nil {
// 		log.Println("Visions url parse error")
// 		log.Println(err)
// 	}
// 	m, eff := url.ParseQuery(u.RawQuery)
// 	if eff != nil {
// 		log.Println("Visions usrl parsequery error")
// 		log.Println(eff)
// 	}
// 	s1 := m["season"][0]
// 	log.Println(s1)
// 	ses := DBcon()
// 	defer ses.Close()
// 	MTyc := ses.DB("tvgobs").C("tvgobs")
// 	var VisionsMedia []map[string]string
// 	b1 := bson.M{"catagory": "Visions", "season": s1}
// 	b2 := bson.M{"_id": 0}
// 	errG := MTyc.Find(b1).Select(b2).Sort("episode").All(&VisionsMedia)
// 	if errG != nil {
// 		log.Println(errG)
// 	}
// 	log.Println(VisionsMedia)
// 	return c.JSON(http.StatusOK, VisionsMedia)
// }

// func IntProdigyHandler(c echo.Context) error {
// 	log.Println("Starting Prodigy")
// 	log.Println("Prodigy started")
// 	u, err := url.Parse(r.URL.String())
// 	if err != nil {
// 		log.Println("Prodigy url parse error")
// 		log.Println(err)
// 	}
// 	m, eff := url.ParseQuery(u.RawQuery)
// 	if eff != nil {
// 		log.Println("Prodigy usrl parsequery error")
// 		log.Println(eff)
// 	}
// 	s1 := m["season"][0]
// 	log.Println(s1)
// 	ses := DBcon()
// 	defer ses.Close()
// 	MTyc := ses.DB("tvgobs").C("tvgobs")
// 	var ProdigyMedia []map[string]string
// 	b1 := bson.M{"catagory": "Prodigy", "season": s1}
// 	b2 := bson.M{"_id": 0}
// 	errG := MTyc.Find(b1).Select(b2).Sort("episode").All(&ProdigyMedia)
// 	if errG != nil {
// 		log.Println(errG)
// 	}
// 	log.Println(ProdigyMedia)
// 	return c.JSON(http.StatusOK, ProdigyMedia)
// }

// func IntTheBadBatchHandler(c echo.Context) error {
// 	log.Println("TheBadBatch started")
// 	u, err := url.Parse(r.URL.String())
// 	if err != nil {
// 		log.Println("parse url string error")
// 		log.Println(err)
// 	}
// 	m, eff := url.ParseQuery(u.RawQuery)
// 	if eff != nil {
// 		log.Println("usrl parsequery error")
// 		log.Println(eff)
// 	}
// 	s1 := m["season"][0]
// 	ses := DBcon()
// 	defer ses.Close()
// 	MTyc := ses.DB("tvgobs").C("tvgobs")
// 	var thebadbatchMedia []map[string]string
// 	b1 := bson.M{"catagory": "TheBadBatch", "season": s1}
// 	b2 := bson.M{"_id": 0}
// 	errG := MTyc.Find(b1).Select(b2).Sort("episode").All(&thebadbatchMedia)
// 	if errG != nil {
// 		log.Println(errG)
// 	}
// 	return c.JSON(http.StatusOK, thebadbatchMedia)
// }

// func IntMastersOfTheUniverseHandler(c echo.Context) error {
// 	log.Println("MastersOfTheUniverse started")
// 	u, err := url.Parse(r.URL.String())
// 	if err != nil {
// 		log.Println("url parse error")
// 		log.Println(err)
// 	}
// 	_, eff := url.ParseQuery(u.RawQuery)
// 	if eff != nil {
// 		log.Println("MastersOfTheUniverse url parse querry error")
// 		log.Println(eff)
// 	}
// 	ses := DBcon()
// 	defer ses.Close()
// 	MTyc := ses.DB("tvgobs").C("tvgobs")
// 	var mastersOfTheUniverseMedia []map[string]string
// 	b1 := bson.M{"catagory": "MastersOfTheUniverse", "season": `01`}
// 	b2 := bson.M{"_id": 0}
// 	errG := MTyc.Find(b1).Select(b2).Sort("episode").All(&mastersOfTheUniverseMedia)
// 	if errG != nil {
// 		log.Println("mastersOfTheUniverse db call error")
// 		log.Println(errG)
// 	}
// 	return c.JSON(http.StatusOK, mastersOfTheUniverseMedia)
// }

// func IntWheelOfTimeHandler(c echo.Context) error {
// 	log.Println("WheelOfTime started")
// 	u, err := url.Parse(r.URL.String())
// 	if err != nil {
// 		log.Println("url parse error")
// 		log.Println(err)
// 	}
// 	_, eff := url.ParseQuery(u.RawQuery)
// 	if eff != nil {
// 		log.Println("WheelOfTime url parse querry error")
// 		log.Println(eff)
// 	}
// 	ses := DBcon()
// 	defer ses.Close()
// 	MTyc := ses.DB("tvgobs").C("tvgobs")
// 	var wheeloftimeMedia []map[string]string
// 	b1 := bson.M{"catagory": "WheelOfTime", "season": `01`}
// 	b2 := bson.M{"_id": 0}
// 	errG := MTyc.Find(b1).Select(b2).Sort("episode").All(&wheeloftimeMedia)
// 	if errG != nil {
// 		log.Println("WheelOfTime db call error")
// 		log.Println(errG)
// 	}
// 	return c.JSON(http.StatusOK, wheeloftimeMedia)
// }

// func IntCowboyBebopHandler(c echo.Context) error {
// 	log.Println("CowboyBebop started")
// 	u, err := url.Parse(r.URL.String())
// 	if err != nil {
// 		log.Println("url parse error")
// 		log.Println(err)
// 	}
// 	_, eff := url.ParseQuery(u.RawQuery)
// 	if eff != nil {
// 		log.Println("CowboyBebop url parse querry error")
// 		log.Println(eff)
// 	}
// 	ses := DBcon()
// 	defer ses.Close()
// 	MTyc := ses.DB("tvgobs").C("tvgobs")
// 	var CowboyBebopMedia []map[string]string
// 	b1 := bson.M{"catagory": "CowboyBebop", "season": `01`}
// 	b2 := bson.M{"_id": 0}
// 	errG := MTyc.Find(b1).Select(b2).Sort("episode").All(&CowboyBebopMedia)
// 	if errG != nil {
// 		log.Println("CowboyBebop db call error")
// 		log.Println(errG)
// 	}
// 	return c.JSON(http.StatusOK, CowboyBebopMedia)
// }

// func IntHawkeyeHandler(c echo.Context) error {
// 	log.Println("Hawkeye started")
// 	u, err := url.Parse(r.URL.String())
// 	if err != nil {
// 		log.Println("url parse error")
// 		log.Println(err)
// 	}
// 	_, eff := url.ParseQuery(u.RawQuery)
// 	if eff != nil {
// 		log.Println("Hawkeye url parse querry error")
// 		log.Println(eff)
// 	}
// 	ses := DBcon()
// 	defer ses.Close()
// 	MTyc := ses.DB("tvgobs").C("tvgobs")
// 	var HawkeyeMedia []map[string]string
// 	b1 := bson.M{"catagory": "Hawkeye", "season": `01`}
// 	b2 := bson.M{"_id": 0}
// 	errG := MTyc.Find(b1).Select(b2).Sort("episode").All(&HawkeyeMedia)
// 	if errG != nil {
// 		log.Println("Hawkeye db call error")
// 		log.Println(errG)
// 	}
// 	return c.JSON(http.StatusOK, HawkeyeMedia)
// }

// func IntBookOfBobaFettHandler(c echo.Context) error {
// 	log.Println("BookOfBobaFett started")
// 	u, err := url.Parse(r.URL.String())
// 	if err != nil {
// 		log.Println("url parse error")
// 		log.Println(err)
// 	}
// 	_, eff := url.ParseQuery(u.RawQuery)
// 	if eff != nil {
// 		log.Println("BookOfBobaFett url parse querry error")
// 		log.Println(eff)
// 	}
// 	ses := DBcon()
// 	defer ses.Close()
// 	MTyc := ses.DB("tvgobs").C("tvgobs")
// 	var BookOfBobaFettMedia []map[string]string
// 	b1 := bson.M{"catagory": "BookOfBobaFett", "season": `01`}
// 	b2 := bson.M{"_id": 0}
// 	errG := MTyc.Find(b1).Select(b2).Sort("episode").All(&BookOfBobaFettMedia)
// 	if errG != nil {
// 		log.Println("BookOfBobaFett db call error")
// 		log.Println(errG)
// 	}
// 	return c.JSON(http.StatusOK, BookOfBobaFettMedia)
// }

// func IntReacherHandler(c echo.Context) error {
// 	log.Println("Reacher started")
// 	u, err := url.Parse(r.URL.String())
// 	if err != nil {
// 		log.Println("url parse error")
// 		log.Println(err)
// 	}
// 	_, eff := url.ParseQuery(u.RawQuery)
// 	if eff != nil {
// 		log.Println("Reacher url parse querry error")
// 		log.Println(eff)
// 	}
// 	ses := DBcon()
// 	defer ses.Close()
// 	MTyc := ses.DB("tvgobs").C("tvgobs")
// 	var ReacherMedia []map[string]string
// 	b1 := bson.M{"catagory": "Reacher", "season": `01`}
// 	b2 := bson.M{"_id": 0}
// 	errG := MTyc.Find(b1).Select(b2).Sort("episode").All(&ReacherMedia)
// 	if errG != nil {
// 		log.Println("Reacher db call error")
// 		log.Println(errG)
// 	}
// 	return c.JSON(http.StatusOK, ReacherMedia)
// }

// func IntHaloHandler(c echo.Context) error {
// 	log.Println("Halo started")
// 	u, err := url.Parse(r.URL.String())
// 	if err != nil {
// 		log.Println("url parse error")
// 		log.Println(err)
// 	}
// 	_, eff := url.ParseQuery(u.RawQuery)
// 	if eff != nil {
// 		log.Println("Halo url parse querry error")
// 		log.Println(eff)
// 	}
// 	ses := DBcon()
// 	defer ses.Close()
// 	MTyc := ses.DB("tvgobs").C("tvgobs")
// 	var HaloMedia []map[string]string
// 	b1 := bson.M{"catagory": "Halo", "season": `01`}
// 	b2 := bson.M{"_id": 0}
// 	errG := MTyc.Find(b1).Select(b2).Sort("episode").All(&HaloMedia)
// 	if errG != nil {
// 		log.Println("Halo db call error")
// 		log.Println(errG)
// 	}
// 	return c.JSON(http.StatusOK, HaloMedia)
// }

// func IntMoonKnightHandler(c echo.Context) error {
// 	log.Println("MoonKnight started")
// 	u, err := url.Parse(r.URL.String())
// 	if err != nil {
// 		log.Println("url parse error")
// 		log.Println(err)
// 	}
// 	_, eff := url.ParseQuery(u.RawQuery)
// 	if eff != nil {
// 		log.Println("MoonKnight url parse querry error")
// 		log.Println(eff)
// 	}
// 	ses := DBcon()
// 	defer ses.Close()
// 	MTyc := ses.DB("tvgobs").C("tvgobs")
// 	var MoonKnightMedia []map[string]string
// 	b1 := bson.M{"catagory": "MoonKnight", "season": `01`}
// 	b2 := bson.M{"_id": 0}
// 	errG := MTyc.Find(b1).Select(b2).Sort("episode").All(&MoonKnightMedia)
// 	if errG != nil {
// 		log.Println("MoonKnight db call error")
// 		log.Println(errG)
// 	}
// 	return c.JSON(http.StatusOK, MoonKnightMedia)
// }

// func IntStrangeNewWorldsHandler(c echo.Context) error {
// 	log.Println("StrangeNewWorlds started")
// 	u, err := url.Parse(r.URL.String())
// 	if err != nil {
// 		log.Println("url parse error")
// 		log.Println(err)
// 	}
// 	_, eff := url.ParseQuery(u.RawQuery)
// 	if eff != nil {
// 		log.Println("StrangeNewWorlds url parse querry error")
// 		log.Println(eff)
// 	}
// 	ses := DBcon()
// 	defer ses.Close()
// 	MTyc := ses.DB("tvgobs").C("tvgobs")
// 	var StrangeNewWorldsMedia []map[string]string
// 	b1 := bson.M{"catagory": "StrangeNewWorlds", "season": `01`}
// 	b2 := bson.M{"_id": 0}
// 	errG := MTyc.Find(b1).Select(b2).Sort("episode").All(&StrangeNewWorldsMedia)
// 	if errG != nil {
// 		log.Println("StrangeNewWorlds db call error")
// 		log.Println(errG)
// 	}
// 	return c.JSON(http.StatusOK, StrangeNewWorldsMedia)
// }

// func IntArnoldHandler(c echo.Context) error {
// 	log.Println("IntArnoldHandler started")
// 	ses := DBcon()
// 	defer ses.Close()
// 	MTc := ses.DB("moviegobs").C("moviegobs")
// 	var arnoldMedia []map[string]string
// 	b1 := bson.M{"catagory": "Arnold"}
// 	b2 := bson.M{"_id": 0}
// 	err := MTc.Find(b1).Select(b2).All(&arnoldMedia)
// 	if err != nil {
// 		log.Printf("arnoldMedia db call error")
// 		log.Println(err)
// 	}
// 	return c.JSON(http.StatusOK, arnoldMedia)
// }

// func IntPrehistoricPlanetHandler(c echo.Context) error {
// 	log.Println("PrehistoricPlanet started")
// 	u, err := url.Parse(r.URL.String())
// 	if err != nil {
// 		log.Println("url parse error")
// 		log.Println(err)
// 	}
// 	_, eff := url.ParseQuery(u.RawQuery)
// 	if eff != nil {
// 		log.Println("PrehistoricPlanet url parse querry error")
// 		log.Println(eff)
// 	}
// 	ses := DBcon()
// 	defer ses.Close()
// 	MTyc := ses.DB("tvgobs").C("tvgobs")
// 	var PrehistoricPlanetMedia []map[string]string
// 	b1 := bson.M{"catagory": "PrehistoricPlanet", "season": `01`}
// 	b2 := bson.M{"_id": 0}
// 	errG := MTyc.Find(b1).Select(b2).Sort("episode").All(&PrehistoricPlanetMedia)
// 	if errG != nil {
// 		log.Println("PrehistoricPlanet db call error")
// 		log.Println(errG)
// 	}
// 	return c.JSON(http.StatusOK, PrehistoricPlanetMedia)
// }

// func IntObiWanKenobiHandler(c echo.Context) error {
// 	log.Println("ObiWanKenobi started")
// 	u, err := url.Parse(r.URL.String())
// 	if err != nil {
// 		log.Println("url parse error")
// 		log.Println(err)
// 	}
// 	_, eff := url.ParseQuery(u.RawQuery)
// 	if eff != nil {
// 		log.Println("ObiWanKenobi url parse querry error")
// 		log.Println(eff)
// 	}
// 	ses := DBcon()
// 	defer ses.Close()
// 	MTyc := ses.DB("tvgobs").C("tvgobs")
// 	var ObiWanKenobiMedia []map[string]string
// 	b1 := bson.M{"catagory": "ObiWanKenobi", "season": `01`}
// 	b2 := bson.M{"_id": 0}
// 	errG := MTyc.Find(b1).Select(b2).Sort("episode").All(&ObiWanKenobiMedia)
// 	if errG != nil {
// 		log.Println("ObiWanKenobi db call error")
// 		log.Println(errG)
// 	}
// 	return c.JSON(http.StatusOK, ObiWanKenobiMedia)
// }

// func IntMSMarvelHandler(c echo.Context) error {
// 	log.Println("MSMarvel started")
// 	u, err := url.Parse(r.URL.String())
// 	if err != nil {
// 		log.Println("url parse error")
// 		log.Println(err)
// 	}
// 	_, eff := url.ParseQuery(u.RawQuery)
// 	if eff != nil {
// 		log.Println("MSMarvel url parse querry error")
// 		log.Println(eff)
// 	}
// 	ses := DBcon()
// 	defer ses.Close()
// 	MTyc := ses.DB("tvgobs").C("tvgobs")
// 	var MSMarvelMedia []map[string]string
// 	b1 := bson.M{"catagory": "MSMarvel", "season": `01`}
// 	b2 := bson.M{"_id": 0}
// 	errG := MTyc.Find(b1).Select(b2).Sort("episode").All(&MSMarvelMedia)
// 	if errG != nil {
// 		log.Println("MSMarvel db call error")
// 		log.Println(errG)
// 	}
// 	return c.JSON(http.StatusOK, MSMarvelMedia)
// }

// func IntIAmGrootHandler(c echo.Context) error {
// 	log.Println("IAmGroot started")
// 	u, err := url.Parse(r.URL.String())
// 	if err != nil {
// 		log.Println("url parse error")
// 		log.Println(err)
// 	}
// 	_, eff := url.ParseQuery(u.RawQuery)
// 	if eff != nil {
// 		log.Println("IAmGroot url parse querry error")
// 		log.Println(eff)
// 	}
// 	ses := DBcon()
// 	defer ses.Close()
// 	MTyc := ses.DB("tvgobs").C("tvgobs")
// 	var IAmGrootMedia []map[string]string
// 	b1 := bson.M{"catagory": "IAmGroot", "season": `01`}
// 	b2 := bson.M{"_id": 0}
// 	errG := MTyc.Find(b1).Select(b2).Sort("episode").All(&IAmGrootMedia)
// 	if errG != nil {
// 		log.Println("IAmGroot db call error")
// 		log.Println(errG)
// 	}
// 	return c.JSON(http.StatusOK, IAmGrootMedia)
// }

// func IntSheHulkHandler(c echo.Context) error {
// 	log.Println("SheHulk started")
// 	u, err := url.Parse(r.URL.String())
// 	if err != nil {
// 		log.Println("url parse error")
// 		log.Println(err)
// 	}
// 	_, eff := url.ParseQuery(u.RawQuery)
// 	if eff != nil {
// 		log.Println("SheHulk url parse querry error")
// 		log.Println(eff)
// 	}
// 	ses := DBcon()
// 	defer ses.Close()
// 	MTyc := ses.DB("tvgobs").C("tvgobs")
// 	var SheHulkMedia []map[string]string
// 	b1 := bson.M{"catagory": "SheHulk", "season": `01`}
// 	b2 := bson.M{"_id": 0}
// 	errG := MTyc.Find(b1).Select(b2).Sort("episode").All(&SheHulkMedia)
// 	if errG != nil {
// 		log.Println("SheHulkMedia db call error")
// 		log.Println(errG)
// 	}
// 	return c.JSON(http.StatusOK, SheHulkMedia)
// }

// func IntHouseOfTheDragonHandler(c echo.Context) error {
// 	log.Println("HouseOfTheDragon started")
// 	u, err := url.Parse(r.URL.String())
// 	if err != nil {
// 		log.Println("url parse error")
// 		log.Println(err)
// 	}
// 	_, eff := url.ParseQuery(u.RawQuery)
// 	if eff != nil {
// 		log.Println("HouseOfTheDragon url parse querry error")
// 		log.Println(eff)
// 	}
// 	ses := DBcon()
// 	defer ses.Close()
// 	MTyc := ses.DB("tvgobs").C("tvgobs")
// 	var HouseOfTheDragonMedia []map[string]string
// 	b1 := bson.M{"catagory": "HouseOfTheDragon", "season": `01`}
// 	b2 := bson.M{"_id": 0}
// 	errG := MTyc.Find(b1).Select(b2).Sort("episode").All(&HouseOfTheDragonMedia)
// 	if errG != nil {
// 		log.Println("HouseOfTheDragonMedia db call error")
// 		log.Println(errG)
// 	}
// 	return c.JSON(http.StatusOK, HouseOfTheDragonMedia)
// }

// func IntTheLordOfTheRingsTheRingsOfPowerHandler(c echo.Context) error {
// 	log.Println("TheLordOfTheRingsTheRingsOfPower started")
// 	u, err := url.Parse(r.URL.String())
// 	if err != nil {
// 		log.Println("url parse error")
// 		log.Println(err)
// 	}
// 	_, eff := url.ParseQuery(u.RawQuery)
// 	if eff != nil {
// 		log.Println("TheLordOfTheRingsTheRingsOfPower url parse querry error")
// 		log.Println(eff)
// 	}
// 	ses := DBcon()
// 	defer ses.Close()
// 	MTyc := ses.DB("tvgobs").C("tvgobs")
// 	var TheLordOfTheRingsTheRingsOfPowerMedia []map[string]string
// 	b1 := bson.M{"catagory": "TheLordOfTheRingsTheRingsOfPower", "season": `01`}
// 	b2 := bson.M{"_id": 0}
// 	errG := MTyc.Find(b1).Select(b2).Sort("episode").All(&TheLordOfTheRingsTheRingsOfPowerMedia)
// 	if errG != nil {
// 		log.Println("TheLordOfTheRingsTheRingsOfPowerMedia db call error")
// 		log.Println(errG)
// 	}
// 	return c.JSON(http.StatusOK, TheLordOfTheRingsTheRingsOfPowerMedia)
// }

// func IntAndorHandler(c echo.Context) error {
// 	log.Println("Andor started")
// 	u, err := url.Parse(r.URL.String())
// 	if err != nil {
// 		log.Println("url parse error")
// 		log.Println(err)
// 	}
// 	_, eff := url.ParseQuery(u.RawQuery)
// 	if eff != nil {
// 		log.Println("Andor url parse querry error")
// 		log.Println(eff)
// 	}
// 	ses := DBcon()
// 	defer ses.Close()
// 	MTyc := ses.DB("tvgobs").C("tvgobs")
// 	var AndorMedia []map[string]string
// 	b1 := bson.M{"catagory": "Andor", "season": `01`}
// 	b2 := bson.M{"_id": 0}
// 	errG := MTyc.Find(b1).Select(b2).Sort("episode").All(&AndorMedia)
// 	if errG != nil {
// 		log.Println("AndorMedia db call error")
// 		log.Println(errG)
// 	}
// 	return c.JSON(http.StatusOK, AndorMedia)
// }

// func IntNightSkyHandler(c echo.Context) error {
// 	log.Println("NightSky started")
// 	u, err := url.Parse(r.URL.String())
// 	if err != nil {
// 		log.Println("url parse error")
// 		log.Println(err)
// 	}
// 	_, eff := url.ParseQuery(u.RawQuery)
// 	if eff != nil {
// 		log.Println("NightSky url parse querry error")
// 		log.Println(eff)
// 	}
// 	ses := DBcon()
// 	defer ses.Close()
// 	MTyc := ses.DB("tvgobs").C("tvgobs")
// 	var NightSkyMedia []map[string]string
// 	b1 := bson.M{"catagory": "NightSky", "season": `01`}
// 	b2 := bson.M{"_id": 0}
// 	errG := MTyc.Find(b1).Select(b2).Sort("episode").All(&NightSkyMedia)
// 	if errG != nil {
// 		log.Println("NightSky db call error")
// 		log.Println(errG)
// 	}
// 	return c.JSON(http.StatusOK, NightSkyMedia)
// }

// func IntTalesOfTheJediHandler(c echo.Context) error {
// 	log.Println("TalesOfTheJedi started")
// 	u, err := url.Parse(r.URL.String())
// 	if err != nil {
// 		log.Println("url parse error")
// 		log.Println(err)
// 	}
// 	_, eff := url.ParseQuery(u.RawQuery)
// 	if eff != nil {
// 		log.Println("TalesOfTheJedi url parse querry error")
// 		log.Println(eff)
// 	}
// 	ses := DBcon()
// 	defer ses.Close()
// 	MTyc := ses.DB("tvgobs").C("tvgobs")
// 	var TalesOfTheJediMedia []map[string]string
// 	b1 := bson.M{"catagory": "TalesOfTheJedi", "season": `01`}
// 	b2 := bson.M{"_id": 0}
// 	errG := MTyc.Find(b1).Select(b2).Sort("episode").All(&TalesOfTheJediMedia)
// 	if errG != nil {
// 		log.Println("TalesOfTheJedi db call error")
// 		log.Println(errG)
// 	}
// 	return c.JSON(http.StatusOK, TalesOfTheJediMedia)
// }

// func setupLogging() {
// 	logfile := os.Getenv("MTV_LOG_BASE_PATH") + "/moviegobsServer.log"
// 	// If the file doesn't exist, create it or append to the file
// 	file, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.SetOutput(file)
// 	log.Println("Logging started")
// }

// func init() {
// 	setupLogging()
// 	// movgo.MOVSetup()
// 	// tvgo.TVSetUp()

// }

// func main() {
// 	r := mux.NewRouter()
// 	s := r.PathPrefix("/static").Subrouter()
// 	r.HandleFunc("/IntAction", IntActionHandler)
// 	r.HandleFunc("/IntCartoons", IntCartoonsHandler)
// 	r.HandleFunc("/IntComedy", IntComedyHandler)
// 	r.HandleFunc("/IntDrama", IntDramaHandler)
// 	r.HandleFunc("/IntFantasy", IntFantasyHandler)
// 	r.HandleFunc("/IntGodzilla", IntGodzillaHandler)
// 	r.HandleFunc("/IntHarryPotter", IntHarryPotterHandler)
// 	r.HandleFunc("/IntIndianaJones", IntIndianaJonesHandler)
// 	r.HandleFunc("/IntMenInBlack", IntMenInBlackHandler)
// 	r.HandleFunc("/IntMisc", IntMiscHandler)
// 	r.HandleFunc("/IntJohnWayne", IntJohnWayneHandler)
// 	r.HandleFunc("/IntJurassicPark", IntJurassicParkHandler)
// 	r.HandleFunc("/IntKingsMan", IntKingsManHandler)
// 	r.HandleFunc("/IntSciFi", IntSciFiHandler)
// 	r.HandleFunc("/IntStarTrek", IntStarTrekHandler)
// 	r.HandleFunc("/IntStarWars", IntStarWarsHandler)
// 	r.HandleFunc("/IntSuperHeros", IntSuperHerosHandler)
// 	r.HandleFunc("/IntTremors", IntTremorsHandler)
// 	r.HandleFunc("/IntJohnWick", IntJohnWickHandler)
// 	r.HandleFunc("/IntPirates", IntPiratesHandler)
// 	r.HandleFunc("/IntBruceWillis", IntBruceWillisHandler)
// 	r.HandleFunc("/IntRiddick", IntRiddickHandler)
// 	r.HandleFunc("/IntTomCruize", IntTomCruizeHandler)
// 	r.HandleFunc("/IntXMen", IntXMenHandler)
// 	r.HandleFunc("/IntDocumentary", IntDocumentaryHandler)
// 	r.HandleFunc("/IntTheRock", IntTheRockHandler)
// 	r.HandleFunc("/MovUpdate", MovUpdateHandler)
// 	r.HandleFunc("/IntNicolasCage", IntNicolasCageHandler)
// 	r.HandleFunc("/IntArnold", IntArnoldHandler)
// 	r.HandleFunc("/IntJamesBond", IntJamesBondHandler)
// 	r.HandleFunc("/IntTransformers", IntTransformersHandler)
// 	//TVGOBS_SETUP
// 	r.HandleFunc("/IntFalconWInterSoldier", IntFalconWInterSoldierHandler)
// 	r.HandleFunc("/IntAlteredCarbon", IntAlteredCarbonHandler)
// 	r.HandleFunc("/IntAlienWorlds", IntAlienWorldsHandler)
// 	r.HandleFunc("/IntDiscovery", IntDiscoveryHandler)
// 	r.HandleFunc("/IntEnterprise", IntEnterpriseHandler)
// 	r.HandleFunc("/IntForAllManKind", IntForAllManKindHandler)
// 	r.HandleFunc("/IntLostInSpace", IntLostInSpaceHandler)
// 	r.HandleFunc("/IntLowerDecks", IntLowerDecksHandler)
// 	r.HandleFunc("/IntMandalorian", IntMandalorianHandler)
// 	r.HandleFunc("/IntOrville", IntOrvilleHandler)
// 	r.HandleFunc("/IntPicard", IntPicardHandler)
// 	r.HandleFunc("/IntRaisedByWolves", IntRaisedByWolvesHandler)
// 	r.HandleFunc("/IntSTTV", IntSTTVHandler)
// 	r.HandleFunc("/IntTNG", IntTNGHandler)
// 	r.HandleFunc("/IntVoyager", IntVoyagerHandler)
// 	r.HandleFunc("/IntWandaVision", IntWandaVisionHandler)
// 	r.HandleFunc("/IntMastersOfTheUniverse", IntMastersOfTheUniverseHandler)
// 	r.HandleFunc("/IntLoki", IntLokiHandler)
// 	r.HandleFunc("/IntTheBadBatch", IntTheBadBatchHandler)
// 	r.HandleFunc("/IntWhatIf", IntWhatIfHandler)
// 	r.HandleFunc("/IntYTheLastMan", IntYTheLastManHandler)
// 	r.HandleFunc("/IntFoundation", IntFoundationHandler)
// 	r.HandleFunc("/IntVisions", IntVisionsHandler)
// 	r.HandleFunc("/IntProdigy", IntProdigyHandler)
// 	r.HandleFunc("/IntWheelOfTime", IntWheelOfTimeHandler)
// 	r.HandleFunc("/IntCowboyBebop", IntCowboyBebopHandler)
// 	r.HandleFunc("/IntHawkeye", IntHawkeyeHandler)
// 	r.HandleFunc("/IntBookOfBobaFett", IntBookOfBobaFettHandler)
// 	r.HandleFunc("/IntReacher", IntReacherHandler)
// 	r.HandleFunc("/IntHalo", IntHaloHandler)
// 	r.HandleFunc("/IntMoonKnight", IntMoonKnightHandler)
// 	r.HandleFunc("/IntStrangeNewWorlds", IntStrangeNewWorldsHandler)
// 	r.HandleFunc("/IntPrehistoricPlanet", IntPrehistoricPlanetHandler)
// 	r.HandleFunc("/IntObiWanKenobi", IntObiWanKenobiHandler)
// 	r.HandleFunc("/IntMSMarvel", IntMSMarvelHandler)
// 	r.HandleFunc("/IntIAmGroot", IntIAmGrootHandler)
// 	r.HandleFunc("/IntSheHulk", IntSheHulkHandler)

// 	r.HandleFunc("/IntHouseOfTheDragon", IntHouseOfTheDragonHandler)
// 	r.HandleFunc("/IntTheLordOfTheRingsTheRingsOfPower", IntTheLordOfTheRingsTheRingsOfPowerHandler)
// 	r.HandleFunc("/IntAndor", IntAndorHandler)
// 	r.HandleFunc("/IntNightSky", IntNightSkyHandler)
// 	r.HandleFunc("/IntTalesOfTheJedi", IntTalesOfTheJediHandler)

// 	r.HandleFunc("/TVUpdate", TVUpdateHandler)
// 	r.HandleFunc("/playMediaReact", playMediaReactHandler)

// 	s.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(""))))
// 	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("/media/"))))
// 	http.ListenAndServe(":8888", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
// 		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
// 		handlers.AllowedOrigins([]string{"*"}))(r))
// }

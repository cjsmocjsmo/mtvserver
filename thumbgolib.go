package main

import (
	"fmt"
	//because I want it
	_ "image/jpeg"
	_ "image/png"
	"log"

	"github.com/disintegration/imaging"
	"github.com/globalsign/mgo/bson"

	// "math/rand"
	"os"
	"path"
	// "strconv"
	// "time"
)

func myPathSplit(myPath string) (DirPath string, BaseNAme string, MOvName string, Ext string) {
	DirPath, BaseNAme = path.Split(myPath)
	Ext = BaseNAme[3:]
	factor := len(BaseNAme) - 4
	MOvName = BaseNAme[:factor]
	return
}

func getServerAddr() (addr string) {
	addr = os.Getenv("MTV_SERVER_ADDRESS")
	return
}

func getServerPort() (port string) {
	port = os.Getenv("MTV_SERVER_PORT")
	return
}

func getThumbPath() (tpath string) {
	tpath = os.Getenv("MTV_THUMBNAIL_PIC_PATH")
	return
}

//CreateMoviesThumbnail exported to setup
func CreateMoviesThumbnail(p string) (ThumbINFO ThumbInFo) {
	fmt.Println(p)
	dirpath, basepath, movname, ext := myPathSplit(p)
	MSA := getServerAddr()
	MSP := getServerPort()
	MTPP := getThumbPath()
	// BP := "/" + url.QueryEscape(basepath)
	// thumbpathtwo := MSA + ":" + MSP + MTPP + BP
	// thumbpathone := "./static/" + basepath
	var BP string = "/" + basepath
	var thumbpathtwo string = MSA + ":" + MSP + MTPP + BP
	var thumbpathone string = "./static/" + basepath

	// ThumbINFO.ID = bson.NewObjectId()
	ThumbINFO.MovName = movname
	ThumbINFO.BasePath = basepath
	ThumbINFO.DirPATH = dirpath
	ThumbINFO.HTTPThumbPath = thumbpathtwo
	ThumbINFO.ThumbPath = thumbpathone
	ThumbINFO.ThumbID = UUID()

	if ext == ".txt" {
		fmt.Print("what the fuck a text file remove it")
		os.Remove(p)
	} else if ext == ".srt" {
		os.Remove(p)
	} else {

		_, err := os.Stat(thumbpathone)
		if err == nil {
			log.Printf("FILE %s EXISTS", thumbpathone)
		} else if os.IsNotExist(err) {
			pic, err := imaging.Open(p)
			if err != nil {
				log.Printf("\n this is file Open error jpgthumb %v \n", p)
			}
			thumbImage := imaging.Resize(pic, 0, 250, imaging.Lanczos)
			err = imaging.Save(thumbImage, thumbpathone)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			log.Printf("file %s stat error: %v", thumbpathone, err)
		}
		cmtses := MovDBcon()
		defer cmtses.Close()
		CMTc := cmtses.DB("movbsthumb").C("movbsthumb")
		err = CMTc.Insert(&ThumbINFO)
		if err != nil {
			fmt.Println(err)
		}
	}
	return
}

//NoArtList exported to setup
var NoArtList []string

//a comment
//FindPicPaths exported to setup
func FindPicPaths(mpath string, noartpicpath string) (result, result2 string) {
	_, _, movename, _ := myPathSplit(mpath)
	Tses := MovDBcon()
	defer Tses.Close()
	MTc := Tses.DB("movbsthumb").C("movbsthumb")
	b1 := bson.M{"movname": movename}
	b2 := bson.M{"_Id": 0}
	var ThumbI []map[string]string
	err := MTc.Find(b1).Select(b2).All(&ThumbI)
	if err != nil {
		log.Println(err)
	}
	LenI := len(ThumbI)
	// fmt.Printf("THIS IS THUMBI %s \n", ThumbI)
	// var result string
	// var result2 string
	if LenI == 0 {
		NoArtList = append(NoArtList, mpath)
		result = noartpicpath
		result2 = noartpicpath
	} else {
		result = ThumbI[0]["thumbpath"]
		result2 = ThumbI[0]["httpthumbpath"]
	}
	fmt.Printf("this is result %s", result)
	return
}

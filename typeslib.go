package main

import (

)
type TVShowInfoS struct {
	FilePath      string        `bson:"filepath"`
	Catagory      string        `bson:"catagory"`
	MediaID       string        `bson:"MediaID"`
	Genre         string        `bson:"genre"`
	Season        string        `bson:"season"`
	Episode       string        `bson:"episode"`
	Title         string        `bson:"title"`
	Series        string        `bson:"series"`
	TVShowPicPath string        `bson:"tvshowpicpath"`
	ThumbPath     string        `bson:"thumbpath"`
	TvFSPath      string        `bson:"tvfspath"`
}
type MOVI struct {
	DirPath       string        `bson:"dirpath"`
	Filepath      string        `bson:"filepath"`
	MediaID       string        `bson:"mediaid"`
	Movname       string        `bson:"movname"`
	Genre         string        `bson:"genre"`
	Catagory      string        `bson:"catagory"`
	MovFSPath     string        `bson:"movfspath"`
	ThumbPath     string        `bson:"thumbpath"`
	HTTPThumbPath string        `bson:"httpthumbpath"`
	MovYear       string        `bson:"movyear"`
}

type ThumbInFo struct {
	MovName       string        `bson:"movname"`
	BasePath      string        `bson:"baspath"`
	DirPATH       string        `bson:"dirpath"`
	ThumbPath     string        `bson:"thumbpath"`
	ThumbID       string        `bson:"thumbid"`
	HTTPThumbPath string        `bson:"httpthumbpath"`
}


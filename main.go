package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var db *mgo.Database

// COLLECTION is
const (
	COLLECTION = "Wallpaper"
)

type (
	wallpaperModel struct {
		ID            bson.ObjectId `bson:"_id" json:"id"`
		Thumbnail     string        `bson:"thumbnail" json:"thumbnail"`
		Original      string        `bson:"original" json:"original"`
		TotalLike     string        `bson:"total_like" json:"total_like"`
		TotalView     string        `bson:"total_view" json:"total_view"`
		TotalDownload string        `bson:"total_download" json:"total_download"`
		Category      string        `bson:"category" json:"category"`
		Name          string        `bson:"name" json:"name"`
		Model         string        `bson:"model" json:"model"`
	}
)

func init() {
	session, err := mgo.Dial("45.77.47.217:30590")
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB("admin")
}

func main() {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/wallpapers", fetchAllWallpaper)
	}
	router.Run(":8081")
}

func fetchAllWallpaper(c *gin.Context) {
	var wallpapers []wallpaperModel
	skipStr := c.Query("skip")
	skipInt, _ := strconv.Atoi(skipStr)
	err := db.C(COLLECTION).Find(bson.M{}).Skip(skipInt).Limit(10).All(&wallpapers)
	if len(wallpapers) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No wallpapers found!"})
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": "Internal Server Error"})
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": wallpapers})
}

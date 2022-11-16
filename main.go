package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type NewBlackPinkSongs struct {
	ID       int    `json:"id"`
	SongName string `json:"SongTitle"`
}

var blackpinkSongs = []NewBlackPinkSongs{
	{ID: 1, SongName: "Whistle"},
	{ID: 2, SongName: "Boombyah"},
	{ID: 3, SongName: "Ready for Love"},
	{ID: 4, SongName: "Pretty Savage"},
	{ID: 5, SongName: "Typa girl"},
	{ID: 6, SongName: "Pink Venom"},
	{ID: 7, SongName: "Shut Down"},
	{ID: 8, SongName: "Love to Hate"},
}

func getBlackPinkSongs(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, blackpinkSongs)
}

func main() {
	router := gin.Default()
	router.GET("/bpsongs", getBlackPinkSongs)
	router.Run("")
}

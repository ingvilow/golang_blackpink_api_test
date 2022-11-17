package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
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

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000/"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE", "GET"},
		AllowHeaders: []string{"Content-Type, access-control-allow-origin, access-control-allow-headers"},
	}))
	dataRoutes := router.Group("api/item")
	{
		dataRoutes.GET("/bpsongs", getBlackPinkSongs)
		/* dataRoutes.POST("/", dataController.Insert)
		   dataRoutes.GET("/:id", dataController.FindByID)
		   dataRoutes.PUT("/:id", dataController.Update)
		   dataRoutes.DELETE("/:id", dataController.Delete) */
	}
	router.Run()
}

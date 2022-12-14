package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// allsongs
type NewBlackPinkSongs struct {
	ID       int    `json:"id"`
	SongName string `json:"SongTitle"`
}

// all blackpink members
type BlackPinkMembers struct {
	ID   int    `json:"idMember"`
	Name string `json:"memebersName"`
	Age  int    `json:"Age"`
}

// get Info about BP memders by ID
type FunFactsMemberById struct {
	ID               string `json:"idMember"`
	FactsAboutMember string `json:"aboutMember"`
}

// get All Albums
type AllAlbumList struct {
	Name string `json:"albumName"`
	Date string `json:""`
}

var AlbumsList = []AllAlbumList{
	{Name: "๐ Square One", Date: "August 8, 2016"},
	{Name: "๐ Square Two", Date: "November 1, 2016"},
	{Name: "๐ โAs If It Is Your Lastโ (Digital Single)", Date: "June 22, 2017"},
	{Name: "๐ BLACKPINK (Japanese Album)", Date: "August 30, 2017"},
	{Name: "๐ Re:BLACKPINK (Japanese Album)"},
	{Name: "๐ Square Up"},
	{Name: "๐ โKiss and Make Upโ (with Dua Lipa)"},
	{Name: "๐ โSOLOโ "},
	{Name: "๐ BLACKPINK In Your Area (Japanese Album)"},
	{Name: "๐ BLACKPINK In Your Area (Japanese Album)"},
	{Name: "๐ Kill This Love (Japanese Album)"},
	{Name: "๐ โSour Candyโ (with Lady Gaga)"},
	{Name: "๐ โHow You Like Thatโ "},
	{Name: "๐ โIce Creamโ (with Selena Gomez)"},
	{Name: "๐ The Album"},
	{Name: "๐ โOn The Groundโ "},
	{Name: "๐ โLalisaโ "},
	{Name: "๐ โMoneyโ"},
	{Name: "๐ โBorn Pinkโ "},
	{Name: "๐ โLiar (Jisoo solo on stage while doing concert tour)โ"},
}

var MembersNames = []BlackPinkMembers{
	{ID: 1, Name: "Jisoo", Age: 25},
	{ID: 2, Name: "Jennie", Age: 24},
	{ID: 3, Name: "Rose", Age: 23},
	{ID: 4, Name: "Lisa", Age: 23},
}

var facts = []FunFactsMemberById{
	{ID: "1", FactsAboutMember: "some fact"},
	{ID: "2", FactsAboutMember: "some fact2"},
	{ID: "3", FactsAboutMember: "some fact3"},
	{ID: "4", FactsAboutMember: "some fact4"},
	{ID: "1", FactsAboutMember: "some fact65"},
	{ID: "2", FactsAboutMember: "some fact778"},
	{ID: "3", FactsAboutMember: "some fact789"},
	{ID: "4", FactsAboutMember: "some fact2346"},
	{ID: "1", FactsAboutMember: "some fact12134ewdsf"},
	{ID: "2", FactsAboutMember: "some factefefefdsc"},
	{ID: "3", FactsAboutMember: "some factrththfvvsd"},
	{ID: "4", FactsAboutMember: "some factadfghregewcfsdbfgrefecdsggtrhtg	fewf"},
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

func getAllAlbums(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, AlbumsList)
}

func getBlackPinkMembers(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, MembersNames)
}
func getBlackPinkSongs(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, blackpinkSongs)
}

// at first we need to go though all range of facts and take it by indexes
func getBlackPinkFactsById(id string) *FunFactsMemberById {

	for t, i := range facts {
		if i.ID == id {

			return &facts[t]
		}
	}
	return nil
}

// here we finally get members fact by id and everything works fine
func getById(context *gin.Context) {
	id := context.Param("id")
	fact := getBlackPinkFactsById(id)

	context.IndentedJSON(http.StatusOK, fact)
}

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://127.0.0.1"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE", "GET"},
		AllowHeaders: []string{"Content-Type, access-control-allow-origin, access-control-allow-headers"},
	}))
	dataRoutes := router.Group("api")
	{
		dataRoutes.GET("/bpsongs", getBlackPinkSongs)
		dataRoutes.GET("/bpNames", getBlackPinkMembers)
		dataRoutes.GET("/bpFacts/id", getById)
		dataRoutes.GET("/allAlbums", getAllAlbums)
		/*  dataRoutes.GET("/:id", dataController.FindByID)
		dataRoutes.PUT("/:id", dataController.Update)
		dataRoutes.DELETE("/:id", dataController.Delete) */
	}
	router.Run("localhost:8080")
}

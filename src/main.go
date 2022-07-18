// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth
package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/adglad/web-service-gin/morestrings"
	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
	"github.com/kr/pretty"
	"github.com/swaggo/files"       // swagger embed files
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
	"net/http"
	"os"

	"github.com/adglad/web-service-gin/docs" // docs is generated by Swag CLI, you have to import it.
)

type Fiveletterword struct {
	WORD string `json:"word"`
	L1   string `json:"l1"`
	L2   string `json:"l2"`
	L3   string `json:"l3"`
	L4   string `json:"l4"`
	L5   string `json:"l5"`
}

type album struct {
	ID     string  `json:"id" required:"true" minLength:"1" description:"ID is a unique string that determines album."`
	Title  string  `json:"title" required:"true" minLength:"1" description:"Title of the album."`
	Artist string  `json:"artist,omitempty" description:"Album author, can be empty for multi-artist compilations."`
	Price  float64 `json:"price" minimum:"0" description:"Price in USD."`
}

//var wordlist = []Fiveletterword{{WORD: "APPLE", L1: "A", L2: "P", L3: "P", L4: "L", L5: "E"}}
//var wordlist []Fiveletterword

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// @title Web Service Gin Swagger API
// @version 1.0
// @description This is  a Web Service Gin server.
// @termsOfService http://swagger.io/terms/

// @contact.name  Web Service Gin API Support
// @contact.url http://www.swagger.io/support
// @contact.email adamcgladstone@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8010
// @BasePath /
// @schemes http
func main() {

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is web gin server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8010"
	//docs.SwaggerInfo.BasePath = "/v2"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	router := gin.Default()

	path := "./Data/word-list.csv"
	var wordlist []Fiveletterword
	//var wordlist = make([]Fiveletterword, 8)

	//wordlist := []Fiveletterword{
	//	{WORD: "APPLE", L1: "A", L2: "P", L3: "P", L4: "L", L5: "E"},
	//	{WORD: "CANOE", L1: "C", L2: "A", L3: "N", L4: "O", L5: "E"},
	//}

	wordcount, _, wordlist := readwords(path)
	//wordlist = append(wordlist, readwords(wordlist))
	fmt.Printf("%# v", pretty.Formatter(wordlist)) //It will print all struct details
	fmt.Println("wordcount is %v ", wordcount)
	//fmt.Println("%+v", wordlist)
	//if err != nil {
	//	fmt.Println("error")
	//}
	//router := gin.Default()
	// Routes
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run("0.0.0.0:8010")

	fmt.Println("Hello")
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}

// getAlbums responds with the list of all albums as JSON.
// album represents data about a record album.
// @Summary      List albums
// @Description  get albums
// @Tags         albums
// @Accept       json
// @Produce      json
// @Param        q    query     string  false  "name search by q"  Format(email)
// @Router       /albums [get]
func getAlbums(c *gin.Context) {
	c.JSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.JSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
// @Summary      Show an albums
// @Description  get string by ID
// @Tags         albums
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Album ID"
// @Router       /albums/{id} [get]
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.JSON(http.StatusOK, a)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func readwords(path string) (int, error, []Fiveletterword) {
	var line Fiveletterword
	var i int
	var words []Fiveletterword

	//path := "./Data/word-list.csv"
	file, err := os.Open(path)
	if err != nil {
		return 1, err, words
	}
	defer file.Close()

	//var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		i++
		//lines = append(lines, scanner.Text())
		wordstring := scanner.Text()
		line.WORD = wordstring
		line.L1 = wordstring[0:1]
		line.L2 = wordstring[1:2]
		line.L3 = wordstring[2:3]
		line.L4 = wordstring[3:4]
		line.L5 = wordstring[4:5]

		words = append(words, line)

	}
	//fmt.Printf("%# v", pretty.Formatter(words)) //It will print all struct details

	return i, errors.New("can't work with 42"), words

}

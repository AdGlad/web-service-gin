package main

import (
	"fmt"
	"github.com/adglad/web-service-gin/morestrings"
	"net/http"

	"github.com/gin-gonic/gin"

	"bufio"
	"errors"
	"github.com/google/go-cmp/cmp"
	"github.com/kr/pretty"
	"os"
)

type Fiveletterword struct {
	WORD string `json:"word"`
	L1   string `json:"l1"`
	L2   string `json:"l2"`
	L3   string `json:"l3"`
	L4   string `json:"l4"`
	L5   string `json:"l5"`
}

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

//var wordlist = []Fiveletterword{{WORD: "APPLE", L1: "A", L2: "P", L3: "P", L4: "L", L5: "E"}}
//var wordlist []Fiveletterword

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {

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
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumByID)

	router.Run("0.0.0.0:8010")

	fmt.Println("Hello")
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
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
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
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

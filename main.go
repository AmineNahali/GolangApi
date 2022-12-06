package main

import (
	"bufio"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func getIndx(c *gin.Context) {
	c.File("./index.html")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, album{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99})
}

func getImg(c *gin.Context) {
	c.File("./win.PNG")
}

func getText(c *gin.Context) {
	name := c.Param("name")
	lines, _ := readLines("./" + name + ".txt")
	res := ""
	for _, line := range lines {
		res += line + "|||"
	}
	c.String(200, res)

}

func main() {
	router := gin.Default()
	router.GET("/", getIndx)
	router.GET("/text/:name", getText)
	router.GET("/albums", getAlbums)
	router.GET("/image", getImg)
	router.Run("localhost:8080")
}

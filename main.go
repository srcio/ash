package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	pathBase string
	port     string
)

func init() {
	pathBase = os.Getenv("ASH_PATH_BASE")
	port = os.Getenv("ASH_PORT")
	if port == "" {
		port = "80"
	}
}

func greet(c *gin.Context) {
	c.String(http.StatusOK, fmt.Sprintf("Hello World! \nTime now is: %v\n", time.Now().Format(time.RFC3339)))
}

func headers(c *gin.Context) {
	var headers string
	for k, v := range c.Request.Header {
		headers += fmt.Sprintf("%v: %v\n", k, v)
	}
	c.String(http.StatusOK, "headers \n"+headers)
}

func main() {
	server := gin.Default()
	sg := server.Group(pathBase)
	sg.GET("", greet)
	sg.GET("/headers", headers)

	log.Fatalln(server.Run(":" + port))
}

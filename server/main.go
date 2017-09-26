package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/shawncatz/hanabi"
	"os"
)

var bloom string
var client *hanabi.Client

func main() {
	var err error

	bloom = os.Getenv("HANABI_BLOOM")
	if bloom == "" {
		bloom = "default"
	}

	client, err = hanabi.NewClient(bloom, "localhost:2717")
	if err != nil {
		log.Fatalf("error connecting to hanabi: %s", err)
	}

	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "world"})
	})

	r.POST("/hooks/:service", hookService)
	r.POST("/hooks/:service/:action", hookService)

	r.Run()
}

func hookService(c *gin.Context) {
	service := c.Param("service")
	action := c.Param("action")
	if action == "" {
		action = "default"
	}

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Printf("failed to read body: %s", err)
	}

	fmt.Printf("body:\n%s", string(body))

	client.Send(fmt.Sprintf("%s.%s", service, action), body)
}

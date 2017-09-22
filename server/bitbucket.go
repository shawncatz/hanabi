package main

import (
	"io/ioutil"
	"fmt"

	"github.com/gin-gonic/gin"
)

func bitbucket(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Printf("failed to read body: %s", err)
	}

	fmt.Printf("body:\n%s", string(body))
}

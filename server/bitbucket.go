package main

import (
	"io/ioutil"
	"fmt"

	"github.com/gin-gonic/gin"
)

func bitbucketBuild(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Printf("failed to read body: %s", err)
	}

	fmt.Printf("body:\n%s", string(body))
}

func bitbucketPr(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Printf("failed to read body: %s", err)
	}

	fmt.Printf("body:\n%s", string(body))
}

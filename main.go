package main

import (
	"github.com/gin-gonic/gin"
)

func init() {
}
func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Run(":1337")
}

package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.New()
	router.GET("/info", GetInfo)
}

func GetInfo(c *gin.Context) {
	group := c.DefaultQuery("group", "")
	song := c.DefaultQuery("song", "")

}

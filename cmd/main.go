package main

import (
	"github.com/ching-lamabam/web-tracker/internal/rest"
	"github.com/gin-gonic/gin"
)

func main() {
	//create a router and corresponding groups
	router := gin.New()

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	webAPI := rest.New()

	webTracker := router.Group("/web-tracker")
	webTrackerV1 := webTracker.Group("/v1")
	webTrackerV1.GET("/ping", webAPI.GetPing)
	webTrackerV1.GET("/img", webAPI.GetImg)

	err := router.Run(":3000")
	if err != nil {
		panic(err)
	}
}

package main

import (
	"devjudge/go-in-docker/api"
	"devjudge/go-in-docker/database"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	r.POST("/create", api.CreateUser(db))
	r.POST("/add/:userA/:userB", api.SendFriendRequest(db))
	r.GET("/friends/:userA", api.GetAllFriends(db))
	r.Run(":8080")
}

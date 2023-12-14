package api

import (
	"devjudge/go-in-docker/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserCreateRequest struct {
	Username string `json:"username"`
}

type ErrorResponse struct {
	Status string `json:"status"`
	Reason string `json:"reason"`
}

type UserCreateSuccessResponse struct {
	Username string `json:"username"`
}

type FriendRequestSuccessResponse struct {
	Status string `json:"status"`
}

type FriendsSuccessResponse struct {
	Friends []string `json:"friends"`
}

func CreateUser(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var request UserCreateRequest
		if err := c.BindJSON(&request); err != nil {
			c.IndentedJSON(400, ErrorResponse{Status: "failure", Reason: err.Error()})
			return
		}
		if err := database.CreateUserDB(db, request.Username); err != nil {
			c.IndentedJSON(400, ErrorResponse{Status: "failure", Reason: err.Error()})
			return
		}
		c.IndentedJSON(201, UserCreateSuccessResponse{Username: request.Username})
	}
}

func SendFriendRequest(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		userA := c.Param("userA")
		userB := c.Param("userB")
		if userA == "" || userB == "" {
			c.IndentedJSON(400, ErrorResponse{Status: "failure", Reason: "invalid user given"})
			return
		}
		if err := database.SendFriendRequestDB(db, userA, userB); err != nil {
			c.IndentedJSON(400, ErrorResponse{Status: "failure", Reason: err.Error()})
			return
		}
		c.IndentedJSON(202, FriendRequestSuccessResponse{Status: "success"})
	}
}

func GetAllFriends(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		user := c.Param("userA")
		if user == "" {
			c.IndentedJSON(400, ErrorResponse{Status: "failure", Reason: "invalid user given"})
			return
		}
		if friends, err := database.GetAllFriendsDB(db, user); err != nil {
			c.IndentedJSON(400, ErrorResponse{Status: "failure", Reason: err.Error()})
			return
		} else {
			if len(friends) == 0 {
				c.IndentedJSON(404, ErrorResponse{Status: "failure", Reason: "no friends found"})
				return
			}
			c.IndentedJSON(200, FriendsSuccessResponse{Friends: friends})
		}
	}
}

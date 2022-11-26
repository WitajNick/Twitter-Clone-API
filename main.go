package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type user struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Admin     bool   `json:"admin"`
}
type tweet struct {
	ID       string    `json:"ID"`
	UserID   string    `json:"userID"`
	Content  string    `json:"content"`
	Datetime time.Time `json:"datetime"`
}

var users = []user{
	{ID: "1", FirstName: "Nick", LastName: "Liszewski", Email: "nicholasliszewski@gmail.com", Admin: true},
}
var tweets = []tweet{
	{ID: "1", UserID: "1", Content: "This is the first tweet!", Datetime: time.Now()},
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func getUserByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range users {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

func postUsers(c *gin.Context) {
	var newUser user

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

func getTweets(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tweets)
}

func getTweetByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range tweets {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "tweet not found"})
}

func postTweets(c *gin.Context) {
	var newTweet tweet

	if err := c.BindJSON(&newTweet); err != nil {
		return
	}

	tweets = append(tweets, newTweet)
	c.IndentedJSON(http.StatusCreated, newTweet)
}

func main() {
	router := gin.Default()

	// Users
	router.GET("/users", getUsers)
	router.GET("/users/:id", getUserByID)
	router.POST("/users", postUsers)

	// Tweets
	router.GET("/tweets", getTweets)
	router.GET("/tweets/:id", getTweetByID)
	router.POST("/tweets", postTweets)

	router.Run("localhost:8080")
}

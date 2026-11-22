package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"github.com/google/uuid"
)

// post represents data about a blog post.
type post struct {
	ID        uuid.UUID `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	CreatedAt time.Time `json:"created_at"`
}

// posts slice to seed blog post data.
var posts = []post{
	{ID: uuid.New(), Title: "First Post", Author: "Author One", CreatedAt: time.Date(2024, 6, 1, 0, 0, 0, 0, time.UTC)},
	{ID: uuid.New(), Title: "Second Post", Author: "Author Two", CreatedAt:  time.Now()},
	{ID: uuid.New(), Title: "Third Post", Author: "Author Three", CreatedAt: time.Now()}, 
}

// getPosts responds with the list of all posts as JSON.
func getPosts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, posts)
}

func addPost(c *gin.Context) {
	var newPost post

	if err := c.BindJSON(&newPost); err != nil {
		message := "invalid request"
		if err != nil {
			message = err.Error()
		}
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": message})
		return
	}

	newPost.ID = uuid.New()
	newPost.CreatedAt = time.Now()

	posts = append(posts, newPost)
	c.IndentedJSON(http.StatusCreated, newPost)
}

func getPostByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid UUID"})
		return
	}

	for _, p := range posts {
		if p.ID == id {
			c.IndentedJSON(http.StatusOK, p)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "post not found"})
}

// main function to set up the router and routes.
func main() {
	router := gin.Default()
	router.GET("/posts", getPosts)
	router.GET("/posts/:id", getPostByID)
	router.POST("/posts", addPost)

	router.Run("localhost:8080")
}

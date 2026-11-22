package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// post represents data about a blog post.
type post struct {
    ID      string  `json:"id"`
    Title   string  `json:"title"`
    Author  string  `json:"author"`
    CreatedAt string  `json:"created_at"`
}

// posts slice to seed blog post data.
var posts = []post{
    {ID: "1", Title: "First Post", Author: "Author One", CreatedAt: "2024-06-01"},
    {ID: "2", Title: "Second Post", Author: "Author Two", CreatedAt: "2024-06-02"},
    {ID: "3", Title: "Third Post", Author: "Author Three", CreatedAt: ""},
}

// getPosts responds with the list of all posts as JSON.
func getPosts(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, posts)
}

func addPost(c *gin.Context) {
    var newPost post

    if err := c.BindJSON(&newPost); err != nil {
        return
    }

    posts = append(posts, newPost)
    c.IndentedJSON(http.StatusCreated, newPost)
}

// main function to set up the router and routes.
func main() {
    router := gin.Default()
    router.GET("/posts", getPosts)
    router.POST("/posts", addPost)

    router.Run("localhost:8080")
}
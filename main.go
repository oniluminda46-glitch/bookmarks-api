package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type bookmark struct {
	ID    int    `json:"id"`
	Title string `json:"title" binding:"required"`
	URL   string `json:"url" binding:"required,url"`
}

var bookmarks = []bookmark{}
var nextID = 1

func main() {
	router := gin.Default()

	// Read all bookmarks (/bookmarks)
	router.GET("/bookmarks", func(c *gin.Context) {
		c.JSON(200, bookmarks)
	})

	// Read one bookmark (/bookmarks/xyz)
	router.GET("/bookmarks/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		for _, b := range bookmarks {
			if b.ID == id {
				c.JSON(200, b)
				return
			}
		}

		c.JSON(404, gin.H{"error": "bookmark not found"})
	})

	// Update a bookmark
	router.PUT("/bookmarks/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		var updated bookmark

		if err := c.ShouldBind(&updated); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		for i, b := range bookmarks {
			if b.ID == id {
				updated.ID = id
				bookmarks[i] = updated
				c.JSON(200, updated)
				return
			}
		}

		c.JSON(404, gin.H{"error": "bookmark not found"})
	})

	// Delete a bookmark
	router.DELETE("/bookmarks/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		for i, b := range bookmarks {
			if b.ID == id {
				bookmarks = append(bookmarks[:i], bookmarks[i+1:]...)
				c.JSON(200, gin.H{"message": "deleted"})
				return
			}
		}

		c.JSON(404, gin.H{"error": "bookmark not found"})
	})

	// Create a bookmark
	router.POST("/bookmarks", func(c *gin.Context) {
		var newBookmark bookmark

		// JSON Binding
		if err := c.ShouldBind(&newBookmark); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		newBookmark.ID = nextID
		nextID++

		bookmarks = append(bookmarks, newBookmark)

		c.JSON(201, newBookmark)
	})

	err := router.Run(":8080")
	if err != nil {
		fmt.Println("The router encountered an error while running:", err)
	}
}

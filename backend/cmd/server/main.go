package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func convertToInt(id string) int {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return 0
	}
	return intID
}

var userDB = make(map[string]string)  // Renamed from db to userDB

func setupRouter() *gin.Engine {
    // Create a default Gin router with Logger and Recovery middleware
    r := gin.Default()

    // Basic test endpoint
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "pong",
        })
    })

    // Let's create a group for our API endpoints
    api := r.Group("/api")
    {
        // GET random joke
        api.GET("/joke/random", func(c *gin.Context) {
            joke, err := fetchRandomJoke()
            if err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
                return
            }
            c.JSON(http.StatusOK, gin.H{"joke": joke})
        })

        // GET random joke by category
        api.GET("/joke/random/:category", func(c *gin.Context) {
            category := c.Param("category")
            joke, err := fetchRandomJokeByCategory(category)
            if err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
                return
            }
            c.JSON(http.StatusOK, gin.H{"joke": joke})
        })

		// GET all categories
		api.GET("/joke/categories", func(c *gin.Context) {
			categories, err := fetchCategories()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"categories": categories})
		})

		// GET all jokes
		api.GET("/joke/all", func(c *gin.Context) {
			jokes, err := fetchAllJokes()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"jokes": jokes})
		})

		// GET jokes by category
		api.GET("/joke/category/:category", func(c *gin.Context) {
			category := c.Param("category")
			jokes, err := fetchJokesByCategory(category)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"jokes": jokes})
		})

		// GET joke by ID
		api.GET("/joke/:id", func(c *gin.Context) {
			id := c.Param("id")
			joke, err := fetchJokeByID(convertToInt(id))
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"joke": joke})
		})

		// POST add category
		api.POST("/joke/category", func(c *gin.Context) {
			var category struct {
				Name string `json:"name" binding:"required"`
			}
			if err := c.ShouldBindJSON(&category); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			err := addCategory(category.Name)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "Category added successfully"})
		})

		// POST add joke to category
		api.POST("/joke", func(c *gin.Context) {
			var joke struct {
				Category string `json:"category" binding:"required"`
				Joke     string `json:"joke" binding:"required"`
			}
			if err := c.ShouldBindJSON(&joke); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			err := addJoke(joke.Joke, joke.Category)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "Joke added successfully"})
		})

		// Like a joke
		api.POST("/joke/like/:id", func(c *gin.Context) {
			id := c.Param("id")
			err := likeJoke(convertToInt(id))
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "Joke liked successfully"})
		})

		// dislike a joke
		api.POST("/joke/dislike/:id", func(c *gin.Context) {
			id := c.Param("id")
			err := dislikeJoke(convertToInt(id))
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "Joke disliked successfully"})
		})
    }

    return r
}


func main() {
    // Load .env file first
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Initialize database second
    err = InitDB()
    if err != nil {
        log.Fatal("Database connection failed:", err)
    }
    fmt.Println("Connected to database!")
	testDB()
    // Setup and run router last
    r := setupRouter()
    r.Run(":8080")


}


package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var userDB = make(map[string]string)  // Renamed from db to userDB

func setupRouter() *gin.Engine {
    r := gin.Default()
    
    // Update all references from db to userDB in your handlers
    r.GET("/user/:name", func(c *gin.Context) {
        user := c.Params.ByName("name")
        value, ok := userDB[user]  // Changed from db to userDB
        if ok {
            c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
        } else {
            c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
        }
    })

    authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
        "foo":  "bar",
        "manu": "123",
    }))

    authorized.POST("admin", func(c *gin.Context) {
        user := c.MustGet(gin.AuthUserKey).(string)

        var json struct {
            Value string `json:"value" binding:"required"`
        }

        if c.Bind(&json) == nil {
            userDB[user] = json.Value  // Changed from db to userDB
            c.JSON(http.StatusOK, gin.H{"status": "ok"})
        }
    })

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


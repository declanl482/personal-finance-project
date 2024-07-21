package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Copied and pasted from https://go.dev/doc/tutorial/web-service-gin to test server.

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {

	// The instance for the Gin framework is Engine. It contains the muxer, middleware, and configuration settings.
	// The Default function returns an Engine instance with the Logger (logs requests and responses) and
	// Recovery (catches and handles panics, preventing server crashes) middleware already attached.

	router := gin.Default()

	if router == nil {
		fmt.Println("Error creating web server engine.")
		return
	}

	// Cross-Origin Resource Sharing (CORS) is a browser protocol that allows web applications to securely request resources
	// from different domains.

	// - AllowOrigins specifies the allowed domains that can make requests to the web server.
	// - AllowMethods specifies the allowed HTTP methods that can be used to make requests to the web server.
	// - AllowHeaders specifies the allowed HTTP headers that can be included in requests to the web server.
	// - AllowCredentials enables credentials (cookies, authorization headers, etc.) for requests.
	// - MaxAge sets the maximum age for CORS preflight requests, which are requests that are made before the actual request to determine
	//   if the request is allowed.

	corsConfig := cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "*"}, // Allow the frontend application to make requests to the backend.
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	// The Use function attaches a global middleware to the router. The middleware attached to the router through Use will be included in
	// the handlers chain for every single request.

	router.Use(cors.New(corsConfig))

	// Copied and pasted from https://go.dev/doc/tutorial/web-service-gin to test server.

	router.GET("/albums", getAlbums)

	router.Run("localhost:8080")
}

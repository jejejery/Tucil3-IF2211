package main

import (
  "net/http"
  
  "github.com/gin-gonic/contrib/static"
  "github.com/gin-gonic/gin"
)

func main() {
  // Set the router as the default one shipped with Gin
  router := gin.Default()
  
  // Serve frontend static files
  router.Use(static.Serve("/", static.LocalFile("./views", true)))
  
  // Setup route group for the API
  api := router.Group("/api")
  {
    api.GET("/", func(c *gin.Context) {
      c.JSON(http.StatusOK, gin.H {
        "message": "pong",
      })
    })
  }
  

  
  // Start and run the server
  router.Run(":8080")
}

// Routing Planning
func RoutingPlanning(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H {
	  "message":"Routing not implemented yet",
	})
  }
  
  //Choosing the right algorithm (UCS or A*)
  func ChooseAlgorithm(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H {
	  "message":"ChooseAlgorithm handler not implemented yet",
	})
  }
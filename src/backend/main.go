package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
  algorithm "backend/lib/algorithm"
  reader "backend/lib/reader"
  structs "backend/lib/structs"
  "fmt"
)

var datainfo map[int] structs.MapValue = make(map[int]structs.MapValue)
var graph structs.Graph;
var algo structs.NodeInfo;

func main() {
	
	r := gin.Default()

	// Enable CORS
	r.Use(corsMiddleware())

	// Route for file uploads
    r.POST("/upload", handleFileUpload)

	// Algorithm Calculation
	//request start and end
	r.POST("/algo",handleReq)

	//calculate to throw route
	r.GET("/algo",handleRouting)

    // Start the server
    r.Run(":8000")

}
func handleReq(c *gin.Context){

	var data struct{
		Start string `json:"start"`
		Dest string  `json:"dest"`
		IsUCS bool 	 `json:"isUCS"`
	}
	  if err := c.BindJSON(&data); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	  }

	start := 0
	dest := 0
	for key := range datainfo {
		if(structs.GetName(datainfo[key]) ==data.Start){
			start = key
		}

		if(structs.GetName(datainfo[key]) ==data.Dest){
			dest = key
		}
	}
	if data.IsUCS {
		algo = algorithm.UCS(datainfo,graph,start,dest)
	}

	if !data.IsUCS{
		algo = algorithm.AStar(datainfo,graph,start,dest)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data received successfully"})
}

func handleRouting(c *gin.Context){
	var buff struct{
		Id       int   	`json:"id"`    
		Path     []int   `json:"path"`
		PathCost float64   `json:"pathCost"`
	}

	buff.Id = structs.GetId(algo)
	buff.Path = structs.GetPath(algo)
	buff.PathCost = structs.GetPathCost(algo)

	c.JSON(http.StatusOK, buff)
}

func handleFileUpload(c *gin.Context) {
    // Get the uploaded file
    file, err := c.FormFile("file")
    if err != nil {
        c.String(http.StatusBadRequest, fmt.Sprintf("error uploading file: %s", err.Error()))
        return
    }

    // Save the file to disk
    if err := c.SaveUploadedFile(file, file.Filename); err != nil {
        c.String(http.StatusInternalServerError, fmt.Sprintf("error saving file: %s", err.Error()))
        return
    }
	
	// Open the file
	f, err := file.Open()
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error opening file: %s", err.Error()))
		return
	}
	defer f.Close()

    // Return a success message
    reader.ReadInput(&datainfo, &graph, f,c)
	fmt.Println("UCS:", algorithm.UCS(datainfo, graph, 0, 5))
	fmt.Println("Astar:", algorithm.AStar(datainfo, graph, 0, 5))
}


// CORS middleware
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}


package main 

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {

	r:= gin.Default()
	r.GET("/ping", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"Message": "Pong",
			"New Message": "Server running on port 8000",
		})
	})
	r.Run(":8000")
 
	

}

// func RunServer(){

// 	r:=gin.Default() // r means router, we can use any r or else router, but r is commonly used for router
// 	r.GET("/ping", func(c *gin.Context){
// 		c.JSON(http.StatusOK, gin.H{     // gin.H is a shortcut for map[string]interface{}
// 			"Message": "Pong",
// 			"New Message": "Hello World",
// 		})
// 	})
// 	r.Run(":8000") // listen and serve on default server 8080 if you didn't specify any port
//}
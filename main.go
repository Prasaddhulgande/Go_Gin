package main 

import (
	"net/http"
	"github.com/gin-gonic/gin"
)
var r= gin.Default()
func main() {


	POSTAPI(r)
	GETIdNewId(r)
	GETAPI(r)
	PUTAPI(r)
	UpdateAPI(r)

	r.Run(":8000")
 
	

}
func GETAPI(router *gin.Engine) {
	r.GET("/ping", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"Message": "Pong running on port :8000",
			"Message 2": "This is GET API",
		})
	})

}
func GETIdNewId(router *gin.Engine) {
	r.GET("/me/:id/:newId", func(c *gin.Context){
		id:= c.Param("id")
		newId:= c.Param("newId")

		c.JSON(http.StatusOK, gin.H{
			"User ID": id,
			"New User ID":newId,
		})
	})
	
}
func POSTAPI(router *gin.Engine) {
	//r:= gin.Default() // r declare globally 
	r.POST("/me", func(c *gin.Context){
		// Add Authontications with email and pass
		type MeRequest struct {
			Email string `json:"email" binding: "required"`
			Password string `json: "password"`
		}
		//If you binding email as a required then need to handle error
		var meRequest MeRequest
		if err:= c.BindJSON(&meRequest); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"email": meRequest.Email,
			"password": meRequest.Password,
		})
	})
	//r.Run(":8000") // router provided as a parameter into this func
}
// PUT API

func PUTAPI(router *gin.Engine) {
	r.PUT("/me", func(c *gin.Context){

		type MeRequest struct {
			Email string `json: "email" binding:"request"`
			Password string `json: "password"`
		}
		var meRequest MeRequest

		if err:= c.BindJSON(&meRequest); err!= nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return 
		}
		c.JSON(http.StatusOK, gin.H{
			"email": meRequest.Email,
			"password": meRequest.Password,
		})
	})
}

func UpdateAPI(router *gin.Engine){
	r.PATCH("/me", func(c *gin.Context){
      type MeRequest struct{
		Email string `json: "email" binding:"required"`
		Password string `json: "password"`
	  }
	  var meRequest MeRequest

	  if err:= c.BindJSON(&meRequest); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return 
	  }
	  c.JSON(http.StatusOK, gin.H{
		"email": meRequest.Email,
		"password": meRequest.Password,
	  })
	})
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
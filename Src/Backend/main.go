package main

import (	
	"net/http"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/", func(c *gin.Context){
		c.String(http.StatusOK,"Merchandise Manager Serving . . .")
	})

	//01
	router.GET("/requestItemList", func(c *gin.Context){
		c.String(http.StatusOK,"Request Item List")
	})

	//02
	router.POST("/addItem", func(c *gin.Context){
		c.String(http.StatusOK,"Add Item")
	})

	//03
	router.POST("/deleteItem", func(c *gin.Context){
		c.String(http.StatusOK,"Delete Item")
	})

	//04
	router.POST("/registerOrder", func(c *gin.Context){
		c.String(http.StatusOK,"Register Order")
	})

	//05
	router.POST("/makeOrder", func(c *gin.Context){
		c.String(http.StatusOK,"Make Order")
	})

	//06
	router.GET("/checkOrder", func(c *gin.Context){
		c.String(http.StatusOK,"Check Order")
	})

	//07
	router.POST("/login", func(c *gin.Context){
		c.String(http.StatusOK,"LogIn")
	})

	//08
	router.POST("/requestForm", func(c *gin.Context){
		c.String(http.StatusOK,"Submit Request")
	})

	//09
	router.GET("/requestList", func(c *gin.Context){
		c.String(http.StatusOK,"Request List")
	})

	//10
	router.DELETE("/deleteRequest", func(c *gin.Context){
		c.String(http.StatusOK,"Delete Request")
	})

	router.Run(":13131")
}

func check(e error){
        if e != nil {
           panic(e)
        }
}

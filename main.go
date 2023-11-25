package main

import (
	"github.com/abhinavthapa1998/go-jwt/controllers"
	"github.com/abhinavthapa1998/go-jwt/initializers"
	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}


func main() {
	r := gin.Default()
	r.POST("/signup",controllers.Signup)
	r.POST("/login", controllers.Login)
	r.Run()
}

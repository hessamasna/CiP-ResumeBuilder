package main

import (
	"net/http"
	"userService/controller"
	"userService/grpc_init"
	"userService/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
	go grpc_init.StartGrpcServer()
}


func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin, Accept")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(http.StatusNoContent)
            return
        }

        // if c.Request.Header.Get("Content-Type") != "application/json" {
        //     c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request content type. Expected JSON."})
        //     return
        // }

        c.Next()
    }
}

func main() {
	r := gin.Default()

	    // Add CORS middleware
	r.Use(CORSMiddleware())

	r.POST("auth/signup", controller.Signup)
	r.POST("auth/login", controller.Signin)
	r.POST("auth/refresh", controller.RefreshAccessToken)
	r.POST("/auth/logout", controller.Logout)

	r.POST("/cv/create" , controller.CreateCv)
	r.GET("cv/getAll/:user_id" , controller.GetCvsByUserId)
	r.GET("cv/get/:id" , controller.GetCvById)
	r.Run()

}

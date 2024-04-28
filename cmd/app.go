package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/JerryJeager/will-be-there-backend/api"
	"github.com/JerryJeager/will-be-there-backend/manualwire"
	"github.com/JerryJeager/will-be-there-backend/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var userController = manualwire.GetUserController()
var eventController = manualwire.GetEventController()
var inviteeController = manualwire.GetInviteeController()

func ExecuteApiRoutes() {
	fmt.Println("executing api routes")

	r := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
    // r.MaxMultipartMemory = 8 << 20 // 8 MiB

	r.Use(cors.Default())
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello!",
		})
	})

	v1 := r.Group("/api/v1")

	v1.GET("/info/openapi.yaml", func(c *gin.Context) {
		c.String(200, api.OpenApiDocs())
	})

	users := v1.Group("/users")
	users.POST("/signup", userController.CreateUser)
	users.POST("/login", userController.CreateToken)

	user := users
	user.Use(middleware.JwtAuthMiddleware())
	user.GET("/:user-id", userController.GetUser)

	event := v1.Group("/event")
	event.Use(middleware.JwtAuthMiddleware())
	{
		event.POST("", eventController.CreateEvent)
		event.GET("/:event-id", eventController.GetEvent)
		event.GET("user/:user-id", eventController.GetEvents)
		event.PUT("/:event-id/image", middleware.FileUploadMiddleware(), eventController.UpdateImageurl)
	}

	invitation := v1.Group("/invitation")
	// invitation.Use(middleware.JwtAuthMiddleware())
	// {
		invitation.POST("/guest", inviteeController.CreateInvitee)
		invitation.PATCH("/guest/:invitee-id", inviteeController.UpdateInviteeStatus)
		invitation.PUT("/guest/:invitee-id", inviteeController.UpdateInvitee)
		invitation.GET("/guests/:event-id", inviteeController.GetInvitees)
	// }
	invitation.Use(middleware.JwtAuthMiddleware())
	{
		invitation.DELETE("/guest/:invitee-id", inviteeController.DeleteInvitee)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := r.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}

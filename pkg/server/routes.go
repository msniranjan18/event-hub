package server

import (
	"github.com/gin-gonic/gin"
	"msn.com/event-hub/pkg/handlers"
	"msn.com/event-hub/pkg/middlewares"
)

func RegisterRoute(server *gin.Engine) {
	//events routes
	server.GET("/events", handlers.GetEvents)
	server.GET("/events/:id", handlers.GetEventbyId)

	authenticated := server.Group("")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", handlers.CreateEvent)
	authenticated.PUT("/events/:id", handlers.UpdateEvent)
	authenticated.DELETE("/events/:id", handlers.DeleteEvent)
	//event registration and cancellation
	authenticated.POST("/events/:id/register", handlers.EventRegistration)
	authenticated.POST("/events/:id/cancel", handlers.EventCancellation)
	authenticated.GET("/users/event_registrations", handlers.GetUsersEventRegistrations)
	authenticated.GET("/registrations", handlers.GetAllRegistrations)
	//users routes
	server.POST("/users/signup", handlers.CreateUser)
	server.POST("/users/login", handlers.UserLogin)
	server.PUT("users/:email", handlers.UpdateUser)
	server.GET("/users", handlers.GetUsers)
	server.GET("/users/:email", handlers.GetUserByEmailID)
	server.DELETE("/users/:email", handlers.DeleteUser)

}

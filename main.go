package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"msn.com/event-hub/pkg/db"
	"msn.com/event-hub/pkg/server"
)

var (
	buildVersion string
	appEnv       string
)

func main() {
	fmt.Println("Start...")
	if appEnv == "" {
		appEnv = os.Getenv("APP_ENV")
	}
	if buildVersion == "" {
		buildVersion = os.Getenv("BUILD_VERSION")
	}

	fmt.Printf("App Environment: %s, Build Version: %s\n", appEnv, buildVersion)
	// logging setup
	timeStampStr := fmt.Sprintf("%d", time.Now().Unix())
	logFilePath := "/tmp/event-hub_" + timeStampStr + ".log"
	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("error opening log file: %v", err)
	}
	defer logFile.Close()

	// Create a logger that writes to the MultiWriter
	multiWriter := io.MultiWriter(os.Stdout, logFile)

	logger := log.New(multiWriter, "", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Printf("starting event-hub... Environment: %s, Version: %s\n", appEnv, buildVersion)

	// database setup
	db.InitDB()

	// gin engine setup
	ginServer := gin.Default()
	//ginServer.LoadHTMLGlob()

	// CORS configuration
	ginServer.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Allows all origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Middleware to add logger to the context
	ginServer.Use(func(c *gin.Context) {
		c.Set("logger", logger)
		c.Next()
	})

	// Serve static files from the "public" directory
	ginServer.Static("/static", "./front-end")

	// // Serve the index page at the root URL
	// ginServer.GET("/", func(c *gin.Context) {
	// 	c.File("./front-end/index.html")
	// })

	// Serve the signup page at the /signup URL
	ginServer.GET("/signup", func(c *gin.Context) {
		c.File("./front-end/user-signup.html")
	})

	ginServer.GET("/login", func(c *gin.Context) {
		c.File("./front-end/user-login.html")
	})

	ginServer.GET("/event-list", func(c *gin.Context) {
		c.File("./front-end/events.html")
	})

	ginServer.GET("/registered-event-list", func(c *gin.Context) {
		c.File("./front-end/user-registered-events.html")
	})

	ginServer.GET("/user-dashboard", func(c *gin.Context) {
		c.File("./front-end/user-dashboard.html")
	})

	ginServer.GET("/event-create", func(c *gin.Context) {
		c.File("./front-end/event-create.html")
	})

	server.RegisterRoute(ginServer)

	ginServer.Run(":8080") //localhost:8080
}

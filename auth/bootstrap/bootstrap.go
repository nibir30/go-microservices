package bootstrap

import (
	"fmt"

	"github.com/nibir30/go-microservices/auth/config"
	"github.com/nibir30/go-microservices/auth/db"
	"github.com/nibir30/go-microservices/auth/internal/api/routes"

	"os/exec"

	"github.com/gin-gonic/gin"
	docs "github.com/nibir30/go-microservices/auth/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

// App struct holds dependencies
type App struct {
	Router    *gin.Engine
	DB        *gorm.DB
	Container *Container
}

// InitializeApp sets up DB, repositories, services, and routes
func InitializeApp() *App {
	// Initialize database connection
	dbInstance := config.ConnectDB()

	// Run database migrations
	db.MigrateDB(dbInstance)

	// Initialize container (repositories & services)
	container := NewContainer(dbInstance)

	// Set up Gin router
	router := gin.Default()

	InitializeSwagger(router)

	// Register routes with container services
	routes.RegisterUserRoutes(router, container.UserService)

	// If you have more routes, register them here:
	// routes.RegisterOrderRoutes(router, container.OrderService)

	return &App{
		Router:    router,
		DB:        dbInstance,
		Container: container,
	}
}

func InitializeSwagger(router *gin.Engine) {
	cmd := exec.Command("swag", "init")
	cmd.Stdout = nil
	cmd.Stderr = nil
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error running swag init:", err)
	}

	docs.SwaggerInfo.BasePath = "/api/v1"

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

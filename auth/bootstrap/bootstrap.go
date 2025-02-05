package bootstrap

import (
	"github.com/nibir30/go-microservices/auth/config"
	"github.com/nibir30/go-microservices/auth/db"
	"github.com/nibir30/go-microservices/auth/internal/api/routes"

	"github.com/gin-gonic/gin"
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

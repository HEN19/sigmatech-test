package routes

import (
	"github.com/api-skeleton/config"
	"github.com/api-skeleton/endpoint"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Controller() *gin.Engine {
	// routes := mux.NewRouter()

	// user := routes.PathPrefix("/user").Subrouter()
	// user.HandleFunc("/register", endpoint.RegistrationEndpoint).Methods("POST", "OPTIONS")
	// user.HandleFunc("/login", endpoint.LoginEndpoint).Methods("POST", "OPTIONS")
	routes := gin.Default()

	// CORS configuration
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:5173"}                   // Allow requests from your frontend domain
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"} // Allow methods
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"} // Allow headers

	// Use the CORS middleware with the specified config
	routes.Use(cors.New(corsConfig))
	routes.Use(removeTrailingSlash())

	api := routes.Group("/api")
	{

		// Users Endpoint
		user := api.Group("/user")
		{
			user.POST("/register", endpoint.RegistrationEndpoint)
			user.POST("/login", endpoint.LoginEndpoint)
			user.GET("/profile", endpoint.UserWithParamEndpoint)
		}

		//Customers Endpoint
		customer := api.Group("/customer")
		{
			customer.POST("/", config.AuthMiddleware(), endpoint.CustomerWithoutParamEndpoint)
			customer.GET("/", config.AuthMiddleware(), endpoint.CustomerWithoutParamEndpoint)
			customer.PUT("/:id", config.AuthMiddleware(), endpoint.CustomerWithParamEndpoint)
			customer.GET("/:id", config.AuthMiddleware(), endpoint.CustomerWithParamEndpoint)
			customer.DELETE("/:id", config.AuthMiddleware(), endpoint.CustomerWithParamEndpoint)
		}

		// Transaction Endpoint
		transaction := api.Group("/transaction")
		{
			transaction.GET("/", config.AuthMiddleware(), endpoint.TransactionWithoutParamEndpoint)
			transaction.POST("/", config.AuthMiddleware(), endpoint.TransactionWithoutParamEndpoint)
			transaction.GET("/:id", config.AuthMiddleware(), endpoint.TransactionWithParamEndpoint)
			transaction.PUT("/:id", config.AuthMiddleware(), endpoint.TransactionWithParamEndpoint)
			transaction.DELETE("/:id", config.AuthMiddleware(), endpoint.TransactionWithParamEndpoint)

		}
	}

	// user := routes.PathPrefix("/user").Subrouter()
	// user.HandleFunc("/register", endpoint.RegistrationEndpoint).Methods("POST", "OPTIONS")
	// user.HandleFunc("/login", endpoint.LoginEndpoint).Methods("POST", "OPTIONS")

	return routes
}

func removeTrailingSlash() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		if len(path) > 1 && path[len(path)-1] == '/' {
			c.Request.URL.Path = path[:len(path)-1]
		}
		c.Next()
	}
}

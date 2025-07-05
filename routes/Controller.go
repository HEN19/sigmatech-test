package routes

import (
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
			customer.POST("/", endpoint.CustomerWithoutParamEndpoint)
			customer.GET("/", endpoint.CustomerWithoutParamEndpoint)
			customer.PUT("/:id", endpoint.CustomerWithParamEndpoint)
			customer.GET("/:id", endpoint.CustomerWithParamEndpoint)
			customer.DELETE("/:id", endpoint.CustomerWithParamEndpoint)
		}

		// Transaction Endpoint
		transaction := api.Group("/transaction")
		{
			transaction.GET("/", endpoint.TransactionWithoutParamEndpoint)
			transaction.POST("/", endpoint.TransactionWithoutParamEndpoint)
			transaction.GET("/:id", endpoint.TransactionWithParamEndpoint)
			transaction.PUT("/:id", endpoint.TransactionWithParamEndpoint)
			transaction.DELETE("/:id", endpoint.TransactionWithParamEndpoint)

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

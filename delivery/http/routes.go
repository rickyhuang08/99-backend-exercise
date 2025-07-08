package http

import (
	"github.com/gin-gonic/gin"
	"github.com/rickyhuang08/99-backend-exercise/middleware"
)

// RegisterRoutes sets up API routes
func RegisterRoutes(r *gin.Engine, handler *Handler, mw *middleware.MiddlewareModule) {
	// Public routes
	public := r.Group("/public-api")
	{
		// Register global middleware
		mw.RegisterGlobalMiddleware(public)
		public.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "ok"})
		})
		// Add more public routes here
		// POST
		public.POST("/login", handler.LoginHandler)
		public.POST("/users", handler.CreateUserHandler)
		public.POST("/listings", handler.CreateListingHandler)

		// GET
		public.GET("/users", handler.GetUserHandler)
		public.GET("/listings", handler.GetListingsForPublicHandler)

	}

	// Protected routes
	private := r.Group("/private-api")
	{
		// Register global middleware
		mw.RegisterGlobalMiddleware(private)
		// Register auth middleware for protected routes
		err := mw.RegisterAuthMiddleware(private)
		if err != nil {
			panic("Failed to register auth middleware: " + err.Error())
		}
		// GET
		private.GET("/users", handler.GetUserHandler) // Example protected route
		private.GET("/listings", handler.GetListingHandler) // Example protected route

		// POST
		private.POST("/users", handler.CreateUserHandler) // Example protected route
		private.POST("/listings", handler.CreateListingHandler) // Example protected route
	}
}

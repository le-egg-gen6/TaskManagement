package route

import (
	"github.com/gin-gonic/gin"
	"go-ecommerce/api/middleware"
	"go-ecommerce/bootstrap"
	"go-ecommerce/mongo"
	"time"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine) {
	publicRouter := gin.Group("/api/v1")
	// All Public APIs
	NewSignupRouter(env, timeout, db, publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)
	NewRefreshTokenRouter(env, timeout, db, publicRouter)

	protectedRouter := gin.Group("/api/v1")
	// Middleware to verify AccessToken
	protectedRouter.Use(middleware.JWTAuthMiddleware(env.AccessTokenSecret))
	// All Private APIs
	NewProfileRouter(env, timeout, db, protectedRouter)
	NewTaskRouter(env, timeout, db, protectedRouter)
}

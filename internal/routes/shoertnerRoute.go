package routes

import (
	"shorten-url/internal/handler"
	"shorten-url/internal/middleware"
	"time"

	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SetUpRoutesForURLShortner() {
	router := gin.Default()

	// Rate limiter
	store := ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
		Rate:  10 * time.Second,
		Limit: 5,
	})
	mw := ratelimit.RateLimiter(store, &ratelimit.Options{
		ErrorHandler: errorHandler,
		KeyFunc:      keyFunc,
	})

	// Adding Custom logger in context
	logger, _ := zap.NewProduction()
	router.Use(middleware.LoggerMiddleware(logger))

	router.POST("/short", mw, handler.ShortURL)
	router.GET("/:shortKey", mw, handler.RedirectURL)

	router.Run()

}

func keyFunc(c *gin.Context) string {
	return c.ClientIP()
}

func errorHandler(c *gin.Context, info ratelimit.Info) {
	c.String(429, "Too many requests. Try again in "+time.Until(info.ResetTime).String())
}

package middlewares

import (
	"github.com/didip/tollbooth/v8"
	"github.com/didip/tollbooth/v8/limiter"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LimitByRequest() gin.HandlerFunc {
	lmt := tollbooth.NewLimiter(1, nil)
	lmt.SetIPLookup(limiter.IPLookup{
		Name:           "X-Forwarded-For",
		IndexFromRight: 0,
	})
	return func(c *gin.Context) {
		err := tollbooth.LimitByRequest(lmt, c.Writer, c.Request)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.Next()
	}
}

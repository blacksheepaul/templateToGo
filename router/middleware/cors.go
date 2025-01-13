package middleware

import (
	"github.com/blacksheepaul/templateToGo/core/config"

	"github.com/gin-gonic/gin"
)

func Cors(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
			for _, v := range cfg.Server.AllowOrigins {
				if v == origin {
					c.Header("Access-Control-Allow-Origin", v)
					break
				}
			}
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(204)
		}
		c.Next()
	}
}

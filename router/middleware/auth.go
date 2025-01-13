package middleware

import (
	"errors"
	"strings"
	"sync"

	"github.com/blacksheepaul/templateToGo/model"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session, err := GetSessionFromHeader(c)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"msg": err.Error()})
			return
		}

		if !isValidUserToken(session) {
			c.AbortWithStatusJSON(401, gin.H{
				"msg": "Invalid or expired token",
			})
			return
		}

		c.Next()
	}
}

var dao *model.Dao
var once sync.Once

func isValidUserToken(token string) bool {
	once.Do(func() {
		dao = model.GetDao()
	})
	if _, ok := dao.GetCache(token); ok {
		return true
	}
	return false
}

var (
	ErrNoSession      = errors.New("no session found")
	ErrInvalidSession = errors.New("invalid session")
)

func GetSessionFromHeader(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return "", ErrNoSession
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", ErrInvalidSession
	}

	return parts[1], nil
}

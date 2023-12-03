package app

import (
	"bmstu-web-backend/internal/app/ds"
	"bmstu-web-backend/internal/app/role"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
	"strings"
)

const jwtPrefix = "Bearer "

func (app *Application) WithAuthCheck(assignedRoles ...role.Role) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		jwtStr := c.GetHeader("Authorization")
		if !strings.HasPrefix(jwtStr, jwtPrefix) {
			for _, oneOfAssignedRole := range assignedRoles {
				if role.NotAuthorized == oneOfAssignedRole {
					return
				}
			}
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		jwtStr = jwtStr[len(jwtPrefix):]

		err := app.redisClient.CheckJWTInBlacklist(c.Request.Context(), jwtStr)
		if err == nil {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		if !errors.Is(err, redis.Nil) { // значит что это не ошибка отсуствия - внутренняя ошибка
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}


		claims := &ds.JWTClaims{}
		token, err := jwt.ParseWithClaims(jwtStr, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(app.config.JWT.Token), nil
		})
		if err != nil || !token.Valid {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		for _, oneOfAssignedRole := range assignedRoles {
			if claims.Role == oneOfAssignedRole {
				c.Set("userId", claims.UserUUID)
				c.Set("userRole", claims.Role)
				return
			}
		}
		c.AbortWithStatus(http.StatusForbidden)
		log.Println("role is not assigned")
	}

}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		for _, err := range c.Errors {
			log.Println(err.Err)
		}
		lastError := c.Errors.Last()
		if lastError != nil {
			switch c.Writer.Status() {
			case http.StatusBadRequest:
				c.JSON(-1, gin.H{"error": "wrong request"})
			case http.StatusNotFound:
				c.JSON(-1, gin.H{"error": lastError.Error()})
			case http.StatusMethodNotAllowed:
				c.JSON(-1, gin.H{"error": lastError.Error()})
			default:
				c.Status(-1)
			}
		}
	}
}

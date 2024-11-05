package middleware

import (
	"net/http"

	"github.com/Takina-Space/backend-go/app/service"
	"github.com/Takina-Space/backend-go/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Middleware for Auth JWT
func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		AuthService := service.NewAuthService()
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "No Authorization header found"})
			return
		}
		const BearerSchema = "Bearer "
		tokenString := authHeader[len(BearerSchema):]
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "No Authorization header found"})
			return
		}
		valid, token, err := AuthService.VerifyJWTRSA(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": err.Error()})
			return
		}
		if valid {
			claims := token.Claims.(jwt.MapClaims)

			if claims["iss"] != config.GetEnv("JWT_ISS") {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Issuer not valid"})
				return
			}
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": err.Error()})
			return
		}

	}
}

package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"task-manager/data"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context){
		authHeader := c.GetHeader("Authorization")	
		if authHeader == ""{
			c.JSON(http.StatusUnauthorized,gin.H{"error":"authorization header missing"})
			c.Abort()
			return

		}
		tokenString := strings.TrimPrefix(authHeader,"Bearer ")

		if tokenString == authHeader{
			c.JSON(http.StatusUnauthorized,gin.H{"error":"invalid token format"})
			c.Abort()
			return
		}

		token,err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error){
			if _,ok := token.Method.(*jwt.SigningMethodHMAC); !ok{
				return nil,jwt.ErrSignatureInvalid
			}
			return data.JwtSecret,nil
		})

		if err != nil || !token.Valid{
			c.JSON(http.StatusUnauthorized,gin.H{"error":"invalid token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if ok {
			c.Set("claims", claims)
		}else{
			c.JSON(401, gin.H{"error": "invalid auth claims"})
			c.Abort()
			return 
		}
		c.Set("claims", claims)
		c.Next()
	}
}

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, exists := c.Get("claims")
		if !exists || claims.(jwt.MapClaims)["role"].(string) != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
			c.Abort()
			return
		}
		c.Next()
	}
}
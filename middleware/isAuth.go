package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ankit-k56/Repelit/initializers"
	"github.com/ankit-k56/Repelit/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func CheckAuth(c *gin.Context){
	tokenString , err := c.Cookie("Authorization")
	if err != nil{

		c.AbortWithStatus(http.StatusUnauthorized)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
	
	
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		log.Fatal(err)
	}
	
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) >claims["exp"].(float64){
			c.AbortWithStatus(http.StatusUnauthorized)
			
		}
		var user models.User
		initializers.Db.First(&user, claims["sub"])

		if(user.ID ==0){
			c.JSON(http.StatusBadRequest, gin.H{"message": "User not found"})
			return

		}
		c.Set("user", user)
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
		
	}
}
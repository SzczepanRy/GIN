package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"jwt.com/api/initializers"
	"jwt.com/api/models"
)

func ReqAuthc(c *gin.Context) {

	fmt.Println("middleware")

	//get cookie
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		fmt.Println("EEEEEEEEEEEEEEEEEEEEEEEEE")

		c.AbortWithStatus(http.StatusUnauthorized)
	}
	//decode and valiadte

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		fmt.Println("DDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDd")

		c.AbortWithStatus(http.StatusUnauthorized)

	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {

		//check exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			//expired
			fmt.Println("CCCCCCCCCCCCCCCCCCCCCCCCCCC")

			c.AbortWithStatus(http.StatusUnauthorized)
		}

		//find the user with token

		var user models.User

		initializers.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			fmt.Println("AAAAAAAAAAAAAAA")

			c.AbortWithStatus(http.StatusUnauthorized)
		}

		//attach to req

		c.Set("user", user)
		//continue
		c.Next()

	} else {
		fmt.Println("BBBBBBBBBBBBBBBBBBBBBBBBBb")

		c.AbortWithStatus(http.StatusUnauthorized)

	}

}

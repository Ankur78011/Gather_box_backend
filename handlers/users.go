package handlers

import (
	"fmt"
	"net/http"
	"time"

	"gatherbox.com/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func generateJwt(user_id int) (string, error) {
	var key = []byte("ankur")
	expirationtime := time.Now().Add(time.Hour * 48).Unix()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = expirationtime
	tokenString, _ := token.SignedString(key)
	return tokenString, nil

}

func (api *ApiHandler) CreateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ctx.Header("Content-Type", "application/json")
		var NewUser models.User
		err := ctx.ShouldBindJSON(&NewUser)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Credential missing"})
			return
		}

		api.Storage.PostUser(NewUser.Name, NewUser.Email, NewUser.UserName, NewUser.Password)
		// token generation

		user_id := api.Storage.GetUserId(NewUser.Email)
		fmt.Println(user_id, "jai shree ram")
		token, _ := generateJwt(user_id)
		ctx.JSON(http.StatusOK, gin.H{"token": token})

	}
}

func (api *ApiHandler) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json")
		var loginCredentials models.Login
		err := ctx.ShouldBindJSON(&loginCredentials)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"mess": "credientail missing"})
			return
		}
		mess, result, user_type := api.Storage.ConfirmUser(loginCredentials.Email, loginCredentials.Password)

		if result && user_type == "Customer" {
			user_id := api.Storage.GetUserId(loginCredentials.Email)
			token, _ := generateJwt(user_id)
			ctx.JSON(http.StatusOK, gin.H{"mess": token, "userType": user_type})
			return
		} else if result && user_type == "Owner" {
			owner_id := api.Storage.GetOwnerId(loginCredentials.Email)
			token, _ := generateJwt(owner_id)
			ctx.JSON(http.StatusOK, gin.H{"mess": token, "userType": user_type, "ownerId": owner_id})
			return
		}
		ctx.JSON(http.StatusUnauthorized, gin.H{"mess": mess})

	}
}

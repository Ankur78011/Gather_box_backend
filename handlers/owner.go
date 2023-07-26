package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"gatherbox.com/models"
	"github.com/gin-gonic/gin"
)

func (api *ApiHandler) AdminLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json")
		var loginCredentials models.Login
		err := ctx.ShouldBindJSON(&loginCredentials)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "information missong"})
			fmt.Println(err)
			return
		}
		mess, result := api.Storage.ConfirmOwner(loginCredentials.Email, loginCredentials.Password)
		if result {
			owner_id := api.Storage.GetOwnerId(loginCredentials.Email)
			token, _ := generateJwt(owner_id)
			ctx.JSON(http.StatusOK, gin.H{"Token": token})
			return
		}

		ctx.JSON(http.StatusUnauthorized, gin.H{"message": mess})
	}
}

func (api *ApiHandler) PostAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json")
		var OwnerDetails models.PostOwner
		_ = json.NewDecoder(ctx.Request.Body).Decode(&OwnerDetails)
		api.Storage.CreateOwner(OwnerDetails.Name, OwnerDetails.Email, OwnerDetails.UserName, OwnerDetails.Password)
		ctx.JSON(http.StatusOK, gin.H{"meaasge": "inserted"})
	}
}

func (api *ApiHandler) UpdateOwner() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ctx.Header("Content-Type", "application/json")
		owner_id, _ := strconv.Atoi(ctx.Param("id"))
		var UpdateValues models.PostOwner
		_ = json.NewDecoder(ctx.Request.Body).Decode(&UpdateValues)
		fmt.Println(UpdateValues.Name, "kkk")
		api.Storage.UpdateOwner(UpdateValues.Name, UpdateValues.Email, UpdateValues.UserName, UpdateValues.Password, owner_id)
	}
}

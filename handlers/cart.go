package handlers

import (
	"encoding/json"
	"strconv"

	"gatherbox.com/models"
	"github.com/gin-gonic/gin"
)

func (api *ApiHandler) PostCart() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json")
		order_id, err := strconv.Atoi(ctx.Param("order_id"))
		if err != nil {
			panic(err)
		}
		var allCart []models.CartItem
		err = ctx.ShouldBindJSON(&allCart)
		if err != nil {
			panic(err)
		}
		for i := 0; i < len(allCart); i++ {
			var cartji models.CartItem
			cartji = allCart[i]
			json.NewDecoder(ctx.Request.Body).Decode(&cartji)
			// ctx.ShouldBindJSON(&cartji)
			api.Storage.CartPost(cartji.Name, cartji.Date, cartji.Number, cartji.Price, order_id)
			// fmt.Println(cartji.Name)
		}

	}
}

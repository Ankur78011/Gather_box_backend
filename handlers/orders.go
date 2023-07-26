package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"gatherbox.com/models"
	"github.com/gin-gonic/gin"
)

func (api *ApiHandler) PostOrderDetails() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		user_id := int(ctx.GetFloat64("user_id"))

		ctx.Header("Content-Type", "application/json")
		var OrderDetails models.OrderDetails
		err := ctx.ShouldBindJSON(&OrderDetails)
		if err != nil {
			fmt.Println("jjj")
			fmt.Println(err)
			ctx.AbortWithStatus(http.StatusBadRequest)

			return
		}

		order_id := api.Storage.InsertOrderDetails(OrderDetails.Name, OrderDetails.MobileNumber, OrderDetails.Address, OrderDetails.ZipCode, OrderDetails.Email, user_id)

		ctx.JSON(http.StatusOK, gin.H{"order_id": order_id})
	}
}

func (api *ApiHandler) GetOrders() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-type", "application/json")
		// owner_id := int(ctx.GetFloat64("owner_id"))
		res_id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(err)
		}
		// orderList := api.Storage.GetOrder(res_id)
		orderCart := api.Storage.GetCart(res_id)
		json.NewEncoder(ctx.Writer).Encode(orderCart)

	}
}
func (api *ApiHandler) ChangeStatus() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json")
		order_id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(err)
		}
		api.Storage.ChangeStatus(order_id)
	}
}

func (api *ApiHandler) GetOrderStatus() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json")
		orderId, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(err)
		}
		status := api.Storage.GetStatus(orderId)
		ctx.JSON(http.StatusOK, gin.H{"status": status})
	}
}

func (api *ApiHandler) SetStatusDelivered() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json")
		orderId, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(err)
		}
		api.Storage.UpdateStatusToDelivered(orderId)
	}
}
func (api *ApiHandler) SetStatusCanelled() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json")
		orderId, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(err)
		}
		api.Storage.UpdateStatusToCancelled(orderId)
	}
}

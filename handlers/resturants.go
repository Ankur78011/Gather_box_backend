package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"gatherbox.com/models"
	"github.com/gin-gonic/gin"
)

func (api *ApiHandler) GetResturant() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json")
		ResutrantsList := api.Storage.GetResturants()
		json.NewEncoder(ctx.Writer).Encode(ResutrantsList)
	}
}
func (api *ApiHandler) CreateResturant() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json")
		ownerId := int(ctx.GetFloat64("owner_id"))
		var AddedRes models.NewResturant
		err := ctx.ShouldBindJSON(&AddedRes)
		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "crediential missing"})
			return
		}
		api.Storage.CreateResturant(ownerId, AddedRes.Name, AddedRes.Address, AddedRes.Img_link, AddedRes.Description)
	}
}

func (api *ApiHandler) GetTopMeals() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json")
		RecentlyResName := api.Storage.GetRecentRes()
		json.NewEncoder(ctx.Writer).Encode(RecentlyResName)
	}
}

func (api *ApiHandler) FilterRes() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ownerId := int(ctx.GetFloat64("owner_id"))
		fmt.Println(ownerId)
		ctx.Header("Content-Type", "application/json")
		ListOfRes := api.Storage.GetResList(ownerId)
		json.NewEncoder(ctx.Writer).Encode(ListOfRes)
	}
}

func (api *ApiHandler) DeleteRes() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json")
		res_id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(err)
		}
		api.Storage.DeleteResOrders(res_id)
		api.Storage.DeleteMeals(res_id)
		api.Storage.DeleteResFromData(res_id)
		ctx.Status(http.StatusOK)
	}
}
func (api *ApiHandler) UpdateResturant() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json")
		owner_id := int(ctx.GetFloat64("owner_id"))
		res_id, _ := strconv.Atoi(ctx.Param("id"))
		res := api.Storage.CheckOwner(owner_id, res_id)
		if !res {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		res_Details := api.Storage.SingleResDetails(res_id)
		ctx.JSON(http.StatusOK, res_Details)

	}
}
func (api *ApiHandler) UpdateResData() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json")
		owner_id := int(ctx.GetFloat64("owner_id"))
		res_id, _ := strconv.Atoi(ctx.Param("id"))
		res := api.Storage.CheckOwner(owner_id, res_id)
		if !res {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		var UpdatedValues models.UpdateResturant
		json.NewDecoder(ctx.Request.Body).Decode(&UpdatedValues)
		api.Storage.UpdateRes(UpdatedValues.Name, UpdatedValues.Address, UpdatedValues.Img_link, UpdatedValues.Description, res_id)
	}
}

func (api *ApiHandler) GetMealsResID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json")
		res_id, _ := strconv.Atoi(ctx.Param("id"))
		meals := api.Storage.GetMealsResId(res_id)
		ctx.JSON(http.StatusOK, meals)

	}
}

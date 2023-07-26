package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"gatherbox.com/models"
	"github.com/gin-gonic/gin"
)

func (api *ApiHandler) GetMeals() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json")
		ListOfMeals := api.Storage.GetMeals()
		json.NewEncoder(ctx.Writer).Encode(ListOfMeals)
	}
}
func (api *ApiHandler) GetMealsAsPerPrice() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		price, _ := strconv.Atoi(ctx.Param("price"))
		ListOfMeals := api.Storage.GetMealsAsPerPrice(price)
		json.NewEncoder(ctx.Writer).Encode(ListOfMeals)

	}
}
func (api *ApiHandler) CreateMeals() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json")
		var NewMeal models.NewMeal
		err := ctx.ShouldBindJSON(&NewMeal)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing information"})
			return
		}
		api.Storage.CreateMeal(NewMeal.Name, NewMeal.Description, NewMeal.Price, NewMeal.Img_link, NewMeal.Prep_time, NewMeal.Res_Id)
	}
}

func (api *ApiHandler) GetResMeals() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		ctx.Header("Content-type", "application/json")
		MealsFromRes := api.Storage.GetResWiseMeal(id)
		json.NewEncoder(ctx.Writer).Encode(MealsFromRes)

	}
}
func (api *ApiHandler) PostMeal() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json")
		owner_id := int(ctx.GetFloat64("owner_id"))
		res_id, _ := strconv.Atoi(ctx.Param("id"))
		res := api.Storage.CheckOwner(owner_id, res_id)
		if !res {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		var mealDetails models.NewMeal

		err := ctx.ShouldBindJSON(&mealDetails)
		if err != nil {

			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing information"})
			return
		}
		api.Storage.CreateTheMeal(mealDetails.Name, mealDetails.Img_link, mealDetails.Description, mealDetails.Prep_time, mealDetails.Price, res_id)

	}
}

func (api *ApiHandler) DeleteMeal() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json")
		owner_id := int(ctx.GetFloat64("owner_id"))
		meal_id, _ := strconv.Atoi(ctx.Param("id"))

		result := api.Storage.CheckThatOwnerDeleteItsMeal(owner_id, meal_id)
		if result != true {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		api.Storage.DeleteFromOrder(meal_id)
		api.Storage.DelteFromMealTable(meal_id)

	}
}

func (api *ApiHandler) UpdateMeals() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json")
		owner_id := int(ctx.GetFloat64("owner_id"))
		meal_id, err := strconv.Atoi(ctx.Param(("id")))
		if err != nil {
			panic(err)
		}
		result := api.Storage.CheckThatOwnerDeleteItsMeal(owner_id, meal_id)
		if result != true {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		var UpdatedValues models.NewMeal
		err = ctx.ShouldBindJSON(&UpdatedValues)
		if err != nil {
			panic(err)
		}
		api.Storage.UpdateMeals(UpdatedValues.Name, UpdatedValues.Description, UpdatedValues.Price, UpdatedValues.Prep_time, UpdatedValues.Img_link, meal_id)

	}
}

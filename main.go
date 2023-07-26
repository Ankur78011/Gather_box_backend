package main

import (
	"database/sql"
	"fmt"

	"gatherbox.com/handlers"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "gatherbox"
)

func enableCors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	c.Next()
}
func main() {
	psql := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psql)
	if err != nil {
		fmt.Println("Cannot Connect To DataBase")
		return
	}
	defer db.Close()
	fmt.Println("connected succesfully")
	router := gin.Default()
	router.Use(enableCors)
	apiHandler := handlers.NewApiHandler(db)
	// resturants
	router.GET("/resturants", apiHandler.GetResturant())
	router.POST("/resturants", apiHandler.OwnerToken(), apiHandler.CreateResturant())
	// meals
	router.GET("/meals", apiHandler.GetMeals())
	router.GET("/meals/:price", apiHandler.GetMealsAsPerPrice())
	router.POST("/meals", apiHandler.CreateMeals())
	//resturant wise meals
	router.GET("/resturantsmeals/:id", apiHandler.GetResMeals())
	//rectnyl added resturants
	router.GET("/recentlyaddedresturans", apiHandler.GetTopMeals())
	// //
	router.POST("/signup", apiHandler.CreateUser())
	//login route
	router.POST("/login", apiHandler.Login())
	////token varification
	router.GET("/verifytoken", apiHandler.Token())
	//orderDEtaisl
	router.POST("/postorderdetails", apiHandler.Token(), apiHandler.PostOrderDetails())
	//Cart posting
	router.POST("/cartpost/:order_id", apiHandler.PostCart())
	//owner login
	router.POST("/adminlogin", apiHandler.AdminLogin())
	router.POST("/createAdmin", apiHandler.PostAdmin())
	router.GET("/verifyowner", apiHandler.OwnerToken())
	router.GET("/allresturants", apiHandler.OwnerToken(), apiHandler.FilterRes())
	router.DELETE("/allresturants/:id", apiHandler.OwnerToken(), apiHandler.DeleteRes())
	//Update owenr details
	router.PUT("/update/:id", apiHandler.OwnerToken(), apiHandler.UpdateOwner())
	//Edit Res
	router.GET("editresturant/:id", apiHandler.OwnerToken(), apiHandler.UpdateResturant())
	router.POST("editresturant/:id", apiHandler.OwnerToken(), apiHandler.UpdateResData())
	//Getingmealsaccoring to res id
	router.GET("resturantMeals/:id", apiHandler.OwnerToken(), apiHandler.GetMealsResID())
	router.POST("resturantmeals/:id", apiHandler.OwnerToken(), apiHandler.PostMeal())
	router.DELETE("deletemeal/:id", apiHandler.OwnerToken(), apiHandler.DeleteMeal())
	//edit meals
	router.PUT("updatemeals/:id", apiHandler.OwnerToken(), apiHandler.UpdateMeals())
	//orders
	router.GET("orders/:id", apiHandler.OwnerToken(), apiHandler.GetOrders())
	router.GET("changeorderstatus/:id", apiHandler.OwnerToken(), apiHandler.ChangeStatus())
	//GetOrderStatus
	router.GET("orderstatus/:id", apiHandler.OwnerToken(), apiHandler.GetOrderStatus())
	router.GET("orderstatusdelivered/:id", apiHandler.OwnerToken(), apiHandler.SetStatusDelivered())
	router.GET("orderstatuscancelled/:id", apiHandler.OwnerToken(), apiHandler.SetStatusCanelled())
	router.Run(":8000")
}

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/itnopadol/sys_api/ctrl"
	"gopkg.in/gin-contrib/cors.v1"
)

func main(){
	r := gin.New()

	r.Use(cors.Default())
	r.GET("/user",ctrl.UserGetById)
	r.GET("/users/search",ctrl.UserGetByKeyword)
	r.GET("/users",ctrl.UserGetAll)
	r.POST("/user",ctrl.UserSave)
	r.PUT("/user",ctrl.UserUpdate)
	r.PUT("/user/disable",ctrl.UserDisable)

	r.GET("/app",ctrl.AppGetById)
	r.GET("/apps/search",ctrl.AppGetByKeyword)
	r.GET("/apps",ctrl.AppGetAll)
	r.POST("/app",ctrl.AppSave)
	r.PUT("/app",ctrl.AppUpdate)
	r.PUT("/app/disable",ctrl.AppDisable)



	r.Run(":9000")

}
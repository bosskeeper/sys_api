package main

import (
	"github.com/gin-gonic/gin"
	"github.com/itnopadol/sys_api/ctrl"
	"gopkg.in/gin-contrib/cors.v1"
	//"github.com/gin-gonic/contrib/cors"
)

func main(){
	r := gin.New()

	r.Use(cors.Default())
	//r.GET("/user",ctrl.UserGetById)
	r.GET("/user",ctrl.UserGetByUserCode)
	r.GET("/users/search",ctrl.UserGetByKeyword)
	r.GET("/users",ctrl.UserGetAll)
	r.POST("/user",ctrl.UserSave)
	r.PUT("/user",ctrl.UserUpdate)
	r.PUT("/user/disable",ctrl.UserDisable)

	//r.GET("/app",ctrl.AppGetById)
	r.GET("/app",ctrl.AppGetByAppCode)
	r.GET("/apps/search",ctrl.AppGetByKeyword)
	r.GET("/apps",ctrl.AppGetAll)
	r.POST("/app",ctrl.AppSave)
	r.PUT("/app",ctrl.AppUpdate)
	r.PUT("/app/disable",ctrl.AppDisable)
	r.GET("/app/role",ctrl.AppGetByRole)

	r.GET("/menu",ctrl.MenuGetById)
	r.GET("/menus/search",ctrl.MenuGetByKeyword)
	r.GET("/menu/app_id",ctrl.MenuGetByAppId)
	r.GET("/menus",ctrl.MenuGetAll)
	r.POST("/menu",ctrl.MenuSave)
	r.PUT("/menu",ctrl.MenuUpdate)
	r.PUT("/menu/disable",ctrl.MenuDisable)

	r.GET("/login",ctrl.LoginGet)

	r.GET("/role",ctrl.RoleGetById)
	r.GET("/roles/search",ctrl.RoleGetByKeyword)
	r.GET("/roles",ctrl.RoleGetAll)
	r.POST("/role",ctrl.RoleSave)
	r.PUT("/role",ctrl.RoleUpdate)
	r.PUT("/role/disable",ctrl.RoleDisable)

	r.GET("/userroles",ctrl.UserRoleGetAll)
	r.GET("/userrole",ctrl.UserRoleGetUser)
	r.POST("/userrole",ctrl.UserRoleSave)
	r.PUT("/userrole",ctrl.UserRoleUpdate)

	r.GET("/permissions",ctrl.PermissionGetAll)
	r.GET("/permission",ctrl.PermissionGetByMenu)
	r.POST("/permission",ctrl.PermissionSave)
	r.PUT("/permission",ctrl.PermissionUpdate)

	r.Run(":9000")

}
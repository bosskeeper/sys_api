package ctrl

import (
	"fmt"
	"net/http"
	"github.com/itnopadol/sys_api/api"
	"github.com/itnopadol/sys_api/model"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	//"strconv"
	"strconv"
	//"golang.org/x/tools/go/gcimporter15/testdata"
)

func PermissionGetAll(c *gin.Context){
	log.Println("call GET PermissionGetAll")
	c.Keys = headerKeys

	access_token := c.Request.URL.Query().Get("access_token")
	app_id := c.Request.URL.Query().Get("app_id")
	role_id := c.Request.URL.Query().Get("role_id")

	p := new(model.Permission)
	p.AppId, _ = strconv.ParseInt(app_id,10,64)
	p.RoleId, _ = strconv.ParseInt(role_id,10,64)
	permissions,err := p.PermissionGetAll(dbc,access_token,p.AppId,p.RoleId)
	if err != nil {
		fmt.Println(err)
	}

	rs := api.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content ="+err.Error()
		c.JSON(http.StatusNotFound,rs)
	}else{
		if permissions==nil{
			//fmt.Println("Yes")
			rs.Status = "error"
			rs.Message = "No Content: NotData"
			c.JSON(http.StatusNotFound, rs)
		}else {
			rs.Status = "success"
			rs.Data = permissions
			c.JSON(http.StatusOK,rs)
		}
	}
}

func PermissionGetByMenu(c *gin.Context){
	log.Println("call GET PermissionGetByMenu")
	c.Keys = headerKeys

	access_token := c.Request.URL.Query().Get("access_token")
	app_id := c.Request.URL.Query().Get("app_id")
	role_id := c.Request.URL.Query().Get("role_id")
	menu_id := c.Request.URL.Query().Get("menu_id")

	p := new(model.Permission)
	p.AppId, _ = strconv.ParseInt(app_id,10,64)
	p.RoleId, _ = strconv.ParseInt(role_id,10,64)
	p.MenuId, _ = strconv.ParseInt(menu_id,10,64)
	err := p.PermissionGetByMenu(dbc,access_token,p.AppId,p.RoleId,p.MenuId)
	if err != nil {
		fmt.Println(err)
	}

	rs := api.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content ="+err.Error()
		c.JSON(http.StatusNotFound,rs)
	}else{
		if p==nil{
			//fmt.Println("Yes")
			rs.Status = "error"
			rs.Message = "No Content: NotData"
			c.JSON(http.StatusNotFound, rs)
		}else {
			rs.Status = "success"
			rs.Data = p
			c.JSON(http.StatusOK,rs)
		}
	}
}


func PermissionSave(c *gin.Context){
	log.Println("call POST PermissionSave")
	c.Keys = headerKeys

	//p := model.Permission{}
	//data := []model.Permission{}
	//data := Permissions{}
	x := model.Permissions{}
	err := c.BindJSON(&x)
	fmt.Println(x)

	if err != nil {
		fmt.Println("Binding Json Error Step ",err)
	}

	x.PermissionSaveAll(dbc)

	rs := api.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content = "+err.Error()
		c.JSON(http.StatusNotFound,rs)
	}else{

			rs.Status = "success"
			rs.Data = x
			c.JSON(http.StatusOK,rs)

	}
}

/*
func PermissionUpdate(c *gin.Context){
	c.Keys = headerKeys
	newPermission := &model.Permission{}

	err := c.BindJSON(newPermission)
	fmt.Println("xx",newPermission)
	if err != nil {
		fmt.Println(err)
	}

	ur, _ := newPermission.PermissionUpdate(dbc)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(ur)

	rs := api.Response{}
	if err != nil {
		rs.Status="error"
		rs.Message="No Content "+ err.Error()
		c.JSON(http.StatusNotFound,rs)
	}else {
		if ur==0{
			rs.Status = "error"
			rs.Message = "No Content = Data Not Found "
			c.JSON(http.StatusNotFound,rs)
		}else {
			rs.Status = "success"
			rs.Data = ur
			c.JSON(http.StatusOK,rs)
		}
	}
}
*/

func PermissionUpdateAll(c *gin.Context){
	log.Println("call Put PermissionUpdate")
	c.Keys = headerKeys

	//p := model.Permission{}
	//data := []model.Permission{}
	//data := Permissions{}
	x := model.Permissions{}
	err := c.BindJSON(&x)
	//fmt.Println(x)

	if err != nil {
		fmt.Println("Binding Json Error Step ",err)
	}

	x.PermissionUpdateAll(dbc)

	rs := api.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content = "+err.Error()
		c.JSON(http.StatusNotFound,rs)
	}else{

			rs.Status = "success"
			//rs.Data = x
			c.JSON(http.StatusOK,rs)

	}
}
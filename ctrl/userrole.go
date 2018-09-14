package ctrl

import (
	"fmt"
	"net/http"
	"github.com/itnopadol/sys_api/api"
	"github.com/itnopadol/sys_api/model"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
)

func UserRoleGetAll(c *gin.Context){

	log.Println("call GET UserRoleAll")
	c.Keys = headerKeys

	access_token := c.Request.URL.Query().Get("access_token")
	user_id := c.Request.URL.Query().Get("user_id")

	ur := new(model.UserRole)
	ur.UserId, _ = strconv.ParseInt(user_id,10,64)
	userroles,err := ur.UserRoleGetAll(dbc,access_token,ur.UserId)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("UserCode = ",user_id)

	rs := api.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content ="+err.Error()
		c.JSON(http.StatusNotFound,rs)
	} else {
		rs.Status = "success"
		rs.Data = userroles
		c.JSON(http.StatusOK,rs)
}
}
	//else{
	//	if userroles==nil{
	//		//fmt.Println("Yes")
	//		rs.Status = "error"
	//		rs.Message = "No Content: NotData"
	//		c.JSON(http.StatusNotFound, rs)
	//	}


func UserRoleGetUser(c *gin.Context){
	log.Println("call GET UserRoleByUser")
	c.Keys = headerKeys

	access_token := c.Request.URL.Query().Get("access_token")
	app_id := c.Request.URL.Query().Get("app_id")
	user_id := c.Request.URL.Query().Get("user_id")

	ur := new(model.UserRole)
	ur.AppId, _ = strconv.ParseInt(app_id,10,64)
	ur.UserId, _ = strconv.ParseInt(user_id,10,64)

	err := ur.UserRoleGetUser(dbc,access_token,ur.AppId,ur.UserId)
	if err != nil {
		fmt.Println(err)
	}

	rs := api.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content ="+err.Error()
		c.JSON(http.StatusNotFound,rs)
	}else{
		if ur==nil{
			//fmt.Println("Yes")
			rs.Status = "error"
			rs.Message = "No Content: NotData"
			c.JSON(http.StatusNotFound, rs)
		}else {
			//rs.Status = "success"
			//rs.Data = ur
			c.JSON(http.StatusOK,rs)
		}
	}
}

func UserRoleSave(c *gin.Context){
	log.Println("call POST UserRoleSave")
	c.Keys = headerKeys

	newUserRole := &model.UserRole{}
	err := c.BindJSON(newUserRole)
	if err != nil {
		fmt.Println(err)
	}
	a, _ := newUserRole.UserRoleSave(dbc)

	rs := api.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content = "+err.Error()
		c.JSON(http.StatusNotFound,rs)
	}else{
		if a==0{
			rs.Status = "error"
			rs.Message = "No Content = Duplicate key"
			c.JSON(http.StatusNotFound,rs)
		}else {
			rs.Status = "success"
			rs.Data = a
			c.JSON(http.StatusOK,rs)
			}
	}
}

func UserRoleUpdate(c *gin.Context){
	c.Keys = headerKeys
	newUserRole := &model.UserRole{}

	err := c.BindJSON(newUserRole)
	if err != nil {
		fmt.Println(err)
	}

	ur, _ := newUserRole.UserRoleUpdate(dbc)
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
			rs.Message = "No Content = User Not Found "
			c.JSON(http.StatusNotFound,rs)
		}else {
			rs.Status = "success"
			rs.Data = ur
			c.JSON(http.StatusOK,rs)
		}
	}
}

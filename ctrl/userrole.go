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
	app_id := c.Request.URL.Query().Get("app_id")

	ur := new(model.UserRole)
	ur.AppId, _ = strconv.ParseInt(app_id,10,64)
	userroles,err := ur.UserRoleGetAll(dbc,access_token,ur.AppId)
	if err != nil {
		fmt.Println(err)
	}
	rs := api.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content ="+err.Error()
		c.JSON(http.StatusNotFound,rs)
	}else{
		if userroles==nil{
			//fmt.Println("Yes")
			rs.Status = "error"
			rs.Message = "No Content: NotData"
			c.JSON(http.StatusNotFound, rs)
		}else {
			rs.Status = "success"
			rs.Data = userroles
			c.JSON(http.StatusOK,rs)
		}
	}
}

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
			rs.Status = "success"
			rs.Data = ur
			c.JSON(http.StatusOK,rs)
		}
	}
}

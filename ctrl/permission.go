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

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
)

func init(){
	db, err := ConnectDB("sys")
	if err != nil {
		fmt.Println(err.Error())
	}
	dbc = db
}

func RoleGetAll(c *gin.Context){
	log.Println("call GET RoleGetAll")
	c.Keys = headerKeys

	r := new(model.Role)
	roles, err := r.RoleGetAll(dbc)
	if err != nil {
		fmt.Println(err)
	}

	rs := api.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content ="+err.Error()
		c.JSON(http.StatusNotFound,rs)
	}else{
		if roles==nil{
			//fmt.Println("Yes")
			rs.Status = "error"
			rs.Message = "No Content: NotData"
			c.JSON(http.StatusNotFound, rs)
		}else {
			rs.Status = "success"
			rs.Data = roles
			c.JSON(http.StatusOK,rs)
		}
	}
}

func RoleGetByKeyword(c *gin.Context){
	log.Println("call GET RoleGetByKeyword")
	c.Keys = headerKeys

	access_token := c.Request.URL.Query().Get("access_token")
	keyword := c.Request.URL.Query().Get("keyword")

	r := new(model.Role)
	roles, err := r.RoleGetByKeyword(dbc,access_token,keyword)
	if err != nil {
		fmt.Println(err)
	}

	rs := api.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message= "No Content"+err.Error()
		c.JSON(http.StatusNotFound,rs)
	}else{
		if roles==nil{
			//fmt.Println("Yes")
			rs.Status = "error"
			rs.Message = "No Content: NotData"
			c.JSON(http.StatusNotFound, rs)
		}else {
			rs.Status = "success"
			rs.Data = roles
			c.JSON(http.StatusOK,rs)
		}
	}
}

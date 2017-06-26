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
	//"github.com/matryer/m"
)

func init(){
	db, err := ConnectDB("sys")
	if err != nil {
		fmt.Println(err.Error())
	}
	dbc = db
}

func RoleGetAll(c *gin.Context){
	log.Println("call GET RoleGetAll ")
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
		rs.Status = "success"
		rs.Data = roles
		c.JSON(http.StatusOK,rs)
	}
}

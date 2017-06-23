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
	//"github.com/matryer/m"
	//"strconv"
	"strconv"
)

func LoginGet(c*gin.Context){
	log.Println("call GET Login")
	c.Keys = headerKeys

	access_token := c.Request.URL.Query().Get("access_token")
	user_name := c.Request.URL.Query().Get("username")
	password := c.Request.URL.Query().Get("password")
	appid := c.Request.URL.Query().Get("appid")

	log.Println("call GET username",user_name)
	log.Println("call GET password",password)


	//l := new(model.Login)
	l := model.Login{}


	l.AppID, _ = strconv.ParseInt(appid,10,64)
	 err := l.LoginGetByUser(dbc,access_token,user_name,password,l.AppID)

	fmt.Println("ctrl l: ",l)
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
		rs.Data = l
		c.JSON(http.StatusOK,rs)
	}
}

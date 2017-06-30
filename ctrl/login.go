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
	user_code := c.Request.URL.Query().Get("usercode")
	password := c.Request.URL.Query().Get("password")
	appid := c.Request.URL.Query().Get("appid")

	log.Println("call GET username",user_code)
	log.Println("call GET password",password)


	//l := new(model.Login)
	l := model.Login{}


	l.AppID, _ = strconv.ParseInt(appid,10,64)
	 err := l.LoginGetByUser(dbc,access_token,user_code,password,l.AppID)

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
		if l.Id==0{
			rs.Status = "error"
			rs.Message = "No Content = UserName or Password Invalid"
			c.JSON(http.StatusNotFound,rs)
		}else if l.Id != 0 && l.Menus == nil{
			rs.Status = "error"
			rs.Message = "No Content = UserName Not Permission"
			c.JSON(http.StatusNotFound,rs)
		}else {
			rs.Status = "success"
			rs.Data = l
			c.JSON(http.StatusOK,rs)
			}
	}
}

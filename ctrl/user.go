package ctrl

import (
	"github.com/gin-gonic/gin"
	"github.com/itnopadol/sys_api/model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"github.com/itnopadol/sys_api/api"
	"log"
	"strconv"
)

//var  dbc *sqlx.DB
//
//func init(){
//	db, err := ConnectDB("sys")
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//	dbc = db
//}

//=======================================================API User========================================================
func UserGetById(c *gin.Context){
	log.Println("call GET UserGetByCode()")
	c.Keys = headerKeys

	access_token := c.Request.URL.Query().Get("access_token")
	user_id := c.Request.URL.Query().Get("user_id")

	u := new(model.User)
	u.Id, _ = strconv.ParseInt(user_id, 10, 64)
	err := u.UserGetById(dbc,access_token,u.Id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("UserCode = ",user_id,u.UserCode,u.UserName)

	rs := api.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content: " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = u
		c.JSON(http.StatusOK, rs)
	}

}

func UserGetByUserCode(c *gin.Context){
	log.Println("call GET UserGetByUserCode()")
	c.Keys = headerKeys

	access_token := c.Request.URL.Query().Get("access_token")
	user_code := c.Request.URL.Query().Get("user_code")

	u := new(model.User)
	u.UserCode = user_code
	err := u.UserGetByUserCode(dbc,access_token,u.UserCode)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("UserCode = ",user_code,u.UserCode,u.UserName)

	rs := api.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content: " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = u
		c.JSON(http.StatusOK, rs)
	}

}

func UserGetByKeyword(c *gin.Context){
	log.Println("call GET UserGetByKeyword()")
	c.Keys = headerKeys

	//keyword := c.Param("keyword")
	keyword := c.Request.URL.Query().Get("keyword")
	access_token := c.Request.URL.Query().Get("access_token")

	fmt.Println("Keyword = ",keyword)
	u := new(model.User)
	users, err := u.UserGetByKeyword(dbc,access_token,keyword)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200,users)
}

func UserGetAll(c *gin.Context){
	log.Println("call GET UserGetAll()")
	c.Keys = headerKeys

	access_token := c.Request.URL.Query().Get("access_token")
	u := new(model.User)
	users, err := u.UserGetAll(dbc,access_token)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200,users)
}

func UserSave(c *gin.Context){
	log.Println("call POST UserSave()")
	c.Keys = headerKeys


	newUser := &model.User{}
	err := c.BindJSON(newUser)
	if err != nil {
		fmt.Println(err.Error())
	}

	u, _ := newUser.UserSave(dbc)

	rs := api.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content: " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = u
		c.JSON(http.StatusOK, rs)
	}


}

func UserUpdate(c *gin.Context){
	log.Println("call PUT UserUpdate()")
	c.Keys = headerKeys

	user := &model.User{}
	err := c.BindJSON(user)
	if err != nil {
		fmt.Println(err)
	}
	u, _ := user.UserUpdate(dbc)

	rs := api.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content: " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	}else{
		rs.Status = "success"
		rs.Data = u
		c.JSON(http.StatusOK, rs)
	}

}


func UserDisable(c *gin.Context){
	log.Println("call PUT UserUpdate()")
	c.Keys = headerKeys

	user := &model.User{}
	err := c.BindJSON(user)
	if err != nil {
		fmt.Println(err)
	}
	u, _ := user.UserDisable(dbc)

	rs := api.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content: " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	}else{
		rs.Status = "success"
		rs.Data = u
		c.JSON(http.StatusOK, rs)
	}

}


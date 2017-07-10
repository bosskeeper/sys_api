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

//=======================================================API App========================================================


func init(){
	db, err := ConnectDB("sys")
	if err != nil {
		fmt.Println(err.Error())
	}
	dbc = db
}

func AppGetAll(c *gin.Context){
	log.Println("call GET AppGetAll()")
	c.Keys = headerKeys

	newApp := &model.App{}
	apps, err := newApp.AppGetAll(dbc)
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
		rs.Data = apps
		c.JSON(http.StatusOK,rs)
	}
}

func AppGetById(c *gin.Context){
	log.Println("call GET AppGetById()")
	c.Keys = headerKeys

	access_token := c.Request.URL.Query().Get("access_token")
	app_id := c.Request.URL.Query().Get("app_id")

	a := new(model.App)
	a.Id, _ = strconv.ParseInt(app_id,10,64)
	err := a.AppGetById(dbc, access_token, a.Id)
	if err != nil {
		fmt.Println(err)
	}
	rs := api.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content: " + err.Error()
		c.JSON(http.StatusNotFound,rs)
	}else{
		rs.Status = "success"
		rs.Data = a
		c.JSON(http.StatusOK,rs)
	}
}

func AppGetByRole(c *gin.Context){
	log.Println("call GET AppByRole")
	c.Keys = headerKeys

	access_token := c.Request.URL.Query().Get("access_token")
	app_id := c.Request.URL.Query().Get("app_id")

	a := new(model.App)
	a.Id, _ = strconv.ParseInt(app_id,10,64)
	apps, err := a.AppGetByRole(dbc,access_token,a.Id)
	if err != nil {
		fmt.Println(err)
	}

	rs := api.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message= "No Content"+err.Error()
		c.JSON(http.StatusNotFound,rs)
	}else{
		rs.Status = "success"
		rs.Data = apps
		c.JSON(http.StatusOK,rs)
	}
}

func AppGetByAppCode(c *gin.Context){
	log.Println("call GET AppGetByAppCode()")
	c.Keys = headerKeys

	access_token := c.Request.URL.Query().Get("access_token")
	app_code := c.Request.URL.Query().Get("app_code")

	a := new(model.App)
	a.AppCode = app_code
	err := a.AppGetByAppCode(dbc, access_token, a.AppCode)
	if err != nil {
		fmt.Println(err)
	}
	rs := api.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content: " + err.Error()
		c.JSON(http.StatusNotFound,rs)
	}else{
		rs.Status = "success"
		rs.Data = a
		c.JSON(http.StatusOK,rs)
	}
}

func AppGetByKeyword(c *gin.Context){
	log.Println("call GET AppGetByKeyword()")
	c.Keys = headerKeys

	access_token := c.Request.URL.Query().Get("access_token")
	keyword := c.Request.URL.Query().Get("keyword")

	a := new(model.App)
	apps, err := a.AppGetByKeyword(dbc,access_token,keyword)
	if err != nil {
		fmt.Println(err)
	}

	rs := api.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message= "No Content"+err.Error()
		c.JSON(http.StatusNotFound,rs)
	}else{
		rs.Status = "success"
		rs.Data = apps
		c.JSON(http.StatusOK,rs)
	}
}

func AppSave(c *gin.Context){
	log.Println("call POST AppSave()")
	c.Keys = headerKeys

	newApp := &model.App{}
	err := c.BindJSON(newApp)
	if err != nil {
		fmt.Println(err)
	}
	a, _ := newApp.AppSave(dbc)

	rs := api.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content = "+err.Error()
		c.JSON(http.StatusNotFound,rs)
	}else{
		rs.Status = "success"
		rs.Data = a
		c.JSON(http.StatusOK,rs)
	}
}

func AppUpdate(c *gin.Context){
	log.Println("call PUT AppUpdate()")

	newApp := &model.App{}
	err := c.BindJSON(newApp)
	if err != nil {
		fmt.Println(err)
	}

	a, err := newApp.AppUpdate(dbc)
	if err != nil {
		fmt.Println(err)
	}
	rs := api.Response{}
	if err != nil {
		rs.Status="error"
		rs.Message="No Content "+ err.Error()
		c.JSON(http.StatusNotFound,rs)
	}else {
		rs.Status="success"
		rs.Data =a
		c.JSON(http.StatusOK,rs)
	}

}

func AppDisable(c *gin.Context){
	log.Println("call PUT AppUpdate()")

	newApp := &model.App{}
	err := c.BindJSON(newApp)
	if err != nil {
		fmt.Println(err)
	}

	a, err := newApp.AppDisable(dbc)
	if err != nil {
		fmt.Println(err)
	}
	rs := api.Response{}
	if err != nil {
		rs.Status="error"
		rs.Message="No Content "+ err.Error()
		c.JSON(http.StatusNotFound,rs)
	}else {
		rs.Status="success"
		rs.Data =a
		c.JSON(http.StatusOK,rs)
	}

}

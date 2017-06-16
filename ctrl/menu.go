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

//=======================================================API Menu========================================================


func init(){
	db, err := ConnectDB("sys")
	if err != nil {
		fmt.Println(err.Error())
	}
	dbc = db
}

func MenuGetAll(c *gin.Context){
	log.Println("call GET MenuGetAll()")
	c.Keys = headerKeys

	m := new(model.Menu)
	menus, err := m.MenuGetAll(dbc)
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
		rs.Data = menus
		c.JSON(http.StatusOK,rs)
	}
}

func MenuGetById(c *gin.Context){
	log.Println("call GET AppGetById()")
	c.Keys = headerKeys

	access_token := c.Request.URL.Query().Get("access_token")
	menu_id := c.Request.URL.Query().Get("menu_id")

	m := new(model.Menu)
	m.Id, _ = strconv.ParseInt(menu_id,10,64)
	err := m.MenuGetById(dbc, access_token, m.Id)
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
		rs.Data = m
		c.JSON(http.StatusOK,rs)
	}
}

func MenuGetByAppId(c *gin.Context){
	log.Println("call GET MenuGetByAppId()")
	c.Keys = headerKeys

	access_token := c.Request.URL.Query().Get("access_token")
	app_id := c.Request.URL.Query().Get("app_id")

	m := new(model.Menu)
	m.AppId, _ = strconv.ParseInt(app_id,10,64)
	menus, err := m.MenuGetByAppId(dbc,access_token,m.AppId)
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
		rs.Data = menus
		c.JSON(http.StatusOK,rs)
	}
}

func MenuGetByKeyword(c *gin.Context){
	log.Println("call GET AppGetByKeyword()")
	c.Keys = headerKeys

	access_token := c.Request.URL.Query().Get("access_token")
	keyword := c.Request.URL.Query().Get("keyword")

	m := new(model.App)
	menus, err := m.AppGetByKeyword(dbc,access_token,keyword)
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
		rs.Data = menus
		c.JSON(http.StatusOK,rs)
	}
}

func MenuSave(c *gin.Context){
	log.Println("call POST AppSave()")
	c.Keys = headerKeys

	newMenu := &model.Menu{}
	err := c.BindJSON(newMenu)
	if err != nil {
		fmt.Println(err)
	}
	m, _ := newMenu.MenuSave(dbc)

	rs := api.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content = "+err.Error()
		c.JSON(http.StatusNotFound,rs)
	}else{
		rs.Status = "success"
		rs.Data = m
		c.JSON(http.StatusOK,rs)
	}
}

func MenuUpdate(c *gin.Context){
	log.Println("call PUT AppUpdate()")

	newMenu := &model.Menu{}
	err := c.BindJSON(newMenu)
	if err != nil {
		fmt.Println(err)
	}

	m, err := newMenu.MenuUpdate(dbc)
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
		rs.Data = m
		c.JSON(http.StatusOK,rs)
	}

}

func MenuDisable(c *gin.Context){
	log.Println("call PUT AppUpdate()")

	newMenu := &model.App{}
	err := c.BindJSON(newMenu)
	if err != nil {
		fmt.Println(err)
	}

	m, err := newMenu.AppDisable(dbc)
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
		rs.Data = m
		c.JSON(http.StatusOK,rs)
	}

}

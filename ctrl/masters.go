package ctrl

import (
	"github.com/gin-gonic/gin"
	"github.com/itnopadol/sys_api/model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	//"net/http"
	//"github.com/itnopadol/sys_api/api"
	"log"
	//"strconv"
)

func GetBranchs(c *gin.Context){

	log.Println("call GET BranchMaster")
	c.Keys = headerKeys

	access_token := c.Request.URL.Query().Get("access_token")
	b := new(model.BranchMaster)
	branchs, err := b.GetBranchs(dbc,access_token)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200,branchs)
}

func GetDepartments(c *gin.Context){

	log.Println("call GET DepartmentMaster")
	c.Keys = headerKeys

	access_token := c.Request.URL.Query().Get("access_token")
	d := new(model.DepartmentMaster)
	departments, err := d.GetDepartments(dbc,access_token)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200,departments)
}

func GetExperts(c *gin.Context){

	log.Println("call GET ExpertMaster")
	c.Keys = headerKeys

	access_token := c.Request.URL.Query().Get("access_token")
	e := new(model.ExpertMaster)
	experts, err := e.GetExperts(dbc,access_token)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200,experts)
}

func GetProfitcenters(c *gin.Context){

	log.Println("call GET ProfitcenterMaster")
	c.Keys = headerKeys

	access_token := c.Request.URL.Query().Get("access_token")
	p := new(model.ProfitcenterMaster)
	profitcenters, err := p.GetProfitcenters(dbc,access_token)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200,profitcenters)
}
package model

import (
	"github.com/jmoiron/sqlx"
	"fmt"
)

type Permission struct {
	Id 	 int64 `json:"id" db:"Id"`
	AppId  int64 `json:"appid" db:"AppId"`
	AppCode  string `json:"appcode" db:"AppCode"`
	AppName  string `json:"appname" db:"AppName"`
	RoleId int64 `json:"roleid" db:"RoleId"`
	RoleCode string `json:"rolecode" db:"RoleCode"`
	RoleName string `json:"rolename" db:"RoleName"`
	MenuId int64 `json:"menuid" db:"MenuId"`
	MenuCode string `json:"menucode" db:"MenuCode"`
	MenuName string `json:"menuname" db:"MenuName"`
	Create int64 `json:"create" db:"Create"`
	Read int64 `json:"read" db:"Read"`
	Update int64 `json:"update" db:"Update"`
	Delete int64 `json:"delete" db:"Delete"`
}

func (p *Permission) PermissionGetAll(db *sqlx.DB, access_token string, app_id int64, role_id int64) (permissions []*Permission, err error){
	sql := `select a.Id,a.AppId,b.AppCode,b.AppName,a.RoleId,c.RoleCode,c.RoleName`+
		` ,a.MenuId,d.MenuCode,d.MenuName,a.Create,a.Read,a.Update,a.Delete`+
		` from Permission as a left join App as b on a.AppId=b.Id`+
		` left join Role as c on a.RoleId=c.Id`+
		` left join Menu as d on a.MenuId=d.Id where a.AppId=? and a.RoleId=?`
	p.AppId = app_id
	p.RoleId=role_id
	err = db.Select(&permissions,sql,p.AppId,p.RoleId)
	if err != nil {
		fmt.Println(err)
		return permissions,err
	}
	fmt.Println("RoleCode = ",p.RoleCode)
	return permissions,nil
}

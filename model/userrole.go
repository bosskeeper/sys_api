package model

import (
	"github.com/jmoiron/sqlx"
	"fmt"
)

type UserRole struct {
	UserRoleId       int64 `json:"user_role_id" db:"UserRoleId"`
	AppId       int64 `json:"app_id" db:"AppId"`
	AppCode string `json:"app_code" db:"AppCode"`
	AppName string `json:"app_name" db:"AppName"`
	UserId       int64 `json:"user_id" db:"UserId"`
	UserCode string `json:"user_code" db:"UserCode"`
	UserName string `json:"user_name" db:"UserName"`
	RoleId       int64 `json:"role_id" db:"RoleId"`
	RoleCode string `json:"role_code" db:"RoleCode"`
	RoleName string `json:"role_name" db:"RoleName"`
}

func (ur *UserRole)UserRoleGetAll(db *sqlx.DB, access_token string, app_id int64) (userroles []*UserRole,err error){
	sql := `select a.Id as UserRoleId,a.AppId,b.AppCode,b.AppName,a.UserId,c.UserCode,c.UserName`+
		` ,a.RoleId,d.RoleCode,d.RoleName from UserRole as a`+
		` left join App as b on a.AppId=b.Id`+
		` left join User as c on a.UserId=c.Id`+
		` left join Role as d on a.RoleId=d.Id where a.AppId=?`
	ur.AppId=app_id
	err = db.Select(&userroles,sql,ur.AppId)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return userroles, nil
}

func (ur *UserRole) UserRoleGetUser (db *sqlx.DB, access_token string, app_id int64,user_id int64)  error {
	sql := `select a.Id as UserRoleId,a.AppId,b.AppCode,b.AppName,a.UserId,c.UserCode,c.UserName`+
		` ,a.RoleId,d.RoleCode,d.RoleName from UserRole as a`+
		` left join App as b on a.AppId=b.Id`+
		` left join User as c on a.UserId=c.Id`+
		` left join Role as d on a.RoleId=d.Id where a.AppId=? and a.UserId=?`
	ur.AppId=app_id
	ur.UserId=user_id
	err := db.Get(ur,sql,ur.AppId,ur.UserId)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("RoleCode = ",ur.RoleCode)
	return nil
}
package model

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	"time"
)

type Permission struct {
	Id 	 int64 `json:"id" db:"Id"`
	AppId  int64 `json:"app_id" db:"AppId"`
	AppCode  string `json:"app_code" db:"AppCode"`
	AppName  string `json:"app_name" db:"AppName"`
	RoleId int64 `json:"role_id" db:"RoleId"`
	RoleCode string `json:"role_code" db:"RoleCode"`
	RoleName string `json:"role_name" db:"RoleName"`
	MenuId int64 `json:"menu_id" db:"MenuId"`
	MenuCode string `json:"menu_code" db:"MenuCode"`
	MenuName string `json:"menu_name" db:"MenuName"`
	IsCreate int64 `json:"is_create" db:"IsCreate"`
	IsRead int64 `json:"is_read" db:"IsRead"`
	IsUpdate int64 `json:"is_update" db:"IsUpdate"`
	IsDelete int64 `json:"is_delete" db:"IsDelete"`
	CreatorId int `json:"creator_id" db:"CreatorId"`
	CreateDateTime string `json:"create_date_time,omitempty" db:"CreateDateTime"`
	EditorId int `json:"editor_id" db:"EditorId"`
	EditDateTime string `json:"edit_date_time,omitempty" db:"EditDateTime"`
}

func (p *Permission) PermissionGetAll(db *sqlx.DB, access_token string, app_id int64, role_id int64) (permissions []*Permission, err error){
	sql := `select a.Id,a.AppId,b.AppCode,b.AppName,a.RoleId,c.RoleCode,c.RoleName`+
		` ,a.MenuId,d.MenuCode,d.MenuName,a.IsCreate,a.IsRead,a.IsUpdate,a.IsDelete`+
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

func (p *Permission) PermissionGetByMenu (db *sqlx.DB, access_token string, app_id int64, role_id int64, menu_id int64)  error {
	sql := `select a.Id,a.AppId,b.AppCode,b.AppName,a.RoleId,c.RoleCode,c.RoleName`+
		` ,a.MenuId,d.MenuCode,d.MenuName,a.IsCreate,a.IsRead,a.IsUpdate,a.IsDelete`+
		` from Permission as a left join App as b on a.AppId=b.Id`+
		` left join Role as c on a.RoleId=c.Id`+
		` left join Menu as d on a.MenuId=d.Id where a.AppId=? and a.RoleId=? and a.MenuId=?`
	p.AppId = app_id
	p.RoleId=role_id
	p.MenuId=menu_id
	err := db.Get(p,sql,p.AppId,p.RoleId,p.MenuId)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("RoleCode = ",p.RoleCode)
	return nil
}

func (p *Permission) PermissionSave(db *sqlx.DB) (permission_id int64, err error){
	fmt.Println("Permission = ",p.Id)
	err = p.GetPermissionNotExist(db)

	p.CreateDateTime = time.Now().String()
	p.CreatorId = 1
	sql := `insert into Permission(AppId,RoleId,MenuId,IsCreate,IsRead,IsUpdate,IsDelete,CreatorId,CreateDateTime) values(?,?,?,?,?,?,?,?,?)`
	fmt.Println("sql = ",sql,p.AppId,p.RoleId,p.MenuId,p.IsCreate,p.IsRead,p.IsUpdate,p.IsDelete,p.CreatorId,p.CreateDateTime)
	res, err := db.Exec(sql,p.AppId,p.RoleId,p.MenuId,p.IsCreate,p.IsRead,p.IsUpdate,p.IsDelete,p.CreatorId,p.CreateDateTime)
	if err != nil {
		fmt.Println(err)
		return 0,err
	}
	permission_id = p.Id
	id, _ := res.LastInsertId()
	fmt.Println("Last Insert Id = ",id)
	return id, nil
}

func (p *Permission) GetPermissionNotExist(db *sqlx.DB) error {
	sql := `select Id from Permission where Id = ?`
	err := db.Get(p, sql, p.Id)
	fmt.Println("Get Permission =",p.Id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
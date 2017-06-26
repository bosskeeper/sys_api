package model

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

type Role struct {
	Id int64 `json:"id" db:"Id"`
	RoleCode string `json:"role_code" db:"RoleCode"`
	RoleName string `json:"role_name" db:"RoleName"`
	AppId int `json:"app_id" db:"AppId"`
	Description string `json:"description" db:"Description"`
	ActiveStatus int `json:"active_status" db:"ActiveStatus"`
	CreatorId int `json:"creator_id" db:"CreatorId"`
	CreateDateTime string `json:"create_date_time,omitempty" db:"CreateDateTime"`
	EditorId int `json:"editor_id" db:"EditorId"`
	EditDateTime string `json:"edit_date_time,omitempty" db:"EditDateTime"`
}


func (r *Role) RoleGetAll(db *sqlx.DB) (roles []*Role, err error) {
	sql := `select a.Id,a.RoleCode,a.RoleName,ifnull(b.AppId,'') as AppId,ifnull(a.Description,'') as Description,a.ActiveStatus from Role as a`+
		` left join UserRole as b on a.Id=b.RoleId order by a.Id`
	err = db.Select(&roles,sql)
	if err != nil {
		return nil, err
	}
	return roles, nil

}

func (r *Role) RoleGetByKeyword(db *sqlx.DB, access_token string, keyword string)(roles []*Role, err error){
	sql := `select a.Id,a.RoleCode,a.RoleName,ifnull(b.AppId,'') as AppId,ifnull(a.Description,'') as Description,a.ActiveStatus from Role as a`+
		` left join UserRole as b on a.Id=b.RoleId where a.RoleCode like CONCAT("%",?,"%")  or a.RoleName like CONCAT("%",?,"%") order by a.Id `
	err = db.Select(&roles,sql,keyword,keyword)
	//fmt.Println(sql)
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *Role) RoleGetById(db *sqlx.DB, access_token string, role_id int64) error{
	sql := `select a.Id,a.RoleCode,a.RoleName,ifnull(b.AppId,'') as AppId,ifnull(a.Description,'') as Description,a.ActiveStatus from Role as a`+
		` left join UserRole as b on a.Id=b.RoleId where a.Id = ? order by a.Id limit 1`
	r.Id = role_id
	err := db.Get(r,sql,r.Id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("RoleCode = ",r.RoleCode)
	return nil
}

func (r *Role) RoleSave(db *sqlx.DB) (role_code string, err error){
	r.CreateDateTime = time.Now().String()
	r.ActiveStatus = 1
	sql := `insert into Role(RoleCode,RoleName,AppId,Description,ActiveStatus,CreatorId,CreateDateTime) values(?,?,?,?,?,?,?)`
	res, err := db.Exec(sql,r.RoleCode,r.RoleName,r.AppId,r.Description,r.ActiveStatus,r.CreatorId,r.CreateDateTime)
	if err != nil {
		fmt.Println(err)
		return "",err
	}

	role_code = r.RoleCode
	id, _ := res.LastInsertId()
	fmt.Println("Last Insert Id = ",id)
	return role_code, nil
}


func (r *Role)RoleUpdate(db *sqlx.DB)(role_code string, err error){
	r.EditDateTime = time.Now().String()
	sql := `update Role set RoleCode=?,RoleName=?,AppId=?,Description=?,ActiveStatus=?,EditorId=?,EditDateTime=? where Id=? `
	res, err := db.Exec(sql,r.RoleCode,r.RoleName,r.AppId,r.Description,r.ActiveStatus,r.EditorId,r.EditDateTime,r.Id)

	fmt.Println("sql = ", sql)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	update, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println("Edit Last Id = ",update)

	role_code = r.RoleCode

	return role_code, nil
}

func (r *Role)RoleDisable(db *sqlx.DB)(role_code string, err error){
	r.EditDateTime = time.Now().String()
	sql := `update Role set ActiveStatus=?,EditorId=?,EditDateTime=? where Id=? `
	res, err := db.Exec(sql,r.ActiveStatus,r.EditorId,r.EditDateTime,r.Id)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	update, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println("Edit Last Id = ",update)
	role_code = r.RoleCode
	return role_code, nil
}

package model

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

type Menu struct {
	Id int64 `json:"id" db:"Id"`
	MenuCode string `json:"menu_code" db:"MenuCode"`
	MenuName string `json:"menu_name" db:"MenuName"`
	AppId int64 `json:"app_id" db:"AppId"`
	Description string `json:"description" db:"Description"`
	ActiveStatus int `json:"active_status" db:"ActiveStatus"`
	CreatorId int `json:"creator_id" db:"CreatorId"`
	CreateDateTime string `json:"create_date_time,omitempty" db:"CreateDateTime"`
	EditorId int `json:"editor_id" db:"EditorId"`
	EditDateTime string `json:"edit_date_time,omitempty" db:"EditDateTime"`
}


func (m *Menu) MenuGetAll(db *sqlx.DB) (menus []*Menu, err error) {
	sql := `select Id,MenuCode,MenuName,AppId,ifnull(Description,'') as Description,ActiveStatus from Menu  order by Id `
	err = db.Select(&menus,sql)
	if err != nil {
		return nil, err
	}
	return menus, nil

}

func (m *Menu) MenuGetByKeyword(db *sqlx.DB, access_token string, keyword string)(menus []*Menu, err error){
	sql := `select Id,MenuCode,MenuName,AppId,ifnull(Description,'') as Description,ActiveStatus from Menu  where MenuCode like CONCAT("%",?,"%")  or MenuName like CONCAT("%",?,"%") order by Id `
	err = db.Select(&menus,sql,keyword,keyword)
	if err != nil {
		return nil, err
	}
	return menus, nil
}

func (m *Menu) MenuGetById(db *sqlx.DB, access_token string, menu_id int64) error{
	sql := `select Id,MenuCode,MenuName,AppId,ifnull(Description,'') as Description,ActiveStatus from Menu where Id = ? order by Id limit 1`
	m.Id = menu_id
	err := db.Get(m,sql,m.Id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("MenuCode = ",m.MenuCode)
	return nil

}

func (m *Menu) MenuGetByAppId(db *sqlx.DB, access_token string, app_id int64) (menus []*Menu, err error){
	sql := `select Id,MenuCode,MenuName,AppId,ifnull(Description,'') as Description,ActiveStatus from Menu where AppId = ? order by Id`
	m.AppId = app_id
	err = db.Select(&menus,sql,m.AppId)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("AppId = ",m.AppId)
	return menus, nil

}

//func (m *Menu) MenuGetByAppId(db *sqlx.DB, access_token string, app_id int64) (menus []*Menu, err error){
//	sql := `select Id,MenuCode,MenuName,AppId,ifnull(Description,'') as Description,ActiveStatus from Menu where AppId = ? order by AppId`
//	fmt.Println(sql)
//	err = db.Select(&menus,sql,app_id,app_id)
//	if err != nil {
//		return nil, err
//	}
//	return menus, nil
//}

func (m *Menu) MenuSave(db *sqlx.DB) (menu_code string, err error){
	m.CreateDateTime = time.Now().String()
	m.ActiveStatus = 1
	sql := `insert into Menu(MenuCode,MenuName,AppId,Description,ActiveStatus,CreatorId,CreateDateTime) values(?,?,?,?,?,?,?)`
	res, err := db.Exec(sql,m.MenuCode,m.MenuName,m.AppId,m.Description,m.ActiveStatus,m.CreatorId,m.CreateDateTime)
	if err != nil {
		fmt.Println(err)
		return "",err
	}

	menu_code = m.MenuCode
	id, _ := res.LastInsertId()
	fmt.Println("Last Insert Id = ",id)
	return menu_code, nil
}


func (m *Menu)MenuUpdate(db *sqlx.DB)(menu_code string, err error){
	m.EditDateTime = time.Now().String()
	sql := `update Menu set MenuCode=?,MenuName=?,AppId=?,Description=?,ActiveStatus=?,EditorId=?,EditDateTime=? where Id=? `
	res, err := db.Exec(sql,m.MenuCode,m.MenuName,m.AppId,m.Description,m.ActiveStatus,m.EditorId,m.EditDateTime,m.Id)

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

	menu_code = m.MenuCode

	return menu_code, nil
}


func (m *Menu)MenuDisable(db *sqlx.DB)(menu_code string, err error){
	m.EditDateTime = time.Now().String()
	sql := `update Menu set ActiveStatus=?,EditorId=?,EditDateTime=? where Id=? `
	res, err := db.Exec(sql,m.ActiveStatus,m.EditorId,m.EditDateTime,m.Id)
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

	menu_code = m.MenuCode

	return menu_code, nil
}
package model

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	"time"
)


type App struct {
	Id int64 `json:"id" db:"Id"`
	AppCode string `json:"app_code" db:"AppCode"`
	AppName	string `json:"app_name" db:"AppName"`
	Description string `json:"description" db:"Description"`
	ActiveStatus int `json:"active_status" db:"ActiveStatus"`
	CreatorId int `json:"creator_id" db:"CreatorId"`
	CreateDateTime string `json:"create_date_time,omitempty" db:"CreateDateTime"`
	EditorId int `json:"editor_id" db:"EditorId"`
	EditDateTime string `json:"edit_date_time,omitempty" db:"EditDateTime"`
}


func (a *App) AppGetAll(db *sqlx.DB) (apps []*App, err error) {
	sql := `select Id,AppCode,AppName,ifnull(Description,'') as Description,ActiveStatus from App  order by Id `
	err = db.Select(&apps,sql)
	if err != nil {
		return nil, err
	}
	return apps, nil

}

func (a *App) AppGetByKeyword(db *sqlx.DB, access_token string, keyword string)(apps []*App, err error){
	sql := `select Id,AppCode,AppName,ifnull(Description,'') as Description,ActiveStatus from App  where AppCode like CONCAT("%",?,"%")  or AppName like CONCAT("%",?,"%") order by Id `
	err = db.Select(&apps,sql,keyword,keyword)
	if err != nil {
		return nil, err
	}
	return apps, nil
}

func (a *App) AppGetById(db *sqlx.DB, access_token string, app_id int64) error{
	sql := `select Id,AppCode,AppName,ifnull(Description,'') as Description,ActiveStatus from App where Id = ? order by Id limit 1`
	a.Id = app_id
	err := db.Get(a,sql,a.Id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("AppCode = ",a.AppCode)
	return nil
}

func (a *App) AppSave(db *sqlx.DB) (app_code string, err error){
	a.CreateDateTime = time.Now().String()
	a.ActiveStatus = 1
	sql := `insert into App(AppCode,AppName,Description,ActiveStatus,CreatorId,CreateDateTime) values(?,?,?,?,?,?)`
	res, err := db.Exec(sql,a.AppCode,a.AppName,a.Description,a.ActiveStatus,a.CreatorId,a.CreateDateTime)
	if err != nil {
		fmt.Println(err)
		return "",err
	}

	app_code = a.AppCode
	id, _ := res.LastInsertId()
	fmt.Println("Last Insert Id = ",id)
	return app_code, nil
}


func (a *App)AppUpdate(db *sqlx.DB)(app_code string, err error){
	a.EditDateTime = time.Now().String()
	sql := `update App set AppCode=?,AppName=?,Description=?,ActiveStatus=?,EditorId=?,EditDateTime=? where Id=? `
	res, err := db.Exec(sql,a.AppCode,a.AppName,a.Description,a.ActiveStatus,a.EditorId,a.EditDateTime,a.Id)

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

	app_code = a.AppCode

	return app_code, nil
}


func (a *App)AppDisable(db *sqlx.DB)(app_code string, err error){
	a.EditDateTime = time.Now().String()
	sql := `update App set ActiveStatus=?,EditorId=?,EditDateTime=? where Id=? `
	res, err := db.Exec(sql,a.ActiveStatus,a.EditorId,a.EditDateTime,a.Id)
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

	app_code = a.AppCode

	return app_code, nil
}








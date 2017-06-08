package model

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	"time"
)

type User struct {
	Id int64 `json:"id" db:"Id"`
	UserCode string `json:"user_code" db:"UserCode"`
	UserName string `json:"user_name" db:"UserName"`
	Password string `json:"password" db:"Password"`
	Telephone string `json:"telephone" db:"Telephone,omitempty"`
	SaleId int `json:"sale_id" db:"SaleId"`
	ProfitId int `json:"profit_id" db:"ProfitId"`
	DepartmentId int `json:"department_id" db:"DepartmentId"`
	ExpertId int `json:"expert_id" db:"ExpertId"`
	ActiveStatus int `json:"active_status" db:"ActiveStatus"`
	CreatorId int `json:"creator_id" db:"CreatorId"`
	CreateDateTime string `json:"create_date_time,omitempty" db:"CreateDateTime"`
	EditorId int `json:"editor_id" db:"EditorId"`
	EditDateTime string `json:"edit_date_time,omitempty" db:"EditDateTime"`
}


func (u *User)UserGetById(db *sqlx.DB, access_token string, user_id int64) error{
	sql := `select Id,UserCode,UserName,Password,Telephone,ProfitId,DepartmentId,ExpertId from User where id =? limit 1`
	u.Id = user_id
	err := db.Get(u,sql,u.Id)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func (u *User)UserGetByKeyword(db *sqlx.DB,access_token string, keyword string) (users []*User,err error){
	sql := `select Id,UserCode,UserName,Password,Telephone,ProfitId,DepartmentId,ExpertId from User where UserCode like CONCAT("%",?,"%") or UserName like CONCAT("%",?,"%") order by Id`
	err = db.Select(&users,sql,keyword,keyword)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return users, nil
}


func (u *User)UserGetAll(db *sqlx.DB, access_token string) (users []*User,err error){
	sql := `select Id,UserCode,UserName,Password,Telephone,ProfitId,DepartmentId,ExpertId from User order by usercode`
	err = db.Select(&users,sql)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return users, nil
}


type dateStruct struct {
	time.Time
}


func (u *User)UserSave(db *sqlx.DB) (user_code string, err error){
	fmt.Println("UserCode = ",u.UserCode)
	err = u.GetUserNotExist(db)
	fmt.Println("check nil =",err)
	if err == nil {
		fmt.Println(err)
		return "", err
	}

	u.CreateDateTime = time.Now().String()

	fmt.Println("Date = ",u.CreateDateTime)

	sql := `Insert into User(UserCode,UserName,Password,Telephone,ProfitId,DepartmentId,ExpertId,CreatorId,CreateDateTime) Values(?,?,?,?,?,?,?,?,?)`
	res, err := db.Exec(sql,
		u.UserCode,
		u.UserName,
		u.Password,
		u.Telephone,
		u.ProfitId,
		u.DepartmentId,
		u.ExpertId,
		u.CreatorId,
		u.CreateDateTime)
	if err != nil {
		fmt.Println("Error = ",err.Error())
		return "", err
	}

	id, err := res.LastInsertId()
	u.Id = id
	fmt.Println("Last Insert Id = ",id)

	return u.UserCode, nil
}


func (u *User)UserUpdate(db *sqlx.DB)(user_code string, err error){
	fmt.Println("UserCode = ",u.Id)
	err = u.GetUserNotExist(db)
	fmt.Println("check nil =",err)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	u.EditDateTime = time.Now().String()
	sql := `update User set UserCode=?,UserName=?,Password=?,Telephone=?,ProfitId=?,DepartmentId=?,ExpertId=?,ActiveStatus=?,EditorId=?,EditDateTime=? where id = ?`
	res, err := db.Exec(sql,u.UserCode,u.UserName,u.Password,u.Telephone,u.ProfitId,u.DepartmentId,u.ExpertId,u.ActiveStatus,u.EditorId,u.EditDateTime,u.Id)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	update, err :=res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println("",update)

	return u.UserCode, nil
}


func (u *User) GetUserNotExist(db *sqlx.DB) error {
	sql := `select Id from User where Id = ?`
	err := db.Get(u, sql, u.Id)
	fmt.Println("Get User =",u.Id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
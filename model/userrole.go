package model

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	//"time"
	//"github.com/revel/modules/db/app"
	"time"
)

type UserRole struct {
	Id       		int64 `json:"id" db:"Id"`
	AppId       	int64 `json:"app_id" db:"AppId"`
	AppCode 		string `json:"app_code" db:"AppCode"`
	AppName 		string `json:"app_name" db:"AppName"`
	UserId      	int64 `json:"user_id" db:"UserId"`
	UserCode 		string `json:"user_code" db:"UserCode"`
	UserName 		string `json:"user_name" db:"UserName"`
	Salecode		string `json:"sale_code" db:"Salecode"`
	RoleId       	int64 `json:"role_id" db:"RoleId"`
	RoleCode 		string `json:"role_code" db:"RoleCode"`
	RoleName 		string `json:"role_name" db:"RoleName"`
	CreatorId 		int `json:"creator_id" db:"CreatorId"`
	CreateDateTime 	string `json:"create_date_time,omitempty" db:"CreateDateTime"`
	EditorId 		int `json:"editor_id" db:"EditorId"`
	EditDateTime 	string `json:"edit_date_time,omitempty" db:"EditDateTime"`
}

func (ur *UserRole)UserRoleGetAll(db *sqlx.DB, access_token string, user_id int64) (userroles []*UserRole,err error){
	sql := `select a.Id,a.AppId,b.AppCode,b.AppName,a.UserId,c.Salecode,c.UserCode,c.UserName`+
		` ,a.RoleId,d.RoleCode,d.RoleName,a.CreatorId,a.CreateDateTime,a.EditorId,a.EditDateTime from UserRole as a`+
		` left join App as b on a.AppId=b.Id`+
		` left join User as c on a.UserId=c.Id`+
		` left join Role as d on a.RoleId=d.Id where a.UserId=?`
	ur.UserId=user_id
	err = db.Select(&userroles,sql,ur.UserId)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return userroles, nil
}

func (ur *UserRole) UserRoleGetUser (db *sqlx.DB, access_token string, app_id int64,user_id int64)  error {
	sql := `select a.Id,a.AppId,b.AppCode,b.AppName,a.UserId,c.UserCode,c.UserName`+
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

func (ur *UserRole) UserRoleSave(db *sqlx.DB) (user_id int64, err error){
	fmt.Println("UserRole = ",ur.Id)
	err = ur.GetUserRoleNotExist(db)

	ur.CreateDateTime = time.Now().String()
	//ur.CreatorId = 1
	sql := `insert into UserRole(AppId,UserId,RoleId,CreatorId,CreateDateTime,EditorId,EditDateTime) values(?,?,?,?,?,?,?)`
	fmt.Println("sql = ",sql)
	res, err := db.Exec(sql,ur.AppId,ur.UserId,ur.RoleId,ur.CreatorId,ur.CreateDateTime,ur.CreatorId,ur.CreateDateTime)
	if err != nil {
		fmt.Println(err)
		return 0,err
	}
	user_id = ur.UserId
	id, _ := res.LastInsertId()
	fmt.Println("Last Insert Id = ",id)
	return id, nil
}

func (ur *UserRole) UserRoleUpdate(db *sqlx.DB)(user_id int64, err error){
	fmt.Println("UserRole = ",ur.Id)
	err = ur.GetUserRoleNotExist(db)

	ur.EditDateTime = time.Now().String()
	//ur.EditorId = 3
	sql := `update UserRole set RoleId=?,EditorId=?,EditDateTime=? where id = ?`
	res, err := db.Exec(sql,ur.RoleId,ur.EditorId,ur.EditDateTime,ur.Id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	user_id = ur.Id
	update, err :=res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	fmt.Println("status",update)
	return update, nil
}

func (ur *UserRole) GetUserRoleNotExist(db *sqlx.DB) error {
	sql := `select Id from UserRole where Id = ?`
	err := db.Get(ur, sql, ur.Id)
	fmt.Println("Get UserRole =",ur.Id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
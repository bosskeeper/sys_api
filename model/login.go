package model

import (
	"github.com/jmoiron/sqlx"
	"log"
	"fmt"
	"io"
	"crypto/rand"
)

type Login struct {
	Id               int64       `json:"id" db:"Id"`
	UserCode         string      `json:"usercode" db:"UserCode"`
	UserName         string      `json:"username" db:"UserName"`
	Password         string      `json:"password" db:"Password"`
	CompanyId        int         `json:"company_id" db:"CompanyId"`
	CompanyName      string      `json:"company_name" db:"CompanyName"`
	BranchId         int64       `json:"branch_id" db:"BranchId"`
	BranchName       string      `json:"branch_name" db:"BranchName"`
	PicPath          string      `json:"pic_path" db:"PicPath"`
	UserActiveStatus int64       `json:"usercctivestatus" db:"UserActiveStatus"`
	RoleId           int64       `json:"roleid" db:"RoleId"`
	RoleCode         string      `json:"rolecode" db:"RoleCode"`
	RoleName         string      `json:"rolename" db:"RoleName"`
	AppID            int64       `json:"appid" db:"AppID"`
	AppCode          string      `json:"appcode" db:"AppCode"`
	AppName          string      `json:"appname" db:"AppName"`
	SaleCode         string      `json:"sale_code" db:"SaleCode"`
	AccessToken      string      `json:"access_token" db:"AccessToken"`
	Menus            []*LoginSub `json:"menu"`
}

type LoginSub struct {
	MenuID       int64  `json:"menuid" db:"MenuID"`
	MenuCode     string `json:"menucode" db:"MenuCode"`
	MenuName     string `json:"menuname" db:"MenuName"`
	PermissionID int64  `json:"permissionid" db:"PermissionID"`
	IsCreate     int64  `json:"is_create" db:"IsCreate"`
	IsUpdate     int64  `json:"is_update" db:"IsUpdate"`
	IsRead       int64  `json:"is_read" db:"IsRead"`
	IsDelete     int64  `json:"is_delete" db:"IsDelete"`
}

var dbl *sqlx.DB

func ConnectDB(dbName string) (db *sqlx.DB, err error) {
	dsn := "root:[ibdkifu88@tcp(nopadol.net:3306)/" + dbName + "?parseTime=true&charset=utf8&loc=Local"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, err
	}

	return db, err
}

func (l *Login) LoginGetByUser(db *sqlx.DB, access_token string, user_code string, password string, appid int64) (err error) {
	dbl, err = ConnectDB("sys")
	if err != nil {
		fmt.Println(err.Error())
	}
	sql := `select a.Id,ifnull(a.UserCode,'') as UserCode,ifnull(a.UserName,'') as UserName,ifnull(a.Password,'') as Password,ifnull(a.SaleCode,'') as SaleCode,a.ActiveStatus as UserActiveStatus,c.id as RoleId,ifnull(c.RoleCode,'') as RoleCode,ifnull(c.RoleName,'') as RoleName,b.AppID,ifnull(d.AppCode,'') as AppCode,ifnull(d.AppName,'') as AppName,a.BranchId,ifnull(f.BranchName,'') as BranchName,ifnull(a.PicPath,'') as PicPath, ifnull(a.CompanyId,0) as CompanyId, ifnull(e.CompanyName,'') as CompanyName ` +
		` from User as a` +
		` left join UserRole as b on a.Id=b.UserId` +
		` left join Role as c on b.RoleId=c.Id` +
		` left join App as d on b.AppId=d.Id` +
		` left join CompanyMaster as e on a.CompanyId=e.Id` +
		` left join BranchMaster as f on a.BranchId=f.Id` +
		` where a.UserCode=? and a.Password=? and b.AppID=? limit 1`
	l.UserCode = user_code
	l.Password = password
	l.AppID = appid
	err = dbl.Get(l, sql, user_code, password, appid)
	log.Println(sql)
	fmt.Println("sql = ", sql)
	if err != nil {
		log.Println("Error ", err.Error())
	}

	uuid, err := newUUID()
	if err != nil {
		fmt.Println(err.Error())
	}

	l.AccessToken = uuid

	sqlsub := `select f.Id as MenuID,f.MenuCode,f.MenuName,ifnull(g.Id,0) as PermissionID,ifnull(g.IsCreate,0) as IsCreate` +
		` ,ifnull(g.IsUpdate,0) as IsUpdate,ifnull(g.IsRead,0) as IsRead,ifnull(g.IsDelete,0) as IsDelete` +
		` from User as a` +
		` left join UserRole as b on a.Id=b.UserId` +
		` left join Role as c on b.RoleId=c.Id` +
		` left join App as d on b.AppId=d.Id` +
		` left join Menu as f on d.Id=f.AppId` +
		` left join Permission as g on c.Id=g.RoleId and d.Id=g.AppId and f.Id=g.MenuId` +
		` where a.UserCode=? and a.Password=? and b.AppID=? and f.activestatus=1`
	//sqlsub = "select f.Id as MenuID,f.MenuCode,f.MenuName,g.Id as PermissionID,g.Create,g.Update,g.Read,g.Delete" +
	//	" from User as a" +
	//	" left join UserRole as b on a.Id=b.UserId " +
	//	" left join Role as c on b.RoleId=c.Id" +
	//	" left join App as d on b.RoleId=d.Id " +
	//	" left join Menu as f on d.Id=f.AppId" +
	//	" left join Permission as g on c.Id=g.RoleId and d.Id=g.AppId and f.Id=g.MenuId" +
	//	"  where a.UserName='"+l.UserName+"' and a.Password='"+l.Password+"' and b.AppID="+ strconv.FormatInt(l.AppID, 10)

	fmt.Println("sqlsub = ", sqlsub)
	err = dbl.Select(&l.Menus, sqlsub, l.UserCode, l.Password, l.AppID)
	//err = db.Select(&l.menus,sqlsub)
	fmt.Println("Menus = ", l.UserCode, l.Password, l.AppID)
	if err != nil {
		log.Println("Error ", err.Error())
	}
	fmt.Println(l)

	sql_log := `Insert into UserAccessLogs(UserCode,AccessToken,AppId,BranchId,LoginTime) values(?,?,?,?,CURRENT_TIMESTAMP())`
	fmt.Println("sql_log =", sql_log)
	resp, err := dbl.Exec(sql_log, l.UserCode, l.AccessToken, l.AppID, l.BranchId)
	if err != nil {
		log.Println("Error ", err.Error())
	}
	id, _ := resp.LastInsertId()
	l.Id = id

	return nil
}

// newUUID generates a random UUID according to RFC 4122
func newUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}

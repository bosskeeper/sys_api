package model

import (
	"github.com/jmoiron/sqlx"
	//"fmt"
	//"github.com/matryer/m"
	"log"
	"fmt"
	//"strconv"
	//	"golang.org/x/net/icmp"
)

type Login struct {
	Id               int64       `json:"id" db:"Id"`
	UserCode         string      `json:"usercode" db:"UserCode"`
	UserName         string      `json:"username" db:"UserName"`
	Password         string      `json:"password" db:"Password"`
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
	sql := `select a.Id,ifnull(a.UserCode,'') as UserCode,ifnull(a.UserName,'') as UserName,ifnull(a.Password,'') as Password,ifnull(a.SaleCode,'') as SaleCode,a.ActiveStatus as UserActiveStatus,c.id as RoleId,ifnull(c.RoleCode,'') as RoleCode,ifnull(c.RoleName,'') as RoleName,b.AppID,ifnull(d.AppCode,'') as AppCode,ifnull(d.AppName,'') as AppName,case when a.BranchId = 1 then 'สำนักงานใหญ่' when a.BranchId = 2 then 'สาขาสันกำแพง' end BranchName,ifnull(a.PicPath,'') as PicPath ` +
		` from User as a` +
		` left join UserRole as b on a.Id=b.UserId` +
		` left join Role as c on b.RoleId=c.Id` +
		` left join App as d on b.AppId=d.Id` +
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

	return nil
}

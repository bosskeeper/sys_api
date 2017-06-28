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
	Id 	 int64 `json:"id" db:"Id"`
	UserCode string `json:"usercode" db:"UserCode"`
	UserName string `json:"username" db:"UserName"`
	Password string `json:"password" db:"Password"`
	RoleCode string `json:"rolecode" db:"RoleCode"`
	RoleName string `json:"rolename" db:"RoleName"`
	AppID  int64 `json:"appid" db:"AppID"`
	AppCode  string `json:"appcode" db:"AppCode"`
	AppName  string `json:"appname" db:"AppName"`
	Menus     []*LoginSub `json:"menu"`
}

type LoginSub struct {
	MenuID int64 `json:"menuid" db:"MenuID"`
	MenuCode string `json:"menucode" db:"MenuCode"`
	MenuName string `json:"menuname" db:"MenuName"`
	PermissionID int64 `json:"permissionid" db:"PermissionID"`
	Create int64 `json:"create" db:"Create"`
	Update  int64 `json:"update" db:"Update"`
	Read  int64 `json:"read" db:"Read"`
	Delete  int64 `json:"delete" db:"Delete"`
}


func (l *Login) LoginGetByUser(db *sqlx.DB, access_token string,user_name string,password string,appid int64) (err error) {
	sql := `select a.Id,a.UserCode,a.UserName,a.Password,c.RoleCode,c.RoleName,b.AppID,d.AppCode,d.AppName`+
		` from User as a`+
		` left join UserRole as b on a.Id=b.UserId`+
		` left join Role as c on b.RoleId=c.Id`+
		` left join App as d on b.RoleId=d.Id`+
		` where a.UserName=? and a.Password=? and b.AppID=? limit 1`
	l.UserName = user_name
	l.Password = password
	l.AppID = appid
		err = db.Get(l,sql,l.UserName,l.Password,l.AppID)
	log.Println("Error ",sql)
	//fmt.Println("UserID = ",l.Id)
	if err != nil {
		log.Println("Error ", err.Error())
	}


	sqlsub := `select f.Id as MenuID,f.MenuCode,f.MenuName,g.Id as PermissionID,g.Create,g.Update,g.Read,g.Delete` +
		` from User as a` +
		` left join UserRole as b on a.Id=b.UserId` +
		` left join Role as c on b.RoleId=c.Id` +
		` left join App as d on b.RoleId=d.Id` +
		` left join Menu as f on d.Id=f.AppId` +
		` left join Permission as g on c.Id=g.RoleId and d.Id=g.AppId and f.Id=g.MenuId` +
		` where a.UserName=? and a.Password=? and b.AppID=?`
	//sqlsub = "select f.Id as MenuID,f.MenuCode,f.MenuName,g.Id as PermissionID,g.Create,g.Update,g.Read,g.Delete" +
	//	" from User as a" +
	//	" left join UserRole as b on a.Id=b.UserId " +
	//	" left join Role as c on b.RoleId=c.Id" +
	//	" left join App as d on b.RoleId=d.Id " +
	//	" left join Menu as f on d.Id=f.AppId" +
	//	" left join Permission as g on c.Id=g.RoleId and d.Id=g.AppId and f.Id=g.MenuId" +
	//	"  where a.UserName='"+l.UserName+"' and a.Password='"+l.Password+"' and b.AppID="+ strconv.FormatInt(l.AppID, 10)

	fmt.Println(sqlsub)



		err = db.Select(&l.Menus,sqlsub,l.UserName,l.Password,l.AppID)
		//err = db.Select(&l.menus,sqlsub)

	fmt.Println("Menus = ", l.UserName,l.Password,l.AppID)
		if err != nil {
			log.Println("Error ", err.Error())
		}
	fmt.Println(l)

	return nil
}


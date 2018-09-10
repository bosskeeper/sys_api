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
	SaleCode string `json:"sale_code" db:"SaleCode"`
	BranchId int `json:"branch_id" db:"BranchId"`
	BranchCode string `json:"branch_code,omitempty" db:"BranchCode"`
	BranchName string `json:"branch_name,omitempty" db:"BranchName"`
	ProfitcenterId int `json:"profitcenter_id" db:"ProfitcenterId"`
	ProfitcenterCode string `json:"profitcenter_code,omitempty" db:"ProfitcenterCode"`
	ProfitcenterName string `json:"profitcenter_name,omitempty" db:"ProfitcenterName"`
	DepartmentId int `json:"department_id" db:"DepartmentId"`
	DepartmentCode string `json:"department_code,omitempty" db:"DepartmentCode"`
	DepartmentName string `json:"department_name,omitempty" db:"DepartmentName"`
	ExpertId int `json:"expert_id" db:"ExpertId"`
	ExpertCode string `json:"expert_code,omitempty" db:"ExpertCode"`
	ExpertName string `json:"expert_name,omitempty" db:"ExpertName"`
	ActiveStatus int `json:"active_status" db:"ActiveStatus"`
	CreatorId int `json:"creator_id" db:"CreatorId"`
	CreateDateTime string `json:"create_date_time,omitempty" db:"CreateDateTime"`
	EditorId int `json:"editor_id" db:"EditorId"`
	EditDateTime string `json:"edit_date_time,omitempty" db:"EditDateTime"`
}


func (u *User)UserGetById(db *sqlx.DB, access_token string, user_id int64) error{
	sql := `select a.Id,a.UserCode,a.UserName,a.Password,a.Telephone,a.SaleId,a.SaleCode`+
			` ,a.BranchId,d.BranchCode,d.BranchName,a.ProfitcenterId,b.ProfitcenterCode,b.ProfitcenterName`+
			` ,a.DepartmentId,c.DepartmentCode,c.DepartmentName,a.ExpertId,e.ExpertCode,e.ExpertName,a.ActiveStatus`+
			` from User as a left join ProfitcenterMaster as b on a.ProfitcenterId=b.Id`+
			` left join DepartmentMaster as c on a.DepartmentId=c.Id`+
			` left join BranchMaster as d on a.BranchId=d.Id`+
			` left join ExpertMaster as e on a.ExpertId=e.Id where a.id =? order by a.id limit 1`
	u.Id = user_id
	err := db.Get(u,sql,u.Id)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func (u *User)UserGetByUserCode(db *sqlx.DB, access_token string, user_code string) error{
	sql := `select a.Id,a.UserCode,a.UserName,a.Password,a.Telephone,a.SaleId,a.SaleCode`+
		` ,a.BranchId,d.BranchCode,d.BranchName,a.ProfitcenterId,b.ProfitcenterCode,b.ProfitcenterName`+
		` ,a.DepartmentId,c.DepartmentCode,c.DepartmentName,a.ExpertId,e.ExpertCode,e.ExpertName,a.ActiveStatus`+
		` from User as a left join ProfitcenterMaster as b on a.ProfitcenterId=b.Id`+
		` left join DepartmentMaster as c on a.DepartmentId=c.Id`+
		` left join BranchMaster as d on a.BranchId=d.Id`+
		` left join ExpertMaster as e on a.ExpertId=e.Id where a.UserCode =? limit 1`
	fmt.Println(sql)
	u.UserCode = user_code
	err := db.Get(u,sql,u.UserCode)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (u *User)UserGetByKeyword(db *sqlx.DB,access_token string, keyword string) (users []*User,err error){
	sql := `select a.Id,a.UserCode,a.UserName,a.Password,a.Telephone,a.SaleId,a.SaleCode`+
		` ,a.BranchId,d.BranchCode,d.BranchName,a.ProfitcenterId,b.ProfitcenterCode,b.ProfitcenterName`+
		` ,a.DepartmentId,c.DepartmentCode,c.DepartmentName,a.ExpertId,e.ExpertCode,e.ExpertName,a.ActiveStatus`+
		` from User as a left join ProfitcenterMaster as b on a.ProfitcenterId=b.Id`+
		` left join DepartmentMaster as c on a.DepartmentId=c.Id`+
		` left join BranchMaster as d on a.BranchId=d.Id`+
		` left join ExpertMaster as e on a.ExpertId=e.Id where a.UserCode like CONCAT("%",?,"%") or a.UserName like CONCAT("%",?,"%") order by a.Id asc`
	err = db.Select(&users,sql,keyword,keyword)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return users, nil
}


func (u *User)UserGetAll(db *sqlx.DB, access_token string) (users []*User,err error){
	sql := `select a.Id,a.UserCode,a.UserName,a.Password,a.Telephone,a.SaleId,a.SaleCode`+
			` ,a.BranchId,d.BranchCode,d.BranchName,a.ProfitcenterId,b.ProfitcenterCode,b.ProfitcenterName`+
			` ,a.DepartmentId,c.DepartmentCode,c.DepartmentName,a.ExpertId,e.ExpertCode,e.ExpertName,a.ActiveStatus`+
			` from User as a left join ProfitcenterMaster as b on a.ProfitcenterId=b.Id`+
			` left join DepartmentMaster as c on a.DepartmentId=c.Id`+
			` left join BranchMaster as d on a.BranchId=d.Id`+
			` left join ExpertMaster as e on a.ExpertId=e.Id order by a.Id asc`
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

	sql := `Insert into User(UserCode,UserName,Password,Telephone,SaleId,SaleCode,BranchId,ProfitcenterId,DepartmentId,ExpertId,CreatorId,CreateDateTime) Values(?,?,?,?,?,?,?,?,?,?,?,?)`
	res, err := db.Exec(sql,
		u.UserCode,
		u.UserName,
		u.Password,
		u.Telephone,
		u.SaleId,
		u.SaleCode,
		u.BranchId,
		u.ProfitcenterId,
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
	sql := `update User set UserCode=?,UserName=?,Password=?,Telephone=?,BranchId=?,ProfitcenterId=?,DepartmentId=?,ExpertId=?,ActiveStatus=?,EditorId=?,EditDateTime=? where id = ?`
	res, err := db.Exec(sql,u.UserCode,u.UserName,u.Password,u.Telephone,u.ProfitcenterId,u.DepartmentId,u.ExpertId,u.ActiveStatus,u.EditorId,u.EditDateTime,u.Id)
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


func (u *User)UserDisable(db *sqlx.DB)(user_code string, err error){
	fmt.Println("UserCode = ",u.Id)
	err = u.GetUserNotExist(db)
	fmt.Println("check nil =",err)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	u.EditDateTime = time.Now().String()
	sql := `update User set UserCode=?,ActiveStatus=?,EditorId=?,EditDateTime=? where id = ?`
	res, err := db.Exec(sql,u.UserCode,u.ActiveStatus,u.EditorId,u.EditDateTime,u.Id)
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


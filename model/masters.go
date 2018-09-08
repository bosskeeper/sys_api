package model

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	//"time"
)

type BranchMaster struct {
	Id int64 `json:"id" db:"Id"`
	BranchCode string `json:"branch_code,omitempty" db:"BranchCode"`
	BranchName string `json:"branch_name,omitempty" db:"BranchName"`
	ActiveStatus int `json:"active_status" db:"ActiveStatus"`
	CreatorId int `json:"creator_id" db:"CreatorId"`
	CreateDateTime string `json:"create_date_time,omitempty" db:"CreateDateTime"`
	//EditorId int `json:"editor_id" db:"EditorId"`
	//EditDateTime string `json:"edit_date_time,omitempty" db:"EditDateTime"`
}

type DepartmentMaster struct {
	Id int64 `json:"id" db:"Id"`
	DepartmentCode string `json:"department_code,omitempty" db:"DepartmentCode"`
	DepartmentName string `json:"department_name,omitempty" db:"DepartmentName"`
	ActiveStatus int `json:"active_status" db:"ActiveStatus"`
	CreatorId int `json:"creator_id" db:"CreatorId"`
	CreateDateTime string `json:"create_date_time,omitempty" db:"CreateDateTime"`
	//EditorId int `json:"editor_id" db:"EditorId"`
	//EditDateTime string `json:"edit_date_time,omitempty" db:"EditDateTime"`
}

type ExpertMaster struct {
	Id int64 `json:"id" db:"Id"`
	ExpertCode string `json:"expert_code,omitempty" db:"ExpertCode"`
	ExpertName string `json:"expert_name,omitempty" db:"ExpertName"`
	ActiveStatus int `json:"active_status" db:"ActiveStatus"`
	CreatorId int `json:"creator_id" db:"CreatorId"`
	CreateDateTime string `json:"create_date_time,omitempty" db:"CreateDateTime"`
	//EditorId int `json:"editor_id" db:"EditorId"`
	//EditDateTime string `json:"edit_date_time,omitempty" db:"EditDateTime"`
}

type ProfitcenterMaster struct {
	Id int64 `json:"id" db:"Id"`
	ProfitcenterCode string `json:"profitcenter_code,omitempty" db:"ProfitcenterCode"`
	ProfitcenterName string `json:"profitcenter_name,omitempty" db:"ProfitcenterName"`
	TeamStatus int `json:"team_status" db:"TeamStatus"`
	ActiveStatus int `json:"active_status" db:"ActiveStatus"`
	CreatorId int `json:"creator_id" db:"CreatorId"`
	CreateDateTime string `json:"create_date_time,omitempty" db:"CreateDateTime"`
	//EditorId int `json:"editor_id" db:"EditorId"`
	//EditDateTime string `json:"edit_date_time,omitempty" db:"EditDateTime"`
}

func (b *BranchMaster)GetBranchs(db *sqlx.DB, access_token string) (branchs []*BranchMaster,err error){

	sql := `SELECT Id, BranchCode, BranchName, ActiveStatus, CreatorId, CreateDateTime
			 FROM BranchMaster order by Id asc`
	err = db.Select(&branchs,sql)
	//fmt.Println("sql : ",sql)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return branchs, nil
}

func (d *DepartmentMaster)GetDepartments(db *sqlx.DB, access_token string) (departments []*DepartmentMaster,err error){

	sql := `SELECT Id, DepartmentCode, DepartmentName, ActiveStatus, CreatorId, CreateDateTime
			FROM DepartmentMaster order by Id asc`
	err = db.Select(&departments,sql)
	//fmt.Println("sql : ",sql)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return departments, nil
}

func (e *ExpertMaster)GetExperts(db *sqlx.DB, access_token string) (experts []*ExpertMaster,err error){

	sql := `SELECT Id, ExpertCode, ExpertName, ActiveStatus, CreatorId, CreateDateTime
			FROM ExpertMaster order by Id asc`
	err = db.Select(&experts,sql)
	//fmt.Println("sql : ",sql)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return experts, nil
}

func (p *ProfitcenterMaster)GetProfitcenters(db *sqlx.DB, access_token string) (profitcenters []*ProfitcenterMaster,err error){

	sql := `SELECT Id, ProfitcenterCode, ProfitcenterName, TeamStatus, ActiveStatus, CreatorId, CreateDateTime
			FROM ProfitcenterMaster order by Id asc`
	err = db.Select(&profitcenters,sql)
	//fmt.Println("sql : ",sql)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return profitcenters, nil
}


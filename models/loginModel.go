package models

import (
	"database/sql"
	"userManageSystem/utils"
)

type Login struct {
	VisitTime string `json:"visit_time"`
}

func InsertTime(time string) {
	db := utils.InitDB()
	stmt, err := db.Prepare("insert into login(visit_time) values (?)")
	defer stmt.Close()
	defer db.Close()
	if err != nil {
		utils.ERROR.Println(err)
	}
	stmt.Query(time)
	utils.INFO.Printf("insert into login(visit_time) values ('%s')", time)
}

func QueryTime(time string) *sql.Rows {
	db := utils.InitDB()
	stmt, err := db.Prepare("select count(*) from login where date(visit_time) = ?")
	defer stmt.Close()
	defer db.Close()
	if err != nil {
		utils.ERROR.Println(err)
	}
	res, _ := stmt.Query(time)
	utils.INFO.Printf("select count(*) from login where date(visit_time) = '%s'", time)
	return res
}

func QueryTimeByMonth(month int) *sql.Rows {
	db := utils.InitDB()
	stmt, err := db.Prepare("select count(*) from login where month(visit_time) = ?")
	defer stmt.Close()
	defer db.Close()
	if err != nil {
		utils.ERROR.Println(err)
	}
	res, _ := stmt.Query(month)
	utils.INFO.Printf("select count(*) from login where month(visit_time) = %d", month)
	return res
}

package service

import (
	"math"
	"time"
	"userManageSystem/models"
	"userManageSystem/utils"
)

type UserStatus struct {
	Title  string  `json:"title"`
	Value  int     `json:"value"`
	Change float64 `json:"change"`
}

func DataAnalysis() *utils.Message {
	var lastLogin, lastRegister, lastDelete int
	login := UserStatus{Title: "登录用户"}
	register := UserStatus{Title: "注册用户"}
	status := UserStatus{Title: "注销用户"}
	month := time.Now().Month()

	res := models.GetUserRegisterNumByMonth(int(month))
	for res.Next() {
		res.Scan(&register.Value)
	}

	res = models.QueryTimeByMonth(int(month))
	for res.Next() {
		res.Scan(&login.Value)
	}

	res = models.GetUserDeleteNumByMonth(int(month))
	for res.Next() {
		res.Scan(&status.Value)
	}

	var lastMonth time.Month
	if month == time.January {
		lastMonth = time.December
	} else {
		lastMonth = month - 1
	}

	res = models.QueryTimeByMonth(int(lastMonth))
	for res.Next() {
		res.Scan(&lastLogin)

	}
	res = models.GetUserRegisterNumByMonth(int(lastMonth))
	for res.Next() {
		res.Scan(&lastRegister)

	}
	res = models.GetUserDeleteNumByMonth(int(lastMonth))
	for res.Next() {
		res.Scan(&lastDelete)
	}

	if lastRegister == 0 {
		register.Change = float64(register.Value - lastRegister)
	} else {
		register.Change = float64(register.Value-lastRegister) / float64(lastRegister)
		register.Change = math.Round(register.Change*100) / 100
	}
	if lastLogin == 0 {
		login.Change = float64(login.Value - lastLogin)
	} else {
		login.Change = float64(login.Value-lastLogin) / float64(lastLogin)
		login.Change = math.Round(login.Change*100) / 100
	}
	if lastDelete == 0 {
		status.Change = float64(status.Value - lastDelete)
	} else {
		status.Change = float64(status.Value-lastDelete) / float64(lastDelete)
		status.Change = math.Round(status.Change*100) / 100
	}
	data := []UserStatus{
		register,
		login,
		status,
	}
	return utils.Ok(data, "获取成功")
}

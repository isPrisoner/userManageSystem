package service

import (
	"log"
	"net/http"
	"os"
	"time"
	"userManageSystem/models"
	"userManageSystem/utils"
)

func Login(username, password string, w http.ResponseWriter) *utils.Message {
	us := models.User{Username: username, Password: password}
	res := models.GetUser(&us)
	status := 0
	for res.Next() {
		res.Scan(&status)
	}
	if status == 0 {
		return utils.Lose(nil, "用户登录失败")
	}
	models.UpdateTime(us.Username)
	res = models.GetUserSessionByUsername(username)
	var sessionID string
	for res.Next() {
		res.Scan(&sessionID)
	}
	if sessionID == "" {
		utils.SessionInit(w, username)
	} else {
		utils.Session(w, sessionID)
	}
	models.InsertTime(time.Now().Format("2006-01-02 15:04:05"))
	return utils.Ok(nil, "用户登录成功")
}

func ValidationRole(sessionID string) string {
	res := models.GetUserByRole(sessionID)
	var roleName string
	for res.Next() {
		res.Scan(&roleName)
	}
	return roleName
}

func QueryUser(username string) *utils.Message {
	res := models.GetUserByName(username)
	var user models.User
	for res.Next() {
		res.Scan(&user.Image, &user.ID, &user.Username, &user.Password, &user.Email, &user.Status)
	}
	user.Image = "/images/" + user.Image
	return utils.Ok(user, "获取到用户信息")
}

func DeleteUser(username string) *utils.Message {
	res := models.GetImageByName(username)
	count := models.DeleteUser(username)
	if count == 0 {
		return utils.Lose(nil, "删除失败")
	}
	var url string
	for res.Next() {
		res.Scan(&url)
	}
	if url != "default.png" {
		if err := os.Remove("images/" + url); err != nil {
			utils.ERROR.Println("删除文件失败", err)
		}
	}
	return utils.Ok(nil, "删除成功")
}

func Image(session string) string {
	res := models.GetImageBySessionID(session)
	var url string
	for res.Next() {
		res.Scan(&url)
	}
	url = "/images/" + url
	return url
}

func Password(username, email string) *utils.Message {
	res := models.GetPasswordByEmail(username, email)
	var password string
	for res.Next() {
		res.Scan(&password)
	}
	if password == "" {
		return utils.Lose(nil, "密码获取失败，请确认邮箱和用户名")
	}
	log.Println(password)
	return utils.Ok(nil, "密码已发送至邮箱请查看")
}

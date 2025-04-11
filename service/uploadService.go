package service

import (
	"io"
	"math/rand"
	"mime/multipart"
	"os"
	"path"
	"strconv"
	"userManageSystem/models"
	"userManageSystem/utils"
)

func upload(username, s, id, password, email, currentAvatar string, file multipart.File, header *multipart.FileHeader) *models.User {
	// status 只可能是0或1
	status, stErr := strconv.Atoi(s)
	if id == "" {
		var lastID int
		res := models.GetMaxUserID()
		for res.Next() {
			res.Scan(&lastID)
		}
		id = strconv.Itoa(lastID + 1)
	}
	var url string
	if file == nil {
		// 如果没有上传新图片，使用当前图片路径
		url = path.Base(currentAvatar)
		if url == "." {
			url = strconv.Itoa(rand.Intn(20)) + ".jpg"
		}
	} else {
		bytes := make([]byte, 4)
		_, err := rand.Read(bytes)
		if err != nil {
			utils.ERROR.Println(err)
		}
		url = header.Filename
		b, _ := io.ReadAll(file)
		os.WriteFile("/images/"+url, b, 0777)
	}
	if stErr != nil {
		status = 1
	}
	return &models.User{
		Image:    url,
		Username: username,
		Password: password,
		Email:    email,
		Status:   status,
	}
}

func UpdateUser(id, username, status, userId, password, email, currentAvatar string, file multipart.File, header *multipart.FileHeader) *utils.Message {
	userID, _ := strconv.Atoi(id)
	var lastUrl string
	img := models.GetImageByID(userID)
	for img.Next() {
		img.Scan(&lastUrl)
	}
	res := models.UpdateByID(userID, upload(username, status, userId, password, email, currentAvatar, file, header))
	if res == 0 {
		return utils.Ok(nil, "用户没有修改信息")
	}
	var url string
	img = models.GetImageByID(userID)
	for img.Next() {
		img.Scan(&url)
	}
	if lastUrl != "default.jpg" && lastUrl != url {
		if err := os.Remove("/images/" + lastUrl); err != nil {
			utils.ERROR.Println("删除文件失败", err)
		}
	}
	return utils.Ok(nil, "用户修改成功")
}

func CreateUser(username, status, userId, password, email, currentAvatar string, file multipart.File, header *multipart.FileHeader) *utils.Message {
	count := models.Exists(username)
	for count.Next() {
		var num int
		count.Scan(&num)
		if num != 0 {
			return utils.Lose(nil, "用户已经存在，请重新创建")
		}
	}
	res := models.CreateUser(upload(username, status, userId, password, email, currentAvatar, file, header))
	if res == 0 {
		return utils.Lose(nil, "用户创建失败")
	}
	return utils.Ok(nil, "用户创建成功")
}

package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"userManageSystem/utils"
)

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username,omitempty"`
	Password  string `json:"password,omitempty"`
	Email     string `json:"email,omitempty"`
	Image     string `json:"image"`
	LastLogin string `json:"last_login,default=0000-00-00 00:00:00"`
	SessionID string `json:"session_id"`
	Status    int    `json:"status"`
	Role      string `json:"role"`
}

func CreateUser(us *User) int64 {
	if us == nil {
		return 0
	}
	db := utils.InitDB()
	stmt, err := db.Prepare("INSERT INTO user (image, username, password, email, status) VALUES (?,?,?,?,?)")
	defer stmt.Close()
	defer db.Close()
	if err != nil {
		utils.ERROR.Println(err)
	}
	res, err := stmt.Exec(us.Image, us.Username, us.Password, us.Email, us.Status)
	if err != nil {
		return 0
	}
	utils.INFO.Printf("INSERT INTO user (image, username, password, email, status) VALUES ('%s', '%s', '%s', '%s', %d)",
		us.Image, us.Username, us.Password, us.Email, us.Status)
	num, _ := res.RowsAffected()
	return num
}

func Exists(username string) *sql.Rows {
	db := utils.InitDB()
	stmt, err := db.Prepare("select count(*) from user where username = ? and delete_status = 0")
	defer stmt.Close()
	defer db.Close()
	if err != nil {
		utils.ERROR.Println(err)
	}
	res, _ := stmt.Query(username)
	utils.INFO.Printf("select count(*) from user where username = '%s' and delete_status = 0", username)
	return res
}

func DeleteUser(username string) int64 {
	db := utils.InitDB()
	stmt, err := db.Prepare("update user set delete_status = ? ,delete_time = ? where username = ?")
	defer stmt.Close()
	defer db.Close()
	if err != nil {
		utils.ERROR.Println(err)
	}
	t := time.Now().Format("2006-01-02 15:04:05")
	res, _ := stmt.Exec(1, t, username)
	utils.INFO.Printf("update user set delete_status = %d ,delete_time = '%s' where username = '%s'", 1, t, username)
	num, _ := res.RowsAffected()
	return num
}

func UpdateTime(username string) {
	db := utils.InitDB()
	stmt, err := db.Prepare("update user set last_login = ? where username = ? and delete_status <> 1")
	defer stmt.Close()
	defer db.Close()
	if err != nil {
		utils.ERROR.Println(err)
	}
	t := time.Now().Format("2006-01-02 15:04:05")
	stmt.Exec(t, username)
	utils.INFO.Printf("update user set last_login = '%s' where username = '%s' and delete_status <> 1", t, username)
}

func UpdateByID(userID int, us *User) int64 {
	db := utils.InitDB()
	stmt, err := db.Prepare("update user set image = ?,username = ?,password = ?,email = ?,status = ? where id = ? and delete_status <> 1")
	defer stmt.Close()
	defer db.Close()
	if err != nil {
		utils.ERROR.Println(err)
	}
	res, _ := stmt.Exec(us.Image, us.Username, us.Password, us.Email, us.Status, userID)
	utils.INFO.Printf("update user set image = '%s',username = '%s',password = '%s',email = '%s',status = %d where id = %d and delete_status <> 1", us.Image, us.Username, us.Password, us.Email, us.Status, userID)
	num, _ := res.RowsAffected()
	return num
}

func GetUserByName(username string) *sql.Rows {
	db := utils.InitDB()
	stmt, err := db.Prepare("select image, id, username, password, email, status from user where username = ? and delete_status <> 1")
	defer stmt.Close()
	defer db.Close()
	if err != nil {
		utils.ERROR.Println(err)
	}
	res, _ := stmt.Query(username)
	utils.INFO.Printf("select image, id, username, password, email, status from user where username = '%s' and delete_status <> 1", username)
	return res
}

func GetUserByRole(sessionID string) *sql.Rows {
	db := utils.InitDB()
	stmt, err := db.Prepare("select role from user where session_id = ? and delete_status <> 1")
	defer stmt.Close()
	defer db.Close()
	if err != nil {
		utils.ERROR.Println(err)
	}
	res, _ := stmt.Query(sessionID)
	utils.INFO.Printf("select role from user where session_id = '%s' and delete_status <> 1", sessionID)
	return res
}

func GetUser(us *User) *sql.Rows {
	db := utils.InitDB()
	stmt, err := db.Prepare("select status from user where username = ? and password = ? and delete_status <> 1")
	defer stmt.Close()
	defer db.Close()
	if err != nil {
		utils.ERROR.Println(err)
	}
	res, _ := stmt.Query(us.Username, us.Password)
	utils.INFO.Printf("select status from user where username = '%s' and password = '%s' and delete_status <> 1", us.Username, us.Password)
	return res
}

func GetUserNum() *sql.Rows {
	db := utils.InitDB()
	stmt, err := db.Prepare("select count(*) from user where delete_status <> 1")
	defer stmt.Close()
	defer db.Close()
	if err != nil {
		utils.ERROR.Println(err)
	}
	res, _ := stmt.Query()
	utils.INFO.Printf("select count(*) from user where delete_status <> 1")
	return res
}

func GetUserNumByStatus(status int) *sql.Rows {
	db := utils.InitDB()
	stmt, err := db.Prepare("select count(*) from user where status = ? and delete_status <> 1")
	defer stmt.Close()
	defer db.Close()
	if err != nil {
		utils.ERROR.Println(err)
	}
	res, _ := stmt.Query(status)
	utils.INFO.Printf("select count(*) from user where status = '%s' and delete_status <> 1", status)
	return res
}

func GetUserNumByName(username string) *sql.Rows {
	db := utils.InitDB()
	stmt, err := db.Prepare("select count(*) from user where username like ? and delete_status <> 1")
	defer stmt.Close()
	defer db.Close()
	if err != nil {
		utils.ERROR.Println(err)
	}
	res, _ := stmt.Query("%" + username + "%")
	utils.INFO.Printf("select count(*) from user where username like '%s' and delete_status <> 1", "%"+username+"%")
	return res
}

func GetUserRegisterNumByMonth(month int) *sql.Rows {
	db := utils.InitDB()
	stmt, err := db.Prepare("select count(*) from user where month(register_time) = ?")
	defer stmt.Close()
	defer db.Close()
	if err != nil {
		utils.ERROR.Println(err)
	}
	res, _ := stmt.Query(month)
	utils.INFO.Printf("select count(*) from user where month(register_time) = %d", month)
	return res
}

func GetUserDeleteNumByMonth(month int) *sql.Rows {
	db := utils.InitDB()
	stmt, err := db.Prepare("select count(*) from user where month(delete_time) = ?")
	defer stmt.Close()
	defer db.Close()
	if err != nil {
		utils.ERROR.Println(err)
	}
	res, _ := stmt.Query(month)
	utils.INFO.Printf("select count(*) from user where month(delete_time) = %d", month)
	return res
}

func GetUserSessionByUsername(username string) *sql.Rows {
	db := utils.InitDB()
	stmt, err := db.Prepare("select session_id from user where username = ? and delete_status <> 1")
	defer stmt.Close()
	defer db.Close()
	if err != nil {
		utils.ERROR.Println(err)
	}
	res, _ := stmt.Query(username)
	utils.INFO.Printf("select session_id from user where username = '%s' and delete_status <> 1", username)
	return res
}

func GetMaxUserID() *sql.Rows {
	db := utils.InitDB()
	defer db.Close()
	res, _ := db.Query("select max(id) id from user")
	utils.INFO.Printf("select max(id) id from user")
	return res
}

func GetImageBySessionID(id string) *sql.Rows {
	db := utils.InitDB()
	stmt, err := db.Prepare("select image from user where session_id = ?")
	defer stmt.Close()
	defer db.Close()
	if err != nil {
		utils.ERROR.Println(err)
	}
	res, _ := stmt.Query(id)
	utils.INFO.Printf("select image from user where session_id = '%s'", id)
	return res
}

func GetImageByID(id int) *sql.Rows {
	db := utils.InitDB()
	stmt, err := db.Prepare("select image from user where id = ?")
	defer stmt.Close()
	defer db.Close()
	if err != nil {
		utils.ERROR.Println(err)
	}
	res, _ := stmt.Query(id)
	utils.INFO.Printf("select image from user where id = %d", id)
	return res
}

func GetImageByName(username string) *sql.Rows {
	db := utils.InitDB()
	stmt, err := db.Prepare("select image from user where username = ? and delete_status <> 1")
	defer stmt.Close()
	defer db.Close()
	if err != nil {
		utils.ERROR.Println(err)
	}
	res, _ := stmt.Query(username)
	utils.INFO.Printf("select image from user where id = '%s' and delete_status <> 1", username)
	return res
}

func GetPasswordByEmail(username, email string) *sql.Rows {
	db := utils.InitDB()
	stmt, err := db.Prepare("select password from user where username = ? and email = ? and delete_status <> 1")
	defer stmt.Close()
	defer db.Close()
	if err != nil {
		utils.ERROR.Println(err)
	}
	res, _ := stmt.Query(username, email)
	utils.INFO.Printf("select password from user where username = '%s' and email = '%s' and delete_status <> 1", username, email)
	return res
}

func QueryPage(limit, offset int) *sql.Rows {
	db := utils.InitDB()
	stmt, err := db.Prepare("select image,username, role, email, status, last_login from user where delete_status <> 1 limit ?,?")
	defer stmt.Close()
	defer db.Close()
	if err != nil {
		utils.ERROR.Println(err)
	}
	res, _ := stmt.Query((limit-1)*offset, offset)
	utils.INFO.Printf("select image,username, role, email, status, last_login from user where delete_status <> 1 limit %d,%d", (limit-1)*offset, offset)
	return res
}

func QueryPageByStatus(limit, offset, status int) *sql.Rows {
	db := utils.InitDB()
	stmt, err := db.Prepare("select image,username, role, email, status, last_login from user where status = ? and delete_status <> 1 limit ?,?")
	defer stmt.Close()
	defer db.Close()
	if err != nil {
		utils.ERROR.Println(err)
	}
	res, _ := stmt.Query(status, (limit-1)*offset, offset)
	utils.INFO.Printf("select image,username, role, email, status, last_login from user where status = %d and delete_status <> 1 limit %d,%d", status, (limit-1)*offset, offset)
	return res
}

func QueryPageByName(limit, offset int, username string) *sql.Rows {
	db := utils.InitDB()
	stmt, err := db.Prepare("select image,username, role, email, status, last_login from user where username like ? and delete_status <> 1 limit ?,?")
	defer stmt.Close()
	defer db.Close()
	if err != nil {
		utils.ERROR.Println(err)
	}
	res, _ := stmt.Query("%"+username+"%", (limit-1)*offset, offset)
	utils.INFO.Printf("select image,username, role, email, status, last_login from user where username like '%s' and delete_status <> 1 limit %d,%d", "%"+username+"%", (limit-1)*offset, offset)
	return res
}

func GetUsernameBySessionID(sessionID string) *sql.Rows {
	db := utils.InitDB()
	stmt, err := db.Prepare("select username from user where session_id = ? and delete_status <> 1")
	defer stmt.Close()
	defer db.Close()
	if err != nil {
		utils.ERROR.Println(err)
	}
	res, _ := stmt.Query(sessionID)
	utils.INFO.Printf("select username from user where session_id = '%s' and delete_status <> 1", sessionID)
	return res
}

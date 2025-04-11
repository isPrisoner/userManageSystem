package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"userManageSystem/service"
	"userManageSystem/utils"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("view/login.html")
	if err != nil {
		fmt.Fprintln(w, http.ErrMissingFile)
		utils.ERROR.Println(err)
	}
	t.Execute(w, nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("view/index.html")
	if err != nil {
		fmt.Fprintln(w, http.ErrMissingFile)
		log.Fatal(err)
	}
	session, _ := r.Cookie("sessionID")
	image := service.Image(session.Value)

	t.Execute(w, image)
}

func Login(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	respond := service.Login(username, password, w)
	data, _ := json.Marshal(respond)
	w.Write(data)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	sessionID, _ := r.Cookie("sessionID")
	utils.SessionDelete(w, sessionID.Value)
	http.Redirect(w, r, "/", http.StatusFound)
}

func Register(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	status := r.FormValue("status")
	userId := r.FormValue("userId")
	password := r.FormValue("password")
	email := r.FormValue("email")
	file, header, _ := r.FormFile("avatar")
	currentAvatar := r.FormValue("currentAvatar")
	respond := service.CreateUser(username, status, userId, password, email, currentAvatar, file, header)
	data, _ := json.Marshal(respond)
	w.Write(data)
}

func RegisterPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("view/register.html")
	if err != nil {
		fmt.Fprintln(w, http.ErrMissingFile)
		utils.ERROR.Println(err)
	}
	t.Execute(w, nil)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	msg := service.DeleteUser(username)
	data, _ := json.Marshal(msg)
	w.Write(data)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	var respond *utils.Message
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		username := mux.Vars(r)["username"]
		respond = service.QueryUser(username)
	} else {
		username := r.FormValue("username")
		status := r.FormValue("status")
		userId := r.FormValue("userId")
		password := r.FormValue("password")
		email := r.FormValue("email")
		file, header, _ := r.FormFile("avatar")
		currentAvatar := r.FormValue("currentAvatar")
		id := mux.Vars(r)["id"]

		respond = service.UpdateUser(id, username, status, userId, password, email, currentAvatar, file, header)
	}
	data, _ := json.Marshal(respond)
	w.Write(data)
}

func UserList(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("view/userList.html")
	if err != nil {
		http.Error(w, "Failed to parse template", http.StatusInternalServerError)
		utils.ERROR.Printf("Failed to parse template: %v", err)
		return
	}
	vars := mux.Vars(r)
	limit, err := strconv.Atoi(vars["page"])
	if err != nil {
		http.Error(w, "Invalid page parameter", http.StatusBadRequest)
		utils.ERROR.Printf("Invalid page parameter: %v", err)
		return
	}
	offset, err := strconv.Atoi(vars["page_size"])
	if err != nil {
		http.Error(w, "Invalid page_size parameter", http.StatusBadRequest)
		utils.ERROR.Printf("Invalid page_size parameter: %v", err)
		return
	}
	status, err := strconv.Atoi(vars["status"])
	if err != nil {
		status = -1
	}
	session, err := r.Cookie("sessionID")
	if err != nil {
		http.Error(w, "Session cookie not found", http.StatusUnauthorized)
		utils.ERROR.Printf("Session cookie not found: %v", err)
		return
	}
	username := vars["username"]
	pagination := service.List(limit, offset, status, username)

	role := service.ValidationRole(session.Value)
	image := service.Image(session.Value)
	err = t.Execute(w, struct {
		Role       string
		Pagination service.Pagination
		Image      string
	}{role, pagination, image})
	if err != nil {
		http.Error(w, "Failed to execute template", http.StatusInternalServerError)
		utils.ERROR.Printf("Failed to execute template: %v", err)
		return
	}

}

func StatusData(w http.ResponseWriter, r *http.Request) {
	respond := service.DataAnalysis()
	res, _ := json.Marshal(respond)
	w.Write(res)
}

func GetVisitData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	vars := mux.Vars(r)
	date := vars["days"]
	res := service.GenerateMockData(date)
	data, _ := json.Marshal(res)
	w.Write(data)
}

func Forgot(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	email := r.FormValue("email")
	res := service.Password(username, email)
	data, _ := json.Marshal(res)
	w.Write(data)
}

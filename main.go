package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"userManageSystem/handlers"
	"userManageSystem/middlewares"
	"userManageSystem/utils"
)

func main() {
	router := mux.NewRouter()
	router.Use(middlewares.BindMiddleware())

	go utils.Log()

	// 加载静态文件
	router.PathPrefix("/images/").
		Handler(http.StripPrefix("/images/",
			http.FileServer(http.Dir("images/"))))
	// 注册路由
	router.HandleFunc("/", handlers.Welcome)
	router.HandleFunc("/login", handlers.Login)
	router.HandleFunc("/index", handlers.Index)
	router.HandleFunc("/userList/{page}/{page_size}", handlers.UserList)
	router.HandleFunc("/userList/{page}/{page_size}/{status:[0-9]+}", handlers.UserList)
	router.HandleFunc("/userList/{page}/{page_size}/{username}", handlers.UserList)
	router.HandleFunc("/logout", handlers.Logout)
	router.HandleFunc("/register", handlers.Register)
	router.HandleFunc("/delete", handlers.Delete)
	router.HandleFunc("/edit/{id:[0-9]+}", handlers.Edit)
	router.HandleFunc("/edit/{username}", handlers.Edit)
	router.HandleFunc("/get_visit_data/{days}", handlers.GetVisitData)
	router.HandleFunc("/stats_data", handlers.StatusData)
	router.HandleFunc("/forgot-password", handlers.Forgot)
	router.HandleFunc("/sign_up", handlers.RegisterPage)

	log.Println("服务器启动，监听端口 8090")
	log.Fatal(http.ListenAndServe(":8090", router))
}

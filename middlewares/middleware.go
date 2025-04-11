package middlewares

import (
	"log"
	"net/http"
	"time"
	"userManageSystem/utils"
)

func Check(had http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" ||
			r.URL.Path == "/login" && r.Method == "POST" ||
			r.URL.Path == "/sign_up" ||
			r.URL.Path == "/register" && r.Method == "POST" || r.Method == "PUT" ||
			r.URL.Path == "/forgot-password" {
			had.ServeHTTP(w, r)
			return
		}
		req, err := r.Cookie("sessionID")
		if err != nil || req.Value == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		had.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func Log(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("请求开始: %s %s", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)

		log.Printf("请求结束: %s %s, 耗时: %v", r.Method, r.URL.Path, time.Since(start))
	}
	return http.HandlerFunc(fn)
}

func Catch(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				utils.ERROR.Println(err)
			}
		}()
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func BindMiddleware() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		h = Check(h)
		h = Log(h)
		h = Catch(h)
		return h
	}
}

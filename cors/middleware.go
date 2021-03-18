package cors

import (
	"net/http"
)

// Middleware :
func Middleware(handler http.Handler) http.Handler {
	// ヘッダをラHandlerFunc
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// header に書き込みたいものを書き込む
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		// Handler を生成する
		handler.ServeHTTP(w, r)
	})
}

func HtmlCors(handler http.Handler) http.Handler {
	// ヘッダをラHandlerFunc
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// header に書き込みたいものを書き込む
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Content-Type", "text/html")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		// Handler を生成する
		handler.ServeHTTP(w, r)
	})
}

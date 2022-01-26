package filter

import (
	"net/http"
)

func CorsFilter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		//.@volta 跨域过滤逻辑在这里
		origin := req.Header.Get("origin")
		resp.Header().Set("Access-Control-Allow-Origin", origin)
		resp.Header().Set("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		resp.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		resp.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		resp.Header().Set("Access-Control-Allow-Credentials", "true")
		method := req.Method
		if method == "OPTIONS" {
			resp.Write([]byte("success"))
			return
		}
		next.ServeHTTP(resp, req)
	})
}

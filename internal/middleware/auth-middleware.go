package middleware

import "net/http"

func AuthMiddleWare(next http.Handler) {
	return http.HandleFunc(func(w http.ResponseWriter, r *http.Request){
		
	})

}
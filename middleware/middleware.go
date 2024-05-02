package middleware

import (
	"fmt"
	response "github.com/NicholasLiem/GoLang_Microservice/http"
	jwt2 "github.com/NicholasLiem/GoLang_Microservice/utils/jwt"
	"net/http"
	"strings"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Authorization")

		if len(tokenStr) == 0 {
			response.ErrorResponse(rw, http.StatusBadRequest, "Token is missing")
			return
		}

		tokenStr = strings.Replace(tokenStr, "Bearer ", "", 1)
		claims, err := jwt2.VerifyJWT(tokenStr)

		if err != nil {
			response.ErrorResponse(rw, http.StatusUnauthorized, "Fail to verify JWT token: "+err.Error())
			return
		}

		fmt.Println(claims)
		next.ServeHTTP(rw, r)
	})
}

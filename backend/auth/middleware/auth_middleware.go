package middleware

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/mcfiet/goDo/utils"
)

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, strconv.Itoa(http.StatusUnauthorized)+" - "+http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		tokenParts := strings.Split(tokenString, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			http.Error(w, strconv.Itoa(http.StatusUnauthorized)+" - "+http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		tokenString = tokenParts[1]
		claims, err := utils.VerifyToken(tokenString)
		if err != nil {
			http.Error(w, strconv.Itoa(http.StatusUnauthorized)+" - "+http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user_id", claims["user_id"])
		log.Println("User ID:", claims["user_id"])
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

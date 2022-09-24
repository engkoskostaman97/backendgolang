package middleware

import (
	"context"
	dto "dumbflix/dto/result"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

func ChekAdmin(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
		userStatus := userInfo["status"].(string)
		fmt.Println("ini status", userStatus)
		if userStatus != "admin" {
			w.WriteHeader(http.StatusUnauthorized)
			response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Maaf Anda Bukan Admin"}
			json.NewEncoder(w).Encode(response)
			return

		}
		ctx := context.WithValue(r.Context(), "userInfo", r.Context().Value("userInfo"))
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

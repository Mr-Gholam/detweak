package main

import (
	"context"
	"net/http"

	"github.com/golang-jwt/jwt"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, X-HTTP-Method-Override, Content-Type, Accept")
		next.ServeHTTP(w, r)
	})
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/login" || r.URL.Path == "/api/signup" {
			next.ServeHTTP(w, r)
			return
		}
		cookie, err := r.Cookie("jwt")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		token, _ := jwt.Parse(cookie.Value,
			func(token *jwt.Token) (interface{}, error) {
				return jwtSecret, nil
			},
		)
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			ctx := r.Context()
			Id := uint(claims["id"].(float64))
			handleError(err)
			r = r.WithContext(context.WithValue(ctx, "user", &User{Username: claims["username"].(string), ID: Id}))
			next.ServeHTTP(w, r)
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
	})
}

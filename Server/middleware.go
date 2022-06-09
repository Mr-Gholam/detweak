package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

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
			r = r.WithContext(context.WithValue(ctx, "user", &User{Username: claims["username"].(string), ID: Id, ImgUrl: claims["imgUrl"].(string), Firstname: claims["firstname"].(string), Lastname: claims["lastname"].(string)}))
			next.ServeHTTP(w, r)
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
	})
}

func getIdFromCookie(w http.ResponseWriter, r *http.Request) uint {
	cookie, err := r.Cookie("jwt")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return 0
	}
	token, _ := jwt.Parse(cookie.Value,
		func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		},
	)
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return uint(claims["id"].(float64))
	}
	return 0
}
func FileUpload(r *http.Request) (string, error) {
	r.ParseMultipartForm(32 << 20)

	file, handler, err := r.FormFile("image") // Retrieve the file from form data

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close() // Close the file when we finish
	JavascriptISOString := "2006-01-02T15-04-05.999Z07-00"
	currentTime := time.Now().UTC().Format(JavascriptISOString)
	imgUrl := currentTime + "-" + handler.Filename
	f, err := os.OpenFile("./images/"+imgUrl, os.O_WRONLY|os.O_CREATE, 0666)

	defer f.Close()

	if err != nil {
		fmt.Println(err)
	}

	// Copy the file to the destination path
	io.Copy(f, file)

	return imgUrl, nil
}
func DeleteFile(ImgUrl string) {

	if ImgUrl != "" {
		err := os.Remove("./images/" + ImgUrl)

		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

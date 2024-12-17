package main

import (
	"fmt"
	"net/http"
)

func WriteConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request url", r.URL)
		fmt.Println("request host", r.Host)
		next.ServeHTTP(w, r)
	})
}

func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}

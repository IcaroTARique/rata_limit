package main

import (
	"fmt"
	"github.com/IcaroTARique/rate_limit/configs"
	"github.com/IcaroTARique/rate_limit/pkg/middleware"
	"github.com/go-chi/chi"
	"net/http"
)

func main() {
	conf, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()

	r.Use(func(next http.Handler) http.Handler {
		return middleware.TokenRateLimit(next, *conf)
	})
	r.Use(func(next http.Handler) http.Handler {
		return middleware.IPRateLimit(next, *conf)
	})
	r.Get("/status", func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		str := fmt.Sprintf("Hello %s\n", ip)
		w.Write([]byte(str))
	})

	if err := http.ListenAndServe(":"+conf.Port, r); err != nil {
		panic(err)
	}
}

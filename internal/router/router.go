package router

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/vitalykrupin/loyalty-system/internal/middleware"
)

func Route() {
	router := chi.NewRouter()
	router.Use(middleware.Logging)

	errListen := http.ListenAndServe(":8080", router)
	if errListen != nil {
		log.Fatalf("ListenAndServe returns error: %v", errListen)
	}
}
package router

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/vitalykrupin/loyalty-system/internal/app"
	"github.com/vitalykrupin/loyalty-system/internal/handlers"
	"github.com/vitalykrupin/loyalty-system/internal/middleware"
)

func Route(app *app.App) {
	r := chi.NewRouter()
	r.Use(middleware.Logging)
	r.Use(middleware.GzipCompress)

	r.Method(http.MethodPost, "/api/user/register", handlers.NewPostRegister(app))
	// r.Method(http.MethodPost, "/api/user/login", handlers.NewPostLogin(app))

	r.Group(func(r chi.Router) {
		// r.Use(middleware.IsAuth)
		r.Method(http.MethodGet, "/api/user/balance", handlers.NewGetBalance(app))
		// r.Method(http.MethodPost, "/api/user/balance/withdraw", handlers.NewWithdraw(app))
		// r.Method(http.MethodGet, "/api/user/withdrawals", handlers.NewGetWithdrawals(app))
		// r.Method(http.MethodPost, "/api/user/orders", handlers.NewPostOrders(app))
		// r.Method(http.MethodGet, "/api/user/orders", handlers.NewGetOrders(app))
	})

	errListen := http.ListenAndServe(":8080", r)
	if errListen != nil {
		log.Fatalf("ListenAndServe returns error: %v", errListen)
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

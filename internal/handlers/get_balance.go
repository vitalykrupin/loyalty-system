package handlers

import (
	"net/http"

	"github.com/vitalykrupin/loyalty-system/internal/app"
)

type GetBalance struct {
	BaseHandler
}

func NewGetBalance(app *app.App) *GetBalance {
	return &GetBalance{
		BaseHandler: BaseHandler{
			app: app,
		},
	}
}

func (*GetBalance) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

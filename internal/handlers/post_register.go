package handlers

import (
	"net/http"

	"github.com/vitalykrupin/loyalty-system/internal/app"
)

type PostRegister struct {
	BaseHandler
}

func NewPostRegister(app *app.App) *PostRegister {
	return &PostRegister{
		BaseHandler: BaseHandler{
			app: app,
		},
	}
}

func (*PostRegister) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

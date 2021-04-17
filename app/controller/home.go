package controller

import (
	"github.com/kyleu/admini/app/ctx"
	"github.com/kyleu/admini/views"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	act("home", w, r, func(st *ctx.PageState) (string, error) {
		views.WriteRender(w, &views.Home{Basic: with(st)})
		return "", nil
	})
}

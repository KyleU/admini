package controller

import (
	"net/http"

	"github.com/kyleu/admini/app/ctx"
	"github.com/kyleu/admini/views/vsearch"
)

func Search(w http.ResponseWriter, r *http.Request) {
	act("home", w, r, func(app *ctx.AppState, page *ctx.PageState) (string, error) {
		q := r.URL.Query().Get("q")
		return render(w, &vsearch.SearchResults{Q: q}, page, "search")
	})
}

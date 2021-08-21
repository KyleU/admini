// Code generated by Project Forge, see https://projectforge.dev for details.
package controller

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/site"
	"github.com/kyleu/admini/app/telemetry/httpmetrics"
	"github.com/kyleu/admini/app/util"
)

func SiteRoutes() fasthttp.RequestHandler {
	r := router.New()

	r.GET("/", Site)

	r.GET(defaultProfilePath, ProfileSite)
	r.POST(defaultProfilePath, ProfileSave)
	r.GET("/auth/{key}", AuthDetail)
	r.GET("/auth/callback/{key}", AuthCallback)
	r.GET("/auth/logout/{key}", AuthLogout)

	r.GET("/favicon.ico", Favicon)
	r.GET("/assets/{_:*}", Static)

	r.GET("/{path:*}", Site)

	r.OPTIONS("/", Options)
	r.OPTIONS("/{_:*}", Options)
	r.NotFound = NotFound

	m := httpmetrics.NewMetrics("marketing_site")
	return fasthttp.CompressHandler(m.WrapHandler(r))
}

func Site(ctx *fasthttp.RequestCtx) {
	actSite("site", ctx, func(as *app.State, ps *cutil.PageState) (string, error) {
		path := util.SplitAndTrim(string(ctx.Request.URI().Path()), "/")
		redir, page, bc, err := site.Handle(path, ctx, as, ps)
		if err != nil {
			return "", err
		}
		if redir != "" {
			return redir, nil
		}
		return render(ctx, as, page, ps, bc...)
	})
}

package controller

import (
	"github.com/valyala/fasthttp"

	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/settings"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/views/vsettings"
)

func Settings(ctx *fasthttp.RequestCtx) {
	act("settings", ctx, func(as *app.State, ps *cutil.PageState) (string, error) {
		current := &settings.Settings{}
		ps.Title = "Settings"
		ps.Data = current
		return render(ctx, as, &vsettings.Settings{Settings: current}, ps, "settings")
	})
}

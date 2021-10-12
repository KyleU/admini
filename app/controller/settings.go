// Code generated by projectforge.dev using code from the [core] module, which is under license [CC0]
package controller

import (
	"github.com/valyala/fasthttp"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/settings"
	"github.com/kyleu/admini/app/user"
	"github.com/kyleu/admini/views/vsettings"
)

func Settings(rc *fasthttp.RequestCtx) {
	act("settings", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		current := &settings.Settings{}
		ps.Title = "Settings"
		ps.Data = current
		return render(rc, as, &vsettings.Settings{Settings: current, Perms: user.GetPermissions()}, ps, "settings")
	})
}

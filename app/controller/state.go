// Content managed by Project Forge, see [projectforge.md] for details.
package controller

import (
	"fmt"

	"github.com/valyala/fasthttp"

	"admini.dev/admini/app"
	"admini.dev/admini/app/controller/cmenu"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/lib/theme"
	"admini.dev/admini/app/util"
	"admini.dev/admini/views/verror"
)

var (
	_currentAppState       *app.State
	_currentAppRootLogger  util.Logger
	_currentSiteState      *app.State
	_currentSiteRootLogger util.Logger
	defaultRootTitleAppend = util.GetEnv("app_display_name_append")
	defaultRootTitle       = func() string {
		if tmp := util.GetEnv("app_display_name"); tmp != "" {
			return tmp
		}
		return util.AppName
	}()
)

func SetAppState(a *app.State, logger util.Logger) {
	_currentAppState = a
	_currentAppRootLogger = logger
	initApp(a, logger)
}

func SetSiteState(a *app.State, logger util.Logger) {
	_currentSiteState = a
	_currentSiteRootLogger = logger
	initSite(a, logger)
}

func handleError(key string, as *app.State, ps *cutil.PageState, rc *fasthttp.RequestCtx, err error) (string, error) {
	rc.SetStatusCode(fasthttp.StatusInternalServerError)

	ps.Logger.Errorf("error running action [%s]: %+v", key, err)

	if len(ps.Breadcrumbs) == 0 {
		bc := util.StringSplitAndTrim(string(rc.URI().Path()), "/")
		bc = append(bc, "Error")
		ps.Breadcrumbs = bc
	}

	if cleanErr := clean(as, ps); cleanErr != nil {
		ps.Logger.Error(cleanErr)
		msg := fmt.Sprintf("error while cleaning request: %+v", cleanErr)
		ps.Logger.Error(msg)
		_, _ = rc.WriteString(msg)
		return "", cleanErr
	}

	redir, renderErr := Render(rc, as, &verror.Error{Err: util.GetErrorDetail(err)}, ps)
	if renderErr != nil {
		msg := fmt.Sprintf("error while running error handler: %+v", renderErr)
		ps.Logger.Error(msg)
		_, _ = rc.WriteString(msg)
		return "", renderErr
	}
	return redir, nil
}

func clean(as *app.State, ps *cutil.PageState) error {
	if ps.Profile != nil && ps.Profile.Theme == "" {
		ps.Profile.Theme = theme.ThemeDefault.Key
	}
	if ps.RootIcon == "" {
		ps.RootIcon = defaultIcon
	}
	if ps.RootPath == "" {
		ps.RootPath = "/"
	}
	if ps.RootTitle == "" {
		ps.RootTitle = defaultRootTitle
	}
	if defaultRootTitleAppend != "" {
		ps.RootTitle += " " + defaultRootTitleAppend
	}
	if ps.SearchPath == "" {
		ps.SearchPath = DefaultSearchPath
	}
	if ps.ProfilePath == "" {
		ps.ProfilePath = DefaultProfilePath
	}
	if len(ps.Menu) == 0 {
		m, err := cmenu.MenuFor(ps.Context, ps.Authed, ps.Admin, as, ps.Logger)
		if err != nil {
			return err
		}
		ps.Menu = m
	}
	return nil
}

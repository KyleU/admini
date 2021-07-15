// Package controller: $PF_IGNORE$
package controller

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

//nolint
func AppRoutes() *router.Router {
	w := fasthttp.CompressHandler
	r := router.New()
	r.GET("/", w(Home))

	r.GET("/settings", w(Settings))
	r.GET("/theme", w(ThemeList))
	r.GET("/theme/{key}", w(ThemeEdit))
	r.POST("/theme/{key}", w(ThemeSave))
	r.GET("/refresh", w(Refresh))
	r.GET(defaultSearchPath, w(Search))

	r.GET(defaultProfilePath, w(Profile))
	r.POST(defaultProfilePath, w(ProfileSave))
	r.GET("/auth/{key}", w(AuthDetail))
	r.GET("/auth/{key}/callback", w(AuthCallback))
	r.GET("/auth/{key}/logout", w(AuthLogout))

	r.GET("/admin", w(Admin))
	r.GET("/admin/{path:*}", w(Admin))

	r.GET("/help", w(Help))
	r.GET("/help/{path:*}", w(Help))

	r.GET("/source", w(SourceList))
	r.POST("/source", w(SourceInsert))
	r.GET("/source/_new", w(SourceNew))
	r.GET("/source/{key}", w(SourceDetail))
	r.GET("/source/{key}/edit", w(SourceEdit))
	r.POST("/source/{key}", w(SourceSave))
	r.GET("/source/{key}/refresh", w(SourceRefresh))
	r.GET("/source/{key}/delete", w(SourceDelete))
	r.GET("/source/{key}/model/{path:*}", w(SourceModelDetail))
	r.POST("/source/{key}/model/{path:*}", w(SourceModelSave))

	r.GET("/project", w(ProjectList))
	r.POST("/project", w(ProjectInsert))
	r.GET("/project/_new", w(ProjectNew))
	r.GET("/project/{key}", w(ProjectDetail))
	r.POST("/project/{key}", w(ProjectSave))
	r.GET("/project/{key}/edit", w(ProjectEdit))
	r.POST("/project/{key}/actions", w(ActionOrdering))
	r.GET("/project/{key}/action/{path:*}", w(ActionEdit))
	r.POST("/project/{key}/action/{path:*}", w(ActionSave))
	r.GET("/project/{key}/test", w(ProjectTest))
	r.GET("/project/{key}/delete", w(ProjectDelete))

	r.GET("/x/{key}", w(WorkspaceProject))
	r.GET("/x/{key}/{_:*}", w(WorkspaceProject))
	r.POST("/x/{key}/{_:*}", w(WorkspaceProject))
	r.GET("/s/{key}", w(WorkspaceSource))
	r.GET("/s/{key}/{_:*}", w(WorkspaceSource))
	r.POST("/s/{key}/{_:*}", w(WorkspaceSource))

	r.GET("/sandbox", w(SandboxList))
	r.GET("/sandbox/{key}", w(SandboxRun))

	r.GET("/modules", w(Modules))

	r.GET("/favicon.ico", Favicon)
	r.GET("/robots.txt", RobotsTxt)
	r.GET("/assets/{_:*}", Static)

	r.OPTIONS("/", w(Options))
	r.OPTIONS("/{_:*}", w(Options))
	r.NotFound = NotFound

	return r
}

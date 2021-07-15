package controller

import (
	"fmt"

	"github.com/kyleu/admini/app/action"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/util"
	"github.com/kyleu/admini/app/workspace"
	"github.com/pkg/errors"

	"github.com/kyleu/admini/app"
)

func actWorkspace(key string, ctx *fasthttp.RequestCtx, f func(as *app.State, ps *cutil.PageState) (string, error)) {
	as, ps := actPrepare("workspace", ctx)
	actComplete(key, as, ps, ctx, f)
}

func WorkspaceProject(ctx *fasthttp.RequestCtx) {
	actWorkspace("workspace", ctx, func(as *app.State, ps *cutil.PageState) (string, error) {
		projectKey, err := ctxRequiredString(ctx, "key", false)
		if err != nil {
			return "", err
		}

		paths := util.SplitAndTrim(string(ctx.Path()), "/")
		if len(paths) < 2 {
			return ersp("no source provided in path [%s]", string(ctx.Path()))
		}
		if paths[0] != "x" {
			return ersp("provided path [%s] is not part of the project workspace", string(ctx.Path()))
		}

		pv, err := as.Projects.LoadView(projectKey)
		if err != nil {
			return "", errors.Wrapf(err, "error loading project [%s]", projectKey)
		}

		paths = paths[2:]

		ps.Title = pv.Project.Name()
		ps.RootIcon = pv.Project.IconWithFallback()
		ps.RootPath = fmt.Sprintf("/x/%s", pv.Project.Key)
		ps.RootTitle = pv.Project.Name()
		ps.SearchPath = defaultSearchPath
		ps.ProfilePath = defaultProfilePath

		m, err := workspace.ProjectMenu(as, pv)
		if err != nil {
			return "", errors.Wrapf(err, "error creating menu for project [%s]", pv.Project.Key)
		}
		ps.Menu = m

		a, remaining := pv.Project.Actions.Get(paths)

		wr := &cutil.WorkspaceRequest{
			T: "x", K: pv.Project.Key, Ctx: ctx, AS: as, PS: ps, Item: a, Path: remaining,
			Project: pv.Project, Sources: pv.Sources, Schemata: pv.Schemata,
		}
		return handleAction(wr, a, ctx, ps)
	})
}

func handleAction(req *cutil.WorkspaceRequest, act *action.Action, ctx *fasthttp.RequestCtx, ps *cutil.PageState) (string, error) {
	if req == nil {
		return "", errors.New("nil project request")
	}
	if act == nil {
		act = &action.Action{}
	}
	res, err := workspace.ActionHandler(req, act)
	if err != nil {
		return "", err
	}

	if res.Redirect != "" {
		return flashAndRedir(true, res.Title, res.Redirect, ctx, ps)
	}

	req.PS.Title = res.Title
	req.PS.Data = res.Data

	return render(req.Ctx, req.AS, res.Page, req.PS, res.Breadcrumbs...)
}

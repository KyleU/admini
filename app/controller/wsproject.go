package controller

import (
	"net/http"

	"github.com/kyleu/admini/app/project/action"

	"github.com/gorilla/mux"
	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/util"
	"github.com/kyleu/admini/app/workspace"
	"github.com/pkg/errors"

	"github.com/kyleu/admini/app"
)

func WorkspaceProject(w http.ResponseWriter, r *http.Request) {
	actWorkspace("workspace", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		projectKey := mux.Vars(r)["key"]

		paths := util.SplitAndTrim(r.URL.Path, "/")
		if len(paths) < 2 {
			return ersp("no source provided in path [%v]", r.URL.Path)
		}
		if paths[0] != "x" {
			return ersp("provided path [%v] is not part of the project workspace", r.URL.Path)
		}

		pv, err := currentApp.Projects.LoadView(projectKey)
		if err != nil {
			return "", errors.Wrap(err, "error loading project ["+projectKey+"]")
		}

		paths = paths[2:]

		ps.Title = pv.Project.Name()
		ps.RootPath = currentApp.Route("workspace", "key", pv.Project.Key)
		ps.RootTitle = pv.Project.Name()
		ps.SearchPath = currentApp.Route("search")
		ps.ProfilePath = currentApp.Route("profile")

		m, err := workspace.ProjectMenu(currentApp, pv)
		if err != nil {
			return "", errors.Wrap(err, "error creating menu for project ["+pv.Project.Key+"]")
		}
		ps.Menu = m

		a, remaining := pv.Project.Actions.Get(paths)

		wr := &cutil.WorkspaceRequest{T: "workspace", K: pv.Project.Key, W: w, R: r, AS: as, PS: ps, Item: a, Path: remaining, Project: pv.Project, Sources: pv.Sources, Schemata: pv.Schemata}
		return handleAction(wr, a)
	})
}

func handleAction(req *cutil.WorkspaceRequest, act *action.Action) (string, error) {
	if req == nil {
		return whoops(req, "nil project request")
	}
	if act == nil {
		act = &action.Action{}
	}
	res, err := workspace.ActionHandler(req, act)
	if err != nil {
		return "", err
	}
	req.PS.Title = res.Title
	req.PS.Data = res.Data

	return renderWS(req, res.Page, res.Breadcrumbs...)
}

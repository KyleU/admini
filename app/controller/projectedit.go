package controller

import (
	"fmt"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/project"
	"github.com/kyleu/admini/views/vproject"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"
)

func ProjectNew(ctx *fasthttp.RequestCtx) {
	act("project.new", ctx, func(as *app.State, ps *cutil.PageState) (string, error) {
		ps.Title = "New Project"
		p := &project.Project{}
		ps.Data = p
		avail, err := as.Services.Sources.List()
		if err != nil {
			return "", errors.Wrap(err, "unable to list sources")
		}
		return render(ctx, as, &vproject.New{Project: p, AvailableSources: avail}, ps, "projects", "New")
	})
}

func ProjectInsert(ctx *fasthttp.RequestCtx) {
	act("project.insert", ctx, func(as *app.State, ps *cutil.PageState) (string, error) {
		frm, err := cutil.ParseForm(ctx)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse form")
		}
		key, err := frm.GetString("key", false)
		if err != nil {
			return flashAndRedir(false, err.Error(), "/project/_new", ctx, ps)
		}
		title := frm.GetStringOpt("title")
		icon := frm.GetStringOpt("icon")
		description := frm.GetStringOpt("description")
		sources, err := frm.GetStringArray("sources", true)
		if err != nil {
			return "", err
		}
		ret := &project.Project{Key: key, Title: title, Icon: icon, Description: description, Sources: sources}
		err = as.Services.Projects.Save(ret, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to save project")
		}
		return flashAndRedir(true, "saved new project", fmt.Sprintf("/project/%s", key), ctx, ps)
	})
}

func ProjectEdit(ctx *fasthttp.RequestCtx) {
	act("project.edit", ctx, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := ctxRequiredString(ctx, "key", false)
		if err != nil {
			return "", err
		}
		prj, err := as.Services.Projects.LoadRequired(key, false)
		if err != nil {
			return "", errors.Wrapf(err, "unable to load project [%s]", key)
		}
		ps.Title = fmt.Sprintf("Edit [%s]", prj.Name())
		ps.Data = prj

		avail, err := as.Services.Sources.List()
		if err != nil {
			return "", errors.Wrap(err, "unable to list sources")
		}
		return render(ctx, as, &vproject.Edit{Project: prj, AvailableSources: avail}, ps, "projects", prj.Key, "Edit")
	})
}

func ProjectSave(ctx *fasthttp.RequestCtx) {
	act("project.save", ctx, func(as *app.State, ps *cutil.PageState) (string, error) {
		frm, err := cutil.ParseForm(ctx)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse form")
		}

		key, err := ctxRequiredString(ctx, "key", false)
		if err != nil {
			return "", err
		}

		prj, err := as.Services.Projects.LoadRequired(key, false)
		if err != nil {
			return "", errors.Wrapf(err, "unable to load project [%s]", key)
		}

		prj.Title = frm.GetStringOpt("title")
		prj.Icon = frm.GetStringOpt("icon")
		prj.Description = frm.GetStringOpt("description")
		prj.Sources, err = frm.GetStringArray("sources", true)
		if err != nil {
			return "", err
		}

		err = as.Services.Projects.Save(prj, true)
		if err != nil {
			return "", errors.Wrapf(err, "unable to save project [%s]", key)
		}

		msg := fmt.Sprintf(`saved project "%s"`, key)
		return flashAndRedir(true, msg, fmt.Sprintf("/project/%s", key), ctx, ps)
	})
}

func ProjectDelete(ctx *fasthttp.RequestCtx) {
	act("project.delete", ctx, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := ctxRequiredString(ctx, "key", false)
		if err != nil {
			return "", err
		}
		err = as.Services.Projects.Delete(key)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete project [%s]", key)
		}

		msg := fmt.Sprintf(`deleted project "%s"`, key)
		return flashAndRedir(true, msg, "/project", ctx, ps)
	})
}

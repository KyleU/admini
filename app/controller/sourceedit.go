package controller

import (
	"fmt"
	"strconv"

	"github.com/kyleu/admini/app/database"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/admini/app/schema"
	"github.com/kyleu/admini/app/source"
	"github.com/kyleu/admini/app/util"
	"github.com/pkg/errors"

	"github.com/kyleu/admini/app/controller/cutil"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/views/vsource"
)

func SourceNew(rc *fasthttp.RequestCtx) {
	act("source.new", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ps.Title = "New Source"
		t := schema.OriginPostgres
		s := &source.Source{Type: t}
		ps.Data = s
		return render(rc, as, &vsource.New{Origin: t}, ps, "sources", "New")
	})
}

func SourceInsert(rc *fasthttp.RequestCtx) {
	act("source.insert", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		frm, err := cutil.ParseForm(rc)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse form")
		}

		key, err := frm.GetString("key", false)
		if err != nil {
			return flashAndRedir(false, err.Error(), "/source/_new", rc, ps)
		}
		title := frm.GetStringOpt("title")
		icon := frm.GetStringOpt("icon")
		description := frm.GetStringOpt("description")
		typ := frm.GetStringOpt("type")
		ret := as.Services.Sources.NewSource(key, title, icon, description, schema.OriginFromString(typ))
		err = as.Services.Sources.Save(ret, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to save source")
		}
		return flashAndRedir(true, "saved new source", fmt.Sprintf("/source/%s", key), rc, ps)
	})
}

func SourceEdit(rc *fasthttp.RequestCtx) {
	act("source.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := rcRequiredString(rc, "key", false)
		if err != nil {
			return "", err
		}
		src, err := as.Services.Sources.Load(key, false)
		if err != nil {
			return "", errors.Wrapf(err, "unable to load source [%s]", key)
		}
		ps.Title = fmt.Sprintf("Edit [%s]", src.Name())
		ps.Data = src

		switch src.Type {
		case schema.OriginPostgres, schema.OriginSQLite:
			page := &vsource.Edit{Source: src}
			return render(rc, as, page, ps, "sources", src.Key, "Edit")
		default:
			msg := fmt.Sprintf("unhandled source type [%s]", src.Type.String())
			return flashAndRedir(false, msg, "/source", rc, ps)
		}
	})
}

func SourceSave(rc *fasthttp.RequestCtx) {
	act("source.save", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		frm, err := cutil.ParseForm(rc)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse form")
		}

		key, err := rcRequiredString(rc, "key", false)
		if err != nil {
			return "", err
		}
		src, err := as.Services.Sources.Load(key, false)
		if err != nil {
			return "", errors.Wrapf(err, "unable to load source [%s]", key)
		}

		src.Title = frm.GetStringOpt("title")
		src.Icon = frm.GetStringOpt("icon")
		src.Description = frm.GetStringOpt("description")

		switch src.Type {
		case schema.OriginPostgres:
			ps := frm.GetStringOpt("port")
			params := &database.PostgresParams{}
			params.Host = frm.GetStringOpt("host")
			if ps != "" {
				params.Port, _ = strconv.Atoi(ps)
			}
			params.Username = frm.GetStringOpt("username")
			params.Password = frm.GetStringOpt("password")
			params.Database = frm.GetStringOpt("database")
			params.Schema = frm.GetStringOpt("schema")
			params.Debug, _ = frm.GetBool("debug")

			src.Config = util.ToJSONBytes(params, true)
		case schema.OriginSQLite:
			params := &database.SQLiteParams{}
			params.File = frm.GetStringOpt("file")
			params.Schema = frm.GetStringOpt("schema")
			params.Debug, _ = frm.GetBool("debug")

			src.Config = util.ToJSONBytes(params, true)
		default:
			return "", errors.Errorf("unable to parse config for source type [%s]", src.Type.String())
		}

		err = as.Services.Sources.Save(src, true)
		if err != nil {
			return "", errors.Wrapf(err, "unable to save source [%s]", key)
		}

		msg := fmt.Sprintf(`saved source "%s"`, key)
		return flashAndRedir(true, msg, fmt.Sprintf("/source/%s", key), rc, ps)
	})
}

func SourceDelete(rc *fasthttp.RequestCtx) {
	act("source.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := rcRequiredString(rc, "key", false)
		if err != nil {
			return "", err
		}
		err = as.Services.Sources.Delete(key)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete source [%s]", key)
		}

		msg := fmt.Sprintf(`deleted source "%s"`, key)
		return flashAndRedir(true, msg, "/source", rc, ps)
	})
}

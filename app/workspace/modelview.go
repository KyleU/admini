package workspace

import (
	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/action"
	"github.com/kyleu/admini/views/vmodel"

	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/model"
	"github.com/pkg/errors"
)

func processModelList(req *cutil.WorkspaceRequest, act *action.Action, srcKey string, m *model.Model, as *app.State) (*Result, error) {
	switch m.Type {
	case model.TypeStruct:
		_, ld, err := loaderFor(req, srcKey, as)
		if err != nil {
			return ErrResult(req, act, err)
		}

		optsMap := optionsFor(req)
		opts := optsMap.Get(m.Key)
		rs, err := ld.List(req.Context, m, opts)
		if err != nil {
			return ErrResult(req, act, errors.Wrapf(err, "unable to list model [%s]", m.Key))
		}
		page := &vmodel.List{Req: req, Act: act, Model: m, Options: opts, Result: rs}
		return NewResult("", nil, req, act, rs, page), nil
	case model.TypeEnum:
		page := &vmodel.Enum{Req: req, Act: act, Model: m}
		return NewResult("", nil, req, act, m, page), nil
	default:
		return nil, errors.Errorf("unhandled model type [%s]", m.Type.String())
	}
}

func processModelView(req *cutil.WorkspaceRequest, act *action.Action, srcKey string, m *model.Model, idStrings []string, as *app.State) (*Result, error) {
	_, ld, err := loaderFor(req, srcKey, as)
	if err != nil {
		return ErrResult(req, act, err)
	}

	data, err := getModel(req.Context, m, idStrings, ld)
	if err != nil {
		return ErrResult(req, act, err)
	}

	page := &vmodel.View{Req: req, Act: act, Model: m, Result: data}
	idx := len(req.Path) - len(idStrings) - 1
	if idx < 0 {
		idx = 0
	}
	bc := append(append(act.Path(), req.Path[:idx]...), idStrings...)
	return NewResult("", bc, req, act, data, page), nil
}

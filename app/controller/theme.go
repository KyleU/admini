package controller

import (
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/lib/theme"
	"github.com/kyleu/admini/app/util"
	"github.com/kyleu/admini/views/vtheme"
)

func ThemeList(rc *fasthttp.RequestCtx) {
	act("theme.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ps.Title = "Themes"
		x := as.Themes.All()
		ps.Data = x
		return render(rc, as, &vtheme.List{Themes: x}, ps, "admin", "Themes")
	})
}

func ThemeEdit(rc *fasthttp.RequestCtx) {
	act("theme.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := rcRequiredString(rc, "key", false)
		if err != nil {
			return "", err
		}
		var t *theme.Theme
		if key == theme.KeyNew {
			t = theme.ThemeDefault.Clone(key)
		} else {
			t = as.Themes.Get(key)
		}
		if t == nil {
			return "", errors.Wrap(err, "no theme found with key ["+key+"]")
		}
		ps.Data = t
		ps.Title = "Edit theme [" + t.Key + "]"
		page := &vtheme.Edit{Theme: t}
		return render(rc, as, page, ps, "admin", "Themes||/admin/theme", t.Key)
	})
}

func ThemeSave(rc *fasthttp.RequestCtx) {
	act("theme.save", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := rcRequiredString(rc, "key", false)
		if err != nil {
			return "", err
		}
		frm, err := cutil.ParseForm(rc)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse form")
		}

		orig := as.Themes.Get(key)

		newKey, err := frm.GetString("key", false)
		if err != nil {
			return "", err
		}
		if newKey == theme.KeyNew {
			newKey = util.RandomString(12)
		}

		l := orig.Light.Clone().ApplyMap(frm, "light-")
		d := orig.Dark.Clone().ApplyMap(frm, "dark-")

		t := &theme.Theme{Key: newKey, Light: l, Dark: d}

		err = as.Themes.Save(t)
		if err != nil {
			return "", errors.Wrap(err, "unable to save theme")
		}

		ps.Profile.Theme = newKey
		err = cutil.SaveProfile(ps.Profile, rc, ps.Session, ps.Logger)
		if err != nil {
			return "", err
		}

		msg := "saved changes to theme [" + newKey + "]"
		return returnToReferrer(msg, "/", rc, ps)
	})
}

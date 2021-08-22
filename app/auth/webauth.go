// Code generated by Project Forge, see https://projectforge.dev for details.
package auth

import (
	"github.com/go-gem/sessions"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"

	"github.com/kyleu/admini/app/web"
)

func getAuthURL(prv *Provider, rc *fasthttp.RequestCtx, websess *sessions.Session, logger *zap.SugaredLogger) (string, error) {
	g, err := gothFor(rc, prv)
	if err != nil {
		return "", err
	}

	sess, err := g.BeginAuth(setState(rc))
	if err != nil {
		return "", err
	}

	u, err := sess.GetAuthURL()
	if err != nil {
		return "", err
	}

	err = web.StoreInSession(prv.ID, sess.Marshal(), rc, websess, logger)
	if err != nil {
		return "", err
	}

	return u, err
}

func getCurrentAuths(websess *sessions.Session) web.Accounts {
	authS, err := web.GetFromSession(WebSessKey, websess)
	var ret web.Accounts
	if err == nil && authS != "" {
		ret = web.AccountsFromString(authS)
	}
	return ret
}

func setCurrentAuths(s web.Accounts, rc *fasthttp.RequestCtx, websess *sessions.Session, logger *zap.SugaredLogger) error {
	s.Sort()
	if len(s) == 0 {
		return web.RemoveFromSession(WebSessKey, rc, websess, logger)
	}
	return web.StoreInSession(WebSessKey, s.String(), rc, websess, logger)
}

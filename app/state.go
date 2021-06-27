package app

import (
	"github.com/fasthttp/router"
	"github.com/kyleu/admini/app/filesystem"
	"github.com/kyleu/admini/app/loader"
	"github.com/kyleu/admini/app/project"
	"github.com/kyleu/admini/app/source"
	"go.uber.org/zap"
)

type State struct {
	Debug        bool
	Router       *router.Router
	Files        filesystem.FileLoader
	Sources      *source.Service
	Projects     *project.Service
	Loaders      *loader.Service
	RootLogger   *zap.SugaredLogger
	routerLogger *zap.SugaredLogger
}

func NewState(debug bool, r *router.Router, f filesystem.FileLoader, ls *loader.Service, log *zap.SugaredLogger) (*State, error) {
	rl := log.With(zap.String("service", "router"))
	ss := source.NewService(f, ls, log)
	ps := project.NewService(f, ss, ls, log)

	ret := &State{Debug: debug, Router: r, Files: f, Sources: ss, Projects: ps, Loaders: ls, RootLogger: log, routerLogger: rl}
	return ret, nil
}

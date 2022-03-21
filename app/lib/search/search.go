// Content managed by Project Forge, see [projectforge.md] for details.
package search

import (
	"context"
	"strings"
	"sync"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"admini.dev/admini/app"
	"admini.dev/admini/app/lib/search/result"
	"admini.dev/admini/app/lib/telemetry"
)

type Provider func(context.Context, *app.State, *Params, *zap.SugaredLogger) (result.Results, error)

func Search(ctx context.Context, as *app.State, params *Params) (result.Results, []error) {
	ctx, span, logger := telemetry.StartSpan(ctx, "search", as.Logger)
	defer span.Complete()

	if params.Q == "" {
		return nil, nil
	}
	var allProviders []Provider
	// $PF_SECTION_START(search_functions)$
	projectFunc := func(ctx context.Context, as *app.State, p *Params, logger *zap.SugaredLogger) (result.Results, error) {
		return as.Services.Projects.Search(ctx, p.Q)
	}
	sourceFunc := func(ctx context.Context, as *app.State, p *Params, logger *zap.SugaredLogger) (result.Results, error) {
		return as.Services.Sources.Search(p.Q)
	}
	allProviders = append(allProviders, projectFunc, sourceFunc)
	// $PF_SECTION_END(search_functions)$
	// $PF_INJECT_START(codegen)$
	// $PF_INJECT_END(codegen)$
	if len(allProviders) == 0 {
		return nil, []error{errors.New("no search providers configured")}
	}

	ret := result.Results{}
	var errs []error
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(len(allProviders))
	params.Q = strings.TrimSpace(params.Q)

	for _, p := range allProviders {
		f := p
		go func() {
			res, err := f(ctx, as, params, logger)
			mu.Lock()
			if err != nil {
				errs = append(errs, err)
			}
			ret = append(ret, res...)
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	ret.Sort()
	return ret, errs
}

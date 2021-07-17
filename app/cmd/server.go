// Code generated by Project Forge, see https://projectforge.dev for details.
package cmd

import (
	"fmt"
	"runtime"

	"github.com/fasthttp/router"
	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/controller"
	"github.com/kyleu/admini/app/filesystem"
	"github.com/kyleu/admini/app/util"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

const keyServer = "server"

func serverCmd() *cobra.Command {
	short := fmt.Sprintf("Starts the http server on port %d (by default)", util.AppPort)
	f := func(*cobra.Command, []string) error { return startServer(_flags) }
	ret := &cobra.Command{Use: keyServer, Short: short, RunE: f}
	return ret
}

func startServer(flags *Flags) error {
	err := initIfNeeded()
	if err != nil {
		return errors.Wrap(err, "error initializing application")
	}

	r, _, err := loadServer(flags, _logger)
	if err != nil {
		return err
	}

	_, err = listenandserve(util.AppName, flags.Address, flags.Port, r)
	return err
}

func loadServer(flags *Flags, logger *zap.SugaredLogger) (*router.Router, *zap.SugaredLogger, error) {
	r := controller.AppRoutes()

	f := filesystem.NewFileSystem(flags.ConfigDir, logger)

	st, err := app.NewState(flags.Debug, _buildInfo, r, f, logger)
	if err != nil {
		return nil, logger, err
	}

	svcs, err := app.NewServices(st)
	if err != nil {
		return nil, nil, errors.Wrap(err, "error creating services")
	}
	st.Services = svcs

	controller.SetAppState(st, logger)

	logger.Infof("started %s using address [%s:%d] on %s:%s", util.AppName, flags.Address, flags.Port, runtime.GOOS, runtime.GOARCH)

	return r, logger, nil
}

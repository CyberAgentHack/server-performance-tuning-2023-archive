package v1

import (
	"context"
	"errors"
	stdhttp "net/http"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/gateway/http"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/gateway/http/response"
	v1 "github.com/CyberAgentHack/server-performance-tuning-2023/pkg/gateway/http/v1"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/repository"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/repository/database"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/usecase"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/util/config"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/util/log"
)

type app struct {
	cmd         *cobra.Command
	logger      *zap.Logger
	Level       string `default:"debug"`
	Environment string `default:"local"`
	Port        int    `default:"9000"`
}

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "v1",
		Short: "Server Performance Tuning 2023",
	}
	a := &app{cmd: cmd}
	cmd.RunE = func(_ *cobra.Command, _ []string) error {
		return a.run()
	}
	return cmd
}

func (a *app) run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	// load environment variables
	err := config.Load(a)
	if err != nil {
		return err
	}

	// logger
	a.logger, err = log.NewLogger(a.Level)
	if err != nil {
		return err
	}
	defer a.close(a.logger.Sync)

	// response
	response.SetLogger(a.logger)

	// usecase
	cfg := &usecase.Config{
		DB: &repository.Database{
			Episode: database.NewEpisode(),
			Series:  database.NewSeries(),
			Season:  database.NewSeason(),
		},
	}
	uc := usecase.NewUsecase(cfg)

	// gateway
	params := &http.RunServerParams{
		Group:  &errgroup.Group{},
		Port:   a.Port,
		APIs:   []http.API{v1.NewService(uc, a.logger)},
		Logger: a.logger,
	}
	server, err := http.RunServer(params)
	if err != nil {
		return err
	}
	a.logger.Info("Started server", zap.Any("config", a))

	// shutdown
	<-ctx.Done()
	a.logger.Info("Received signal")
	err = server.Shutdown(context.Background())
	if errors.Is(err, stdhttp.ErrServerClosed) {
		err = nil
	}
	cancel()
	a.logger.Info("Shutdown...")
	return err
}

func (a *app) close(f func() error) {
	if err := f(); err != nil {
		a.logger.Error("Failed to close", zap.Error(err))
	}
}

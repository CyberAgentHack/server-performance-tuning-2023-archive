package app

import (
	"context"
	"errors"
	"net/http"
	"os/signal"
	"syscall"

	apphttp "github.com/CyberAgentHack/server-performance-tuning-2023/pkg/app/http"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/config"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/db"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/log"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/repository"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/repository/database"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/usecase"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

type App struct {
	logger       *zap.Logger
	Level        string `default:"debug"`
	Environment  string `default:"local"`
	Port         int    `default:"9000"`
	DbSecretName string
}

func New() (App, error) {
	const envPrefix = "ENV"
	app := App{}
	err := envconfig.Process(envPrefix, &app)
	return app, err
}

func (a *App) Run() (err error) {
	return a.runWithContext(context.Background())
}

func (a *App) runWithContext(ctx context.Context) (err error) {
	ctx, cancel := signal.NotifyContext(
		ctx,
		syscall.SIGTERM,
		syscall.SIGINT,
	)
	defer cancel()

	group, gCtx := errgroup.WithContext(ctx)

	a.logger, err = log.NewLogger(a.Level)
	if err != nil {
		return err
	}
	defer a.logger.Sync()

	// config
	cfg, err := config.NewConfig(a.Environment, a.DbSecretName)
	if err != nil {
		return err
	}

	mysql, err := db.NewMySQL(cfg.DBConfig)
	if err != nil {
		return err
	}

	// usecase
	database := &repository.Database{
		Episode:        database.NewEpisode(mysql),
		Series:         database.NewSeries(),
		Season:         database.NewSeason(),
		ViewingHistory: database.NewViewingHistory(),
	}
	uc := usecase.NewUsecase(database)

	// run http server
	service := apphttp.NewService(uc, a.logger)
	params := &apphttp.ServerParams{
		Port:    a.Port,
		Logger:  a.logger,
		Service: service,
	}
	server := apphttp.NewServer(params)

	group.Go(func() error {
		return server.ListenAndServe()
	})

	a.logger.Info("Starting server...", zap.Any("config", a))

	// wait signal or error
	<-gCtx.Done()
	err = server.Shutdown(context.Background())
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	a.logger.Info("Shutdown...")
	err = group.Wait()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		a.logger.Info("Received error", zap.Error(err))
		return err
	}

	return nil
}

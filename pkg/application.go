package application

import (
	"github.com/otaviohenrique/go-app-example/pkg/ports"
	"go.uber.org/zap"
	"os"
	"syscall"
)

type Core interface {
	Run()
}

// Application should hold all logic, and "orchestration". Is specially important to do all work not related to API here.
type Application struct {
	api      ports.Api
	consumer ports.Consumer
	producer ports.Producer
	logger   *zap.SugaredLogger
	config   *Config
	metrics  *Metrics
	osSign   chan os.Signal
}

func NewApp(api ports.Api, consumer ports.Consumer, producer ports.Producer, logger *zap.SugaredLogger, config *Config, metrics *Metrics, sig chan os.Signal) *Application {
	app := new(Application)

	app.api = api
	app.consumer = consumer
	app.producer = producer
	app.logger = logger
	app.config = config
	app.metrics = metrics
	app.osSign = sig

	return app
}

func (app *Application) Run() {
	go app.signalHandler()

	app.api.Serve()
}

func (app *Application) signalHandler() {
	for {
		signal := <-app.osSign

		switch signal {
		case syscall.SIGTERM:
			app.logger.Infow("signal received", "signal", "SIGTERM")
			os.Exit(0)
		case syscall.SIGINT:
			app.logger.Infow("signal received", "signal", "SIGINT")
			os.Exit(0)
		default:
			app.logger.Infow("unknown signal received")
		}
	}
}

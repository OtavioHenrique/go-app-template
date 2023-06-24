package cmd

import (
	application "github.com/otaviohenrique/go-app-example/pkg"
	"github.com/otaviohenrique/go-app-example/pkg/adapters"
	"github.com/otaviohenrique/go-app-example/pkg/controller"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run application",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		config := application.NewConfig("go-app-name")
		logger := application.NewLogger()
		metrics := application.NewMetrics()

		consumer := adapters.NewSqsConsumer()
		producer := adapters.NewSqsProducer()

		userController := controller.NewUserController()
		api := adapters.NewHttpApi(userController, logger, config)

		sig := make(chan os.Signal, 1)

		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

		app := application.NewApp(api, consumer, producer, logger, config, metrics, sig)

		app.Run()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}

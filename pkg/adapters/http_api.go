package adapters

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog"
	application "github.com/otaviohenrique/go-app-example/pkg"
	"github.com/otaviohenrique/go-app-example/pkg/controller"
	"go.uber.org/zap"
	"net/http"
)

type HttpApi struct {
	userController *controller.UserController
	logger         *zap.SugaredLogger
	config         *application.Config
}

// HTTP or any other external-world communication adapter should work just as "router" calling the received controller
// Any logic related to schema validation is its responsability too.
func NewHttpApi(userController *controller.UserController, logger *zap.SugaredLogger, config *application.Config) *HttpApi {
	api := new(HttpApi)

	api.userController = userController
	api.logger = logger
	api.config = config

	return api
}

func (api *HttpApi) healthCheck(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (api *HttpApi) getUsers(w http.ResponseWriter, _ *http.Request) {
	resp := api.userController.GetUser()

	w.WriteHeader(200)
	w.Write([]byte(resp))
}

func (api *HttpApi) Serve() {
	r := chi.NewRouter()

	logger := httplog.NewLogger(api.config.AppName, httplog.Options{
		JSON: true,
	})

	r.Use(httplog.RequestLogger(logger))

	r.Get("/healthcheck", api.healthCheck)
	r.Get("/users", api.getUsers)

	port := api.config.HttpPort

	api.logger.Infow("server started", "port", port)
	api.logger.Fatal(http.ListenAndServe(port, r))
}

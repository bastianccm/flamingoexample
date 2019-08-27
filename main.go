package main

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3"
	"flamingo.me/flamingo/v3/core/requestlogger"
	"flamingo.me/flamingo/v3/framework/config"
	"flamingo.me/flamingo/v3/framework/opencensus"
	"flamingo.me/flamingo/v3/framework/web"
)

func main() {
	flamingo.App([]dingo.Module{
		new(requestlogger.Module),
		new(opencensus.Module),
		new(module),
	})
}

type module struct{}

func (*module) DefaultConfig() config.Map {
	return config.Map{
		"opencensus.jaeger.enable": true,
	}
}

func (*module) Configure(injector *dingo.Injector) {
	web.BindRoutes(injector, new(routes))
}

type routes struct {
	controller *controller
}

func (r *routes) Inject(controller *controller) {
	r.controller = controller
}

func (r *routes) Routes(registry *web.RouterRegistry) {
	registry.Route("/", "home")
	registry.HandleAny("home", r.controller.action)
}

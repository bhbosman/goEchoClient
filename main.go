package main

import (
	"context"
	"github.com/bhbosman/gocommon/comms"
	"github.com/bhbosman/gocommon/comms/examples/echo/Server"

	//"github.com/bhbosman/gocommon/comms/examples/echo/Server"
	"github.com/bhbosman/gocommon/stacks"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.LogName("Main Application"),
		//fx.Provide(fx.Annotated{Group: "ClientContextFactory", Target: Server.CreateClientHandlerFactory}),
		fx.Provide(fx.Annotated{Target: comms.NewClientContextFactories}),
		fx.Provide(fx.Annotated{Target: stacks.NewStackFactory}),
		fx.Provide(fx.Annotated{Name: "EchoServerConnectionManager", Target: comms.NewNetListenApp("EchoServerConnectionManager", 3001, "Empty", Server.CreateClientHandlerFactoryName)}),
		fx.Provide(fx.Annotated{Name: "Application", Target: comms.CreateRootContext}),
		fx.Invoke(
			func(params struct {
				fx.In
				Lifecycle                   fx.Lifecycle
				EchoServerConnectionManager *fx.App `name:"EchoServerConnectionManager"`
				Logger                      fx.ILogger
			}) {
				params.Lifecycle.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {
						return params.EchoServerConnectionManager.Start(ctx)
					},
					OnStop: func(ctx context.Context) error {
						return params.EchoServerConnectionManager.Stop(ctx)
					},
				})

			}),
	)
	if app.Err() != nil {
		return
	}
	app.Run()
}

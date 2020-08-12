package main

import (
	"context"
	echoClient "github.com/bhbosman/goEchoClient/internal/components"
	app2 "github.com/bhbosman/gocommon/app"
	"github.com/bhbosman/gocomms/connectionManager"
	"github.com/bhbosman/gocomms/connectionManager/endpoints"
	"github.com/bhbosman/gocomms/connectionManager/view"
	"github.com/bhbosman/gocomms/impl"
	log2 "github.com/bhbosman/gologging"
	"go.uber.org/fx"
	"log"
	"os"
)

func main() {
	app := fx.New(
		log2.ProvideLogFactory(log.New(os.Stderr, "EchoClient: ", log.LstdFlags), nil),
		connectionManager.RegisterDefaultConnectionManager(),
		impl.RegisterAllConnectionRelatedServices(),
		app2.RegisterRootContext(),
		echoClient.RegisterEchoServiceDialer(),
		endpoints.RegisterConnectionManagerEndpoint(),
		view.RegisterConnectionsHtmlTemplate(),
		fx.Provide(
			func(params struct {
				fx.In
				Factory *log2.Factory
			}) *log2.SubSystemLogger {
				return params.Factory.Create("Main")
			}),
		fx.Invoke(
			func(params struct {
				fx.In
				Lifecycle      fx.Lifecycle
				Apps           []*fx.App `group:"Apps"`
				Logger         *log2.SubSystemLogger
			}) {
				for _, item := range params.Apps {
					localApp := item
					params.Lifecycle.Append(fx.Hook{
						OnStart: func(ctx context.Context) error {
							return localApp.Start(ctx)
						},
						OnStop: func(ctx context.Context) error {
							return localApp.Stop(ctx)
						},
					})
				}
			}),
	)
	if app.Err() != nil {
		return
	}
	app.Run()
}

package components

import (
	"github.com/bhbosman/gocomms/impl"
	commsImpl "github.com/bhbosman/gocomms/intf"
	"github.com/bhbosman/gocomms/netDial"
	"go.uber.org/fx"
)

func RegisterEchoServiceDialer() fx.Option {
	const createServerHandlerFactoryName = "EchoClientConnectionReactorFactory"
	return fx.Options(
		fx.Provide(fx.Annotated{
			Group: impl.ConnectionReactorFactoryConst,
			Target: func() (commsImpl.IConnectionReactorFactory, error) {
				return &connectionReactorFactory{
					name: createServerHandlerFactoryName,
				}, nil

			},
		}),
		fx.Provide(fx.Annotated{
			Group: "Apps",
			Target: netDial.NewNetDialApp(
				"EchoServiceDialer(Empty)",
				"tcp4://127.0.0.1:3000",
				impl.TransportFactoryEmptyName,
				impl.CreateEmptyStack,
				createServerHandlerFactoryName,
				netDial.MaxConnectionsSetting(1)),
		}),
		fx.Provide(fx.Annotated{
			Group: "Apps",
			Target: netDial.NewNetDialApp(
				"EchoServiceDialer(Compressed)",
				"tcp4://127.0.0.1:3001",
				impl.TransportFactoryCompressedName,
				impl.CreateCompressedStack,
				createServerHandlerFactoryName,
				netDial.MaxConnectionsSetting(1)),
		}),
		fx.Provide(fx.Annotated{
			Group: "Apps",
			Target: netDial.NewNetDialApp(
				"EchoServiceDialer(UnCompressed)",
				"tcp4://127.0.0.1:3002",
				impl.TransportFactoryUnCompressedName,
				impl.CreateUnCompressedStack,
				createServerHandlerFactoryName,
				netDial.MaxConnectionsSetting(1)),
		}),
	)
}

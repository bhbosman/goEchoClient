package components

import (
	"github.com/bhbosman/gocomms/impl"
	"github.com/bhbosman/gocomms/netDial"
	"go.uber.org/fx"
)

func RegisterEchoServiceDialer() fx.Option {
	const createServerHandlerFactoryName = "EchoClientConnectionReactorFactory"
	cfr := &connectionReactorFactory{
		name: createServerHandlerFactoryName,
	}
	return fx.Options(
		//fx.Provide(fx.Annotated{
		//	Group: impl.ConnectionReactorFactoryConst,
		//	Target: func() (commsImpl.IConnectionReactorFactory, error) {
		//		return cfr, nil
		//	},
		//}),
		fx.Provide(fx.Annotated{
			Group: "Apps",
			Target: netDial.NewNetDialApp(
				"EchoServiceDialer(Empty)",
				"tcp4://127.0.0.1:3000",
				impl.CreateEmptyStack,
				//createServerHandlerFactoryName,
				cfr,
				netDial.MaxConnectionsSetting(1)),
		}),
		fx.Provide(fx.Annotated{
			Group: "Apps",
			Target: netDial.NewNetDialApp(
				"EchoServiceDialer(Compressed)",
				"tcp4://127.0.0.1:3001",
				impl.CreateCompressedStack,
				//createServerHandlerFactoryName,
				cfr,
				netDial.MaxConnectionsSetting(1)),
		}),
		fx.Provide(fx.Annotated{
			Group: "Apps",
			Target: netDial.NewNetDialApp(
				"EchoServiceDialer(UnCompressed)",
				"tcp4://127.0.0.1:3002",
				impl.CreateUnCompressedStack,
				//createServerHandlerFactoryName,
				cfr,
				netDial.MaxConnectionsSetting(1)),
		}),
	)
}

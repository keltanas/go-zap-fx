package zapfx

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

const (
	moduleName = "zap"
)

func Module() fx.Option {
	return fx.Options(
		WithLogger(),
		fx.Module(
			moduleName,
			fx.Provide(
				func() (*zap.Logger, error) {
					return zap.NewProduction()
				},
			),
		),
	)
}

func WithLogger() fx.Option {
	return fx.WithLogger(
		func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		},
	)
}

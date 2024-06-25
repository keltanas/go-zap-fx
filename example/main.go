package main

import (
	"context"

	"go.uber.org/fx"

	zapfx "github.com/keltanas/go-zap-fx"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cnt := fx.New(zapfx.Module())
	if err := cnt.Start(ctx); err != nil {
		panic(err)
	}
	if err := cnt.Stop(ctx); err != nil {
		panic(err)
	}
}

# Module Zap integration to FX

[![Go](https://github.com/keltanas/go-zap-fx/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/keltanas/go-zap-fx/actions/workflows/go.yml)

```
go get github.com/keltanas/go-zap-fx
```

## Example

``` go
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
```

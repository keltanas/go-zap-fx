package zapfx_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/fx"

	zapfx "github.com/keltanas/go-zap-fx"
)

func TestModule(t *testing.T) {
	app := []fx.Option{
		zapfx.Module(),
	}

	err := fx.ValidateApp(app...)
	require.NoError(t, err)
}

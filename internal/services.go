package internal

import (
	"github.com/ankorstore/yokai-worker-template/internal/worker"
	"github.com/ankorstore/yokai/fxworker"
	"go.uber.org/fx"
)

// ProvideServices is used to register the application services.
func ProvideServices() fx.Option {
	return fx.Options(
		fxworker.AsWorker(worker.NewExampleWorker),
	)
}

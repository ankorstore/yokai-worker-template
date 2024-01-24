package internal

import (
	"github.com/ankorstore/yokai-worker-template/internal/worker"
	"github.com/ankorstore/yokai/fxworker"
	"go.uber.org/fx"
)

func ProvideServices() fx.Option {
	return fx.Options(
		fxworker.AsWorker(worker.NewExampleWorker),
	)
}

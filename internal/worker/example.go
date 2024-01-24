package worker

import (
	"context"
	"time"

	"github.com/ankorstore/yokai/config"
	"github.com/ankorstore/yokai/worker"
)

type ExampleWorker struct {
	config *config.Config
}

func NewExampleWorker(config *config.Config) *ExampleWorker {
	return &ExampleWorker{
		config: config,
	}
}

func (w *ExampleWorker) Name() string {
	return "example-worker"
}

func (w *ExampleWorker) Run(ctx context.Context) error {
	logger := worker.CtxLogger(ctx)

	for {
		select {
		case <-ctx.Done():
			logger.Info().Msg("stopping")

			return nil
		default:
			logger.Info().Msg("running")

			time.Sleep(time.Duration(w.config.GetFloat64("config.example-worker.interval")) * time.Second)
		}
	}
}

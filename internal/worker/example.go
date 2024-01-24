package worker

import (
	"context"
	"time"

	"github.com/ankorstore/yokai/config"
	"github.com/ankorstore/yokai/worker"
)

// ExampleWorker is an example worker.
type ExampleWorker struct {
	config *config.Config
}

// NewExampleWorker returns a new [ExampleWorker].
func NewExampleWorker(config *config.Config) *ExampleWorker {
	return &ExampleWorker{
		config: config,
	}
}

// Name returns the [ExampleWorker] name.
func (w *ExampleWorker) Name() string {
	return "example-worker"
}

// Run executes the [ExampleWorker].
func (w *ExampleWorker) Run(ctx context.Context) error {
	logger := worker.CtxLogger(ctx)

	for {
		select {
		case <-ctx.Done():
			logger.Info().Msg("stopping")

			return nil
		default:
			logger.Info().Msg("running")

			// The sleep interval can be configured in the application config files.
			time.Sleep(time.Duration(w.config.GetFloat64("config.example-worker.interval")) * time.Second)
		}
	}
}

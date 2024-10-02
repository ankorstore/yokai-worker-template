package worker_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/ankorstore/yokai-worker-template/internal"
	"github.com/ankorstore/yokai/fxconfig"
	"github.com/ankorstore/yokai/log/logtest"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
)

func TestExampleWorker(t *testing.T) {
	var logBuffer logtest.TestLogBuffer
	var metricsRegistry *prometheus.Registry

	// bootstrap test app
	app := internal.Bootstrapper.BootstrapTestApp(
		t,
		fxconfig.AsConfigPath(fmt.Sprintf("%s/configs/", internal.RootDir)),
		fx.Populate(&logBuffer, &metricsRegistry),
	)

	// start test app
	app.RequireStart()

	// give time to worker to start
	time.Sleep(1 * time.Millisecond)

	// run log assertion
	logtest.AssertHasLogRecord(t, logBuffer, map[string]interface{}{
		"level":   "info",
		"service": "worker-app",
		"module":  "worker",
		"worker":  "example-worker",
		"message": "running",
	})

	// stop test app
	app.RequireStop()

	// give time to worker to stop
	time.Sleep(1 * time.Millisecond)

	// stop log assertion
	logtest.AssertHasLogRecord(t, logBuffer, map[string]interface{}{
		"level":   "info",
		"service": "worker-app",
		"module":  "worker",
		"worker":  "example-worker",
		"message": "stopping",
	})

	// metrics assertion
	expectedMetric := `
		# HELP worker_executions_total Total number of workers executions
		# TYPE worker_executions_total counter
		worker_executions_total{status="started",worker="example_worker"} 1
		worker_executions_total{status="success",worker="example_worker"} 1
	`

	err := testutil.GatherAndCompare(
		metricsRegistry,
		strings.NewReader(expectedMetric),
		"worker_executions_total",
	)
	assert.NoError(t, err)
}

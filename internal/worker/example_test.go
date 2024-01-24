package worker_test

import (
	"strings"
	"testing"

	"github.com/ankorstore/yokai-worker-template/internal"
	"github.com/ankorstore/yokai/log/logtest"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
)

func TestExampleWorker(t *testing.T) {
	var logBuffer logtest.TestLogBuffer
	var metricsRegistry *prometheus.Registry

	internal.RunTest(
		t,
		fx.Populate(
			&logBuffer,
			&metricsRegistry,
		),
	)

	// logs assertion
	logtest.AssertHasLogRecord(t, logBuffer, map[string]interface{}{
		"level":   "info",
		"service": "worker-app",
		"module":  "worker",
		"worker":  "example-worker",
		"message": "running",
	})

	logtest.AssertHasLogRecord(t, logBuffer, map[string]interface{}{
		"level":   "info",
		"service": "worker-app",
		"module":  "worker",
		"worker":  "example-worker",
		"message": "stopping",
	})

	// metrics assertion
	expectedMetric := `
		# HELP worker_app_worker_worker_execution_total Total number of workers executions
		# TYPE worker_app_worker_worker_execution_total counter
		worker_app_worker_worker_execution_total{status="started",worker="example_worker"} 1
		worker_app_worker_worker_execution_total{status="success",worker="example_worker"} 1
	`

	err := testutil.GatherAndCompare(
		metricsRegistry,
		strings.NewReader(expectedMetric),
		"worker_app_worker_worker_execution_total",
	)
	assert.NoError(t, err)
}

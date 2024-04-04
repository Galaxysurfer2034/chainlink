package plugins

import (
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/config"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/store/dialects"
)

func TestPluginPortManager(t *testing.T) {
	// register one
	m := NewLoopRegistry(logger.TestLogger(t), nil, nil)
	pFoo, err := m.Register("foo")
	require.NoError(t, err)
	require.Equal(t, "foo", pFoo.Name)
	require.Greater(t, pFoo.EnvCfg.PrometheusPort, 0)
	// test duplicate
	pNil, err := m.Register("foo")
	require.ErrorIs(t, err, ErrExists)
	require.Nil(t, pNil)
	// ensure increasing port assignment
	pBar, err := m.Register("bar")
	require.NoError(t, err)
	require.Equal(t, "bar", pBar.Name)
	require.Equal(t, pFoo.EnvCfg.PrometheusPort+1, pBar.EnvCfg.PrometheusPort)
}

// Mock tracing config
type MockCfgTracing struct{}

func (m *MockCfgTracing) Attributes() map[string]string {
	return map[string]string{"attribute": "value"}
}
func (m *MockCfgTracing) Enabled() bool           { return true }
func (m *MockCfgTracing) NodeID() string          { return "" }
func (m *MockCfgTracing) CollectorTarget() string { return "http://localhost:9000" }
func (m *MockCfgTracing) SamplingRatio() float64  { return 0.1 }
func (m *MockCfgTracing) TLSCertPath() string     { return "/path/to/cert.pem" }
func (m *MockCfgTracing) Mode() string            { return "tls" }

type MockCfgDatabase struct{}

func (m MockCfgDatabase) Backup() config.Backup { panic("unimplemented") }

func (m MockCfgDatabase) Listener() config.Listener { panic("unimplemented") }

func (m MockCfgDatabase) Lock() config.Lock { panic("unimplemented") }

func (m MockCfgDatabase) DefaultIdleInTxSessionTimeout() time.Duration { return time.Hour }

func (m MockCfgDatabase) DefaultLockTimeout() time.Duration { return time.Minute }

func (m MockCfgDatabase) DefaultQueryTimeout() time.Duration { return time.Second }

func (m MockCfgDatabase) Dialect() dialects.DialectName { panic("unimplemented") }

func (m MockCfgDatabase) LogSQL() bool { return true }

func (m MockCfgDatabase) MaxIdleConns() int { return 99 }

func (m MockCfgDatabase) MaxOpenConns() int { return 42 }

func (m MockCfgDatabase) MigrateDatabase() bool { panic("unimplemented") }

func (m MockCfgDatabase) URL() url.URL {
	return url.URL{Scheme: "fake", Host: "database.url"}
}

func TestLoopRegistry_Register(t *testing.T) {
	mockCfgTracing := &MockCfgTracing{}
	mockCfgDatabase := &MockCfgDatabase{}
	registry := make(map[string]*RegisteredLoop)

	// Create a LoopRegistry instance with mockCfgTracing
	loopRegistry := &LoopRegistry{
		lggr:        logger.TestLogger(t),
		registry:    registry,
		cfgTracing:  mockCfgTracing,
		cfgDatabase: mockCfgDatabase,
	}

	// Test case 1: Register new loop
	registeredLoop, err := loopRegistry.Register("testID")
	require.Nil(t, err)
	require.Equal(t, "testID", registeredLoop.Name)

	cfg := registeredLoop.EnvCfg
	require.True(t, cfg.TracingEnabled)
	require.Equal(t, "http://localhost:9000", cfg.TracingCollectorTarget)
	require.Equal(t, map[string]string{"attribute": "value"}, cfg.TracingAttributes)
	require.Equal(t, 0.1, cfg.TracingSamplingRatio)
	require.Equal(t, "/path/to/cert.pem", cfg.TracingTLSCertPath)

	require.Equal(t, "fake://database.url", cfg.DatabaseURL)
	require.Equal(t, time.Hour, cfg.DatabaseIdleInTxSessionTimeout)
	require.Equal(t, time.Minute, cfg.DatabaseLockTimeout)
	require.Equal(t, time.Second, cfg.DatabaseQueryTimeout)
	require.Equal(t, true, cfg.DatabaseLogSQL)
	require.Equal(t, 99, cfg.DatabaseMaxOpenConns)
	require.Equal(t, 42, cfg.DatabaseMaxIdleConns)
}

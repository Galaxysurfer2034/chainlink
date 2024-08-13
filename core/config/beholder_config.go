package config

type Telemetry interface {
	InsecureConnection() bool
	CACertFile() string
	OtelExporterGRPCEndpoint() string
	ResourceAttributes() map[string]string
	TraceSampleRatio() float64
}

package metrics

type Metrics struct{}

// This class should take care of all logic related to metrics. If needed, it can receive config, or other port needed
func NewMetrics() *Metrics {
	return new(Metrics)
}

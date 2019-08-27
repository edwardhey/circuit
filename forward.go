package circuit

import (
	"github.com/edwardhey/circuit/v3"
)

// generated by goforward github.com/edwardhey/circuit/v3 github.com/edwardhey/circuit

// Circuit is a circuit breaker pattern implementation that can accept commands and open/close on failures
type Circuit = circuit.Circuit

// NewCircuitFromConfig creates an inline circuit.  If you want to group all your circuits together, you should probably
// just use Manager struct instead.
func NewCircuitFromConfig(name string, config Config) *Circuit {
	return circuit.NewCircuitFromConfig(name, config)
}

// ClosedToOpen receives events and controls if the circuit should open or close as a result of those events.
// Return true if the circuit should open, false if the circuit should close.
type ClosedToOpen = circuit.ClosedToOpen

// OpenToClosed controls logic that tries to close an open circuit
type OpenToClosed = circuit.OpenToClosed

// Config controls how a circuit operates
type Config = circuit.Config

// GeneralConfig controls the non general logic of the circuit.  Things specific to metrics, execution, or fallback are
// in their own configs
type GeneralConfig = circuit.GeneralConfig

// ExecutionConfig is https://github.com/Netflix/Hystrix/wiki/Configuration#execution
type ExecutionConfig = circuit.ExecutionConfig

// FallbackConfig is https://github.com/Netflix/Hystrix/wiki/Configuration#fallback
type FallbackConfig = circuit.FallbackConfig

// MetricsCollectors can receive metrics during a circuit.  They should be fast, as they will
// block circuit operation during function calls.
type MetricsCollectors = circuit.MetricsCollectors

// TimeKeeper allows overriding time to test the circuit
type TimeKeeper = circuit.TimeKeeper

// Configurable is anything that can receive configuration changes while live
type Configurable = circuit.Configurable

// BadRequest is implemented by an error returned by runFunc if you want to consider the requestor bad, not the circuit
// bad.  See http://netflix.github.io/Hystrix/javadoc/com/netflix/hystrix/exception/HystrixBadRequestException.html
// and https://github.com/Netflix/Hystrix/wiki/How-To-Use#error-propagation for information.
type BadRequest = circuit.BadRequest

// IsBadRequest returns true if the error is of type BadRequest
func IsBadRequest(err error) bool {
	return circuit.IsBadRequest(err)
}

// SimpleBadRequest is a simple wrapper for an error to mark it as a bad request
type SimpleBadRequest = circuit.SimpleBadRequest

// CommandPropertiesConstructor is a generic function that can create command properties to configure a circuit by name
// It is safe to leave not configured properties their empty value.
type CommandPropertiesConstructor = circuit.CommandPropertiesConstructor

// Manager manages circuits with unique names
type Manager = circuit.Manager

// RunMetricsCollection send metrics to multiple RunMetrics
type RunMetricsCollection = circuit.RunMetricsCollection

// FallbackMetricsCollection sends fallback metrics to all collectors
type FallbackMetricsCollection = circuit.FallbackMetricsCollection

// MetricsCollection allows reporting multiple circuit metrics at once
type MetricsCollection = circuit.MetricsCollection

// Metrics reports internal circuit metric events
type Metrics = circuit.Metrics

// RunMetrics is guaranteed to execute one (and only one) of the following functions each time the circuit
// attempts to call a run function. Methods with durations are when run was actually executed.  Methods without
// durations never called run, probably because of the circuit.
type RunMetrics = circuit.RunMetrics

// FallbackMetrics is guaranteed to execute one (and only one) of the following functions each time a fallback is executed.
// Methods with durations are when the fallback is actually executed.  Methods without durations are when the fallback was
// never called, probably because of some circuit condition.
type FallbackMetrics = circuit.FallbackMetrics

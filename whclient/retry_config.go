package whclient

import (
	"math/rand"
	"time"
)

const (
	defaultInitialDelay        = 500 * time.Millisecond
	defaultMaxDelay            = 60 * time.Second
	defaultMaxElapsedTime      = 3 * time.Minute
	defaultMultiplier          = 1.5
	defaultRandomizationFactor = 0.5
)

// RetryConfig contains exponential backoff parameters for retrying connections
type RetryConfig struct {
	// Retry values
	InitialDelay        time.Duration // Default = 500 * time.Millisecond
	MaxDelay            time.Duration // Default = 60 * time.Second
	MaxElapsedTime      time.Duration // Default = 3 * time.Minute
	Multiplier          float64       // Default = 1.5
	RandomizationFactor float64       // Default = 0.5
}

// NextDelay calculates the next interval based on the current interval
func (r RetryConfig) NextDelay(currentDelay time.Duration) time.Duration {
	// check if current interval is max interval
	// avoid calculation
	if currentDelay == r.MaxDelay {
		return currentDelay
	}

	delta := r.RandomizationFactor * float64(currentDelay)
	minDelay := r.Multiplier*float64(currentDelay) - delta
	maxDelay := r.Multiplier*float64(currentDelay) + delta
	nextDelay := minDelay + (rand.Float64() * (maxDelay - minDelay + 1))
	Delay := time.Duration(nextDelay)
	if Delay > r.MaxDelay {
		Delay = r.MaxDelay
	}
	return Delay
}

// initializeRetryValues sets the RetryConfig parameteres to their
// default value
func (r RetryConfig) defaultValues() RetryConfig {
	conf := RetryConfig{
		InitialDelay:        r.InitialDelay,
		MaxDelay:            r.MaxDelay,
		MaxElapsedTime:      r.MaxElapsedTime,
		Multiplier:          r.Multiplier,
		RandomizationFactor: r.RandomizationFactor,
	}

	if r.InitialDelay == 0 {
		conf.InitialDelay = defaultInitialDelay
	}
	if r.MaxDelay == 0 {
		conf.MaxDelay = defaultMaxDelay
	}
	if r.MaxElapsedTime == 0 {
		conf.MaxElapsedTime = defaultMaxElapsedTime
	}

	if r.Multiplier < 1.0 {
		conf.Multiplier = defaultMultiplier
	}

	if r.RandomizationFactor == 0 {
		conf.RandomizationFactor = defaultRandomizationFactor
	}

	return conf
}

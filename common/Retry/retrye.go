package Retry

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"govenv/pkg/common/MyLogger"
)

// ####################################################################
// ##################### Retry for coroutines #########################
type RetryOptions struct {
	Logger    *MyLogger.MyLogger // optional logger
	Name      string
	Tries     int                          // Number of attempts (-1 for infinite)
	Delay     time.Duration                // Initial delay between attempts
	MaxDelay  time.Duration                // Maximum delay allowed
	Backoff   float64                      // Multiplier applied to the delay
	Jitter    time.Duration                // Fixed jitter or random range jitter
	OnRetry   func(attempt int, err error) // Hook before each retry
	OnFailure func(err error)              // Hook after exhausting retries
}

func NewSocketRetryer(logger *MyLogger.MyLogger, name string, tries int, delay time.Duration, maxDelay time.Duration, backOff float64, jitter time.Duration,
	onRetry func(int, error), onFailure func(error),
) *RetryOptions {
	return &RetryOptions{
		Logger:   logger,
		Name:     name,
		Tries:    tries,
		Delay:    delay,
		MaxDelay: maxDelay,
		Backoff:  backOff,
		Jitter:   jitter,
		OnRetry: func(attempt int, err error) {
			fmt.Printf("Retry %d failed: %v\n", attempt, err)
		},
		OnFailure: func(err error) {
			fmt.Printf("Task failed after all retries: %v\n", err)
		},
	}
}

func Retry(ctx context.Context, fn func() error, opts RetryOptions) error {
	if opts.Tries == 0 {
		return nil // No retries configured
	}
	triesLeft := opts.Tries
	delay := opts.Delay

	for triesLeft != 0 {
		err := fn()
		if err == nil {
			return nil // Success
		}

		if opts.OnRetry != nil {
			opts.OnRetry(opts.Tries-triesLeft+1, err)
		}

		triesLeft--
		if triesLeft == 0 && opts.Tries > 0 {
			if opts.OnFailure != nil {
				opts.OnFailure(err)
			}
			return fmt.Errorf("retry failed after %d attempts: %w", opts.Tries, err)
		}

		time.Sleep(addJitter(delay, opts.Jitter))
		delay = time.Duration(float64(delay) * opts.Backoff)
		if opts.MaxDelay > 0 && delay > opts.MaxDelay {
			delay = opts.MaxDelay
		}
	}
	return nil
}

func addJitter(delay time.Duration, jitter time.Duration) time.Duration {
	if jitter <= 0 {
		return delay
	}
	randomJitter := time.Duration(rand.Int63n(int64(jitter)))
	return delay + randomJitter
}

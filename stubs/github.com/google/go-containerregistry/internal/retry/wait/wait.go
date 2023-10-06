package wait

import "time"

type ConditionFunc func() (done bool, err error)

type Backoff struct {
	Duration time.Duration
	Factor   float64
	Jitter   float64
	Steps    int
	Cap      time.Duration
}

func Jitter(duration time.Duration, maxFactor float64) time.Duration {
	panic("stub")
}

func (b *Backoff) Step() time.Duration {
	panic("stub")
}

func ExponentialBackoff(backoff Backoff, condition ConditionFunc) error {
	panic("stub")
}

type Embedme interface{}

package pattern

import (
	"errors"
	"time"
)

/**
 * created: 2019/6/14 15:56
 * By Will Fan
 */

var (
	ErrNoTickets		= errors.New("semaphore: could not acquire semaphore")
	ErrIllegalRelease	= errors.New("semaphore: can't release the semaphore without acquiring it first")
)

type Interface interface {
	Acquire() error
	Release() error
}

type implementation struct {
	sem			chan struct{}
	timeout		time.Duration
}

func (s *implementation) Acquire() error  {
	select {
	case s.sem <- struct{}{}:
		return nil
	case <-time.After(s.timeout):
		return ErrNoTickets
	}
}

func (s *implementation) Release() error {
	select {
	case _ = <-s.sem:
		return nil
	case <-time.After(s.timeout):
		return ErrIllegalRelease
	}

	//return nil
}

func New(tickets int, timeout time.Duration) Interface {
	return &implementation{
		sem:		make(chan struct{}, tickets),
		timeout:	timeout,
	}
}
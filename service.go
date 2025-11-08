package subsys

import (
	"github.com/Station-Manager/errors"
	"sync"
	"sync/atomic"
)

type Service struct {
	isInitialized atomic.Bool
	isStarted     atomic.Bool
	initOnce      sync.Once
	initErr       error
	mu            sync.Mutex
}

// Initialize ensures the service is initialized and ready for use. It is idempotent and thread-safe.
func (s *Service) Initialize() error {
	const op errors.Op = "subsys.Service.Initialize"
	if s == nil {
		return errors.New(op).Msg(errMsgNilService)
	}

	if s.isInitialized.Load() {
		return nil // Exit gracefully
	}

	s.initOnce.Do(func() {

		s.isInitialized.Store(true)
	})

	return s.initErr
}

// Start starts the service. This is a blocking call and is not idempotent, if there is an issue starting the subsystem,
// it will return an error.
func (s *Service) Start() error {
	const op errors.Op = "subsys.Service.Start"
	if s == nil {
		return errors.New(op).Msg(errMsgNilService)
	}

	if !s.isInitialized.Load() {
		return errors.New(op).Msg(errMsgNotInitialized)
	}

	if s.isStarted.Load() {
		return errors.New(op).Msg(errMsgAlreadyStarted)
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if s.isStarted.Load() {
		return errors.New(op).Msg(errMsgAlreadyStarted)
	}

	s.isStarted.Store(true)

	return nil
}

// Stop stops the service. This is a blocking call and is not idempotent, if there is an issue stopping the subsystem,
// it will return an error.
func (s *Service) Stop() error {
	const op errors.Op = "subsys.Service.Stop"
	if s == nil {
		return errors.New(op).Msg(errMsgNilService)
	}

	if !s.isInitialized.Load() {
		return errors.New(op).Msg(errMsgNotInitialized)
	}

	if !s.isStarted.Load() {
		return errors.New(op).Msg(errMsgAlreadyStopped)
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.isStarted.Load() {
		return errors.New(op).Msg(errMsgAlreadyStopped)
	}

	s.isStarted.Store(false)

	return nil
}

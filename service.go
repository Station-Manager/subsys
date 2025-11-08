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
		// Perform initialization and capture any error.
		// Replace the following with real initialization logic.
		s.initErr = func() error {
			// Do some initialization here
			// If there is an error, store it in s.initErr and return
			return nil
		}()

		// Only set isInitialized to true if there was no error during initialization
		if s.initErr == nil {
			s.isInitialized.Store(true)
		}
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

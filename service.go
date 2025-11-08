package subsys

import (
	"github.com/Station-Manager/errors"
	"sync/atomic"
)

type Service struct {
	isInitialized atomic.Bool
}

func (s *Service) Initialize() error {
	const op errors.Op = "subsys.Service.Initialize"
	if s == nil {
		return errors.New(op).Msg(errMsgNilService)
	}

	if s.isInitialized.Load() {
		return nil // Exit gracefully
	}

	s.isInitialized.Store(true)

	return nil
}

func (s *Service) Start() error {
	return nil
}

func (s *Service) Stop() error {
	return nil
}

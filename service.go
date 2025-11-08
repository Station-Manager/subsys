package subsys

import "sync/atomic"

type Service struct {
	inIntialized atomic.Bool
}

func (s *Service) Initialize() error {
	return nil
}

package service

import (
	"github.com/hung/tdd/foo"
)

// Service handles business logic using the Processor
type Service struct {
	processor foo.Processor
}

// NewService creates a new Service instance
func NewService(processor foo.Processor) *Service {
	return &Service{
		processor: processor,
	}
}

// ProcessSync calls Func1 synchronously
func (s *Service) ProcessSync(value int) error {
	return s.processor.Func1(value)
}

// ProcessAsync calls Func1 in a goroutine (fire-and-forget)
func (s *Service) ProcessAsync(value int) {
	go func() {
		_ = s.processor.Func1(value)
	}()

	go func() {
		_ = s.processor.Func2(value)
	}()
}

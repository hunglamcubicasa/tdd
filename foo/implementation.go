package foo

import "fmt"

// ProcessorImpl is the concrete implementation of Processor
type ProcessorImpl struct{}

// NewProcessor creates a new ProcessorImpl instance
func NewProcessor() *ProcessorImpl {
	return &ProcessorImpl{}
}

// Func1 processes the given value
func (p *ProcessorImpl) Func1(value int) error {
	if value < 0 {
		return fmt.Errorf("invalid value: %d", value)
	}
	// Simulate some processing
	return nil
}

// Func2 processes the given value
func (p *ProcessorImpl) Func2(value int) error {
	if value < 0 {
		return fmt.Errorf("invalid value: %d", value)
	}
	// Simulate some processing
	return nil
}

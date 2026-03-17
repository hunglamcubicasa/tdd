package foo

// Processor defines the interface for processing operations
type Processor interface {
	Func1(value int) error
	Func2(value int) error
}

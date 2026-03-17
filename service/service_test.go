package service

import (
	"errors"
	"sync"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	foo "github.com/hung/tdd/mocks"
)

// BaseTestSuite contains shared test data and setup logic
type BaseTestSuite struct {
	testValue    int
	testError    error
	expectedData map[string]int
}

// setupTestData initializes common test data
func (b *BaseTestSuite) setupTestData() {
	b.testValue = 42
	b.testError = errors.New("test error")
	b.expectedData = map[string]int{
		"success": 100,
		"failure": -1,
	}
}

// ProcessSyncSuite tests the ProcessSync function
type ProcessSyncSuite struct {
	suite.Suite
	BaseTestSuite
	mockProcessor *foo.MockProcessor
	service       *Service
}

// SetupSuite is called once before the suite runs
func (s *ProcessSyncSuite) SetupSuite() {
	s.setupTestData()
}

// SetupTest is called before each test
func (s *ProcessSyncSuite) SetupTest() {
	s.mockProcessor = foo.NewMockProcessor(s.T())
	s.service = NewService(s.mockProcessor)
}

// TearDownTest is called after each test
func (s *ProcessSyncSuite) TearDownTest() {
	s.mockProcessor = nil
	s.service = nil
}

// TestProcessSync_Success tests successful synchronous processing
func (s *ProcessSyncSuite) TestProcessSync_Success() {
	// Arrange
	s.mockProcessor.EXPECT().Func1(s.testValue).Return(nil).Once()

	// Act
	err := s.service.ProcessSync(s.testValue)

	// Assert
	s.NoError(err)
}

// TestProcessSync_Error tests error handling in synchronous processing
func (s *ProcessSyncSuite) TestProcessSync_Error() {
	// Arrange
	s.mockProcessor.EXPECT().Func1(s.testValue).Return(s.testError).Once()

	// Act
	err := s.service.ProcessSync(s.testValue)

	// Assert
	s.Error(err)
	s.Equal(s.testError, err)
}

// TestProcessSync_WithNegativeValue tests processing with negative value
func (s *ProcessSyncSuite) TestProcessSync_WithNegativeValue() {
	// Arrange
	negativeValue := s.expectedData["failure"]
	expectedErr := errors.New("invalid value")
	s.mockProcessor.EXPECT().Func1(negativeValue).Return(expectedErr).Once()

	// Act
	err := s.service.ProcessSync(negativeValue)

	// Assert
	s.Error(err)
	s.Equal(expectedErr, err)
}

// TestProcessSync_WithSuccessValue tests processing with success value
func (s *ProcessSyncSuite) TestProcessSync_WithSuccessValue() {
	// Arrange
	successValue := s.expectedData["success"]
	s.mockProcessor.EXPECT().Func1(successValue).Return(nil).Once()

	// Act
	err := s.service.ProcessSync(successValue)

	// Assert
	s.NoError(err)
}

// ProcessAsyncSuite tests the ProcessAsync function
type ProcessAsyncSuite struct {
	suite.Suite
	BaseTestSuite
	mockProcessor *foo.MockProcessor
	service       *Service
}

// SetupSuite is called once before the suite runs
func (s *ProcessAsyncSuite) SetupSuite() {
	s.setupTestData()
}

// SetupTest is called before each test
func (s *ProcessAsyncSuite) SetupTest() {
	s.mockProcessor = foo.NewMockProcessor(s.T())
	s.service = NewService(s.mockProcessor)
}

// TearDownTest is called after each test
func (s *ProcessAsyncSuite) TearDownTest() {
	s.mockProcessor = nil
	s.service = nil
}

// TestProcessAsync_Success tests successful asynchronous processing
func (s *ProcessAsyncSuite) TestProcessAsync_Success() {
	// Arrange
	var wg sync.WaitGroup
	wg.Add(2)

	s.mockProcessor.EXPECT().Func1(s.testValue).Return(nil).Once().Run(func(args mock.Arguments) {
		wg.Done()
	})
	s.mockProcessor.EXPECT().Func2(s.testValue).Return(nil).Once().Run(func(args mock.Arguments) {
		wg.Done()
	})

	// Act
	s.service.ProcessAsync(s.testValue)

	// Assert
	wg.Wait()
}

// TestProcessAsync_Error tests error handling in asynchronous processing
func (s *ProcessAsyncSuite) TestProcessAsync_Error() {
	// Arrange
	var wg sync.WaitGroup
	wg.Add(2)
	s.mockProcessor.EXPECT().Func1(s.testValue).Return(s.testError).Once().Run(func(args mock.Arguments) {
		wg.Done()
	})
	s.mockProcessor.EXPECT().Func2(s.testValue).Return(s.testError).Once().Run(func(args mock.Arguments) {
		wg.Done()
	})

	// Act
	s.service.ProcessAsync(s.testValue)

	// Assert
	wg.Wait()
}

// TestProcessAsync_WithNegativeValue tests async processing with negative value
func (s *ProcessAsyncSuite) TestProcessAsync_WithNegativeValue() {
	// Arrange
	negativeValue := s.expectedData["failure"]
	var wg sync.WaitGroup
	wg.Add(2)
	expectedErr := errors.New("invalid value")
	s.mockProcessor.EXPECT().Func1(negativeValue).Return(expectedErr).Once().Run(func(args mock.Arguments) {
		wg.Done()
	})
	s.mockProcessor.EXPECT().Func2(negativeValue).Return(expectedErr).Once().Run(func(args mock.Arguments) {
		wg.Done()
	})

	// Act
	s.service.ProcessAsync(negativeValue)

	// Assert
	wg.Wait()
}

// TestProcessAsync_WithSuccessValue tests async processing with success value
func (s *ProcessAsyncSuite) TestProcessAsync_WithSuccessValue() {
	// Arrange
	successValue := s.expectedData["success"]
	var wg sync.WaitGroup
	wg.Add(2)
	s.mockProcessor.EXPECT().Func1(successValue).Return(nil).Once().Run(func(args mock.Arguments) {
		wg.Done()
	})
	s.mockProcessor.EXPECT().Func2(successValue).Return(nil).Once().Run(func(args mock.Arguments) {
		wg.Done()
	})

	// Act
	s.service.ProcessAsync(successValue)

	// Assert
	wg.Wait()
}

// TestProcessSyncSuite runs the ProcessSyncSuite
func TestProcessSyncSuite(t *testing.T) {
	suite.Run(t, new(ProcessSyncSuite))
}

// TestProcessAsyncSuite runs the ProcessAsyncSuite
func TestProcessAsyncSuite(t *testing.T) {
	suite.Run(t, new(ProcessAsyncSuite))
}

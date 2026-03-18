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

func (s *ProcessSyncSuite) TestProcessSync_SuccessCases() {
	testCases := []struct {
		name  string
		value int
	}{
		{name: "default value", value: s.testValue},
		{name: "success value", value: s.expectedData["success"]},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.mockProcessor = foo.NewMockProcessor(s.T())
			s.service = NewService(s.mockProcessor)

			s.mockProcessor.EXPECT().Func1(tc.value).Return(nil).Once()

			err := s.service.ProcessSync(tc.value)

			s.NoError(err)
		})
	}
}

func (s *ProcessSyncSuite) TestProcessSync_FailureCases() {
	testCases := []struct {
		name        string
		value       int
		returnError error
	}{
		{name: "processor returns error", value: s.testValue, returnError: s.testError},
		{name: "negative value", value: s.expectedData["failure"], returnError: errors.New("invalid value")},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.mockProcessor = foo.NewMockProcessor(s.T())
			s.service = NewService(s.mockProcessor)

			s.mockProcessor.EXPECT().Func1(tc.value).Return(tc.returnError).Once()

			err := s.service.ProcessSync(tc.value)

			s.Error(err)
			s.Equal(tc.returnError, err)
		})
	}
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

func (s *ProcessAsyncSuite) TestProcessAsync_SuccessCases() {
	testCases := []struct {
		name  string
		value int
	}{
		{name: "default value", value: s.testValue},
		{name: "success value", value: s.expectedData["success"]},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.mockProcessor = foo.NewMockProcessor(s.T())
			s.service = NewService(s.mockProcessor)

			var wg sync.WaitGroup
			wg.Add(2)

			s.mockProcessor.EXPECT().Func1(tc.value).Return(nil).Once().Run(func(args mock.Arguments) { wg.Done() })
			s.mockProcessor.EXPECT().Func2(tc.value).Return(nil).Once().Run(func(args mock.Arguments) { wg.Done() })

			s.service.ProcessAsync(tc.value)

			wg.Wait()
		})
	}
}

func (s *ProcessAsyncSuite) TestProcessAsync_FailureCases() {
	testCases := []struct {
		name        string
		value       int
		returnError error
	}{
		{name: "processor returns error", value: s.testValue, returnError: s.testError},
		{name: "negative value", value: s.expectedData["failure"], returnError: errors.New("invalid value")},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.mockProcessor = foo.NewMockProcessor(s.T())
			s.service = NewService(s.mockProcessor)

			var wg sync.WaitGroup
			wg.Add(2)

			s.mockProcessor.EXPECT().Func1(tc.value).Return(tc.returnError).Once().Run(func(args mock.Arguments) { wg.Done() })
			s.mockProcessor.EXPECT().Func2(tc.value).Return(tc.returnError).Once().Run(func(args mock.Arguments) { wg.Done() })

			s.service.ProcessAsync(tc.value)

			wg.Wait()
		})
	}
}

// TestProcessSyncSuite runs the ProcessSyncSuite
func TestProcessSyncSuite(t *testing.T) {
	suite.Run(t, new(ProcessSyncSuite))
}

// TestProcessAsyncSuite runs the ProcessAsyncSuite
func TestProcessAsyncSuite(t *testing.T) {
	suite.Run(t, new(ProcessAsyncSuite))
}

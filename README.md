# TDD with Mockery Example

This project demonstrates Test-Driven Development (TDD) in Go using mockery for generating mocks.

## Project Structure

```
.
├── foo/
│   ├── interface.go        # Processor interface with Func1(int) error
│   └── implementation.go   # ProcessorImpl concrete implementation
├── service/
│   ├── service.go          # Service with sync and async methods
│   └── service_test.go     # Test suites for the service
├── mocks/
│   └── mock_Processor.go   # Generated mock for Processor interface
├── .mockery.yaml           # Mockery configuration
└── go.mod
```

## Features

### Package `foo`
- **Interface**: `Processor` with method `Func1(int) error`
- **Implementation**: `ProcessorImpl` that implements the Processor interface

### Package `service`
- **Service**: Contains two methods:
  - `ProcessSync(int) error` - Calls `Func1` synchronously
  - `ProcessAsync(int)` - Calls `Func1` in a goroutine (fire-and-forget)

### Test Architecture

The test suite follows a shared base pattern:

#### BaseTestSuite
- Contains common test data (`testValue`, `testError`, `expectedData`)
- Provides `setupTestData()` method to initialize shared test data
- Both test suites embed this base structure

#### ProcessSyncSuite
- Tests the `ProcessSync` method
- Inherits from `BaseTestSuite`
- Calls `setupTestData()` in `SetupSuite`
- Test cases:
  - Success scenario
  - Error handling
  - Negative value handling
  - Success value from test data

#### ProcessAsyncSuite
- Tests the `ProcessAsync` method
- Inherits from `BaseTestSuite`
- Calls `setupTestData()` in `SetupSuite`
- Uses channels and timeouts to verify async execution
- Test cases:
  - Success scenario
  - Error handling (fire-and-forget)
  - Negative value handling
  - Success value from test data

## Running Tests

### Run all tests
```bash
go test -v ./service/...
```

### Run only ProcessSyncSuite
```bash
go test -v ./service/... -run TestProcessSyncSuite
```

### Run only ProcessAsyncSuite
```bash
go test -v ./service/... -run TestProcessAsyncSuite
```

### Run a specific test
```bash
go test -v ./service/... -run TestProcessSyncSuite/TestProcessSync_Success
```

## Generating Mocks

Mocks are generated using mockery. To regenerate mocks:

```bash
~/go/bin/mockery
```

Configuration is in `.mockery.yaml`.

## Dependencies

- `github.com/stretchr/testify` - Testing framework with suite support
- `github.com/vektra/mockery/v2` - Mock generation tool

## Key TDD Principles Demonstrated

1. **Test Suites**: Using testify's suite package for organized tests
2. **Shared Test Data**: `BaseTestSuite` with common test data and setup
3. **Mock Generation**: Using mockery for interface mocking
4. **Independent Tests**: Each suite can run independently
5. **Async Testing**: Proper testing of goroutines with channels and timeouts
6. **Arrange-Act-Assert**: Clear test structure in all test cases

package subsys

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TestSuite struct {
	suite.Suite
}

func TestSubSys(T *testing.T) {
	suite.Run(T, new(TestSuite))
}

func (t *TestSuite) TestConcurrentStart() {
	service := &Service{}

	const goroutines = 100
	done := make(chan error, goroutines)

	// Launch concurrent Initialize and Start calls
	for i := 0; i < goroutines; i++ {
		go func() {
			err := service.Initialize()
			require.NoError(t.T(), err)
			done <- service.Start()
		}()
	}

	// Collect results - exactly one should succeed, others should fail
	var successCount, errorCount int
	for i := 0; i < goroutines; i++ {
		if err := <-done; err == nil {
			successCount++
		} else {
			errorCount++
		}
	}

	// Verify only one goroutine successfully started the service
	assert.Equal(t.T(), 1, successCount, "Expected exactly one successful Start()")
	assert.Equal(t.T(), goroutines-1, errorCount, "Expected all other Start() calls to fail")

	// Verify service is in started state
	assert.True(t.T(), service.isStarted.Load())

	// Stop the service
	err := service.Stop()
	assert.NoError(t.T(), err)
	assert.False(t.T(), service.isStarted.Load())
}

func (t *TestSuite) TestConcurrentStop() {
	service := &Service{}

	// Initialize and start the service
	err := service.Initialize()
	require.NoError(t.T(), err)
	err = service.Start()
	require.NoError(t.T(), err)

	const goroutines = 100
	done := make(chan error, goroutines)

	// Launch concurrent Stop calls
	for i := 0; i < goroutines; i++ {
		go func() {
			done <- service.Stop()
		}()
	}

	// Collect results - exactly one should succeed, others should fail
	var successCount, errorCount int
	for i := 0; i < goroutines; i++ {
		if err := <-done; err == nil {
			successCount++
		} else {
			errorCount++
		}
	}

	// Verify only one goroutine successfully stopped the service
	assert.Equal(t.T(), 1, successCount, "Expected exactly one successful Stop()")
	assert.Equal(t.T(), goroutines-1, errorCount, "Expected all other Stop() calls to fail")

	// Verify service is in stopped state
	assert.False(t.T(), service.isStarted.Load())
}

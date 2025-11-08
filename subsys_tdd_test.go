package subsys

import (
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

func (t *TestSuite) TestInitializeConcurrent() {

	service := &Service{}

	const goroutines = 100
	done := make(chan bool, goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			err := service.Initialize()
			require.NoError(t.T(), err)
			done <- true
		}()
	}

	for i := 0; i < goroutines; i++ {
		<-done
	}
}

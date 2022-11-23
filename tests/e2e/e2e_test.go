package e2e

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type RunChainSuite struct {
	BaseSetupSuite
}

func (s *RunChainSuite) SetupSuite() {
	s.BaseSetupSuite.SetupSuite(false, false)
}

func TestRunChainSuite(t *testing.T) {
	suite.Run(t, new(RunChainSuite))
}

func (s *RunChainSuite) TestRunChainSuiteEmpty() {}

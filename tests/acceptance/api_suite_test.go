package acceptance

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestApiSuite(t *testing.T) {
	suite.Run(t, new(APISuite))
}

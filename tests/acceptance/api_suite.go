package acceptance

import (
	"github.com/stretchr/testify/suite"
	"os"
)

type APISuite struct {
	suite.Suite
	baseUrl string
}

func (suite *APISuite) SetupSuite() {
	base := os.Getenv("API_HOST")
	if base == "" {
		base = "http://localhost:8080"
	}
	suite.baseUrl = base
}

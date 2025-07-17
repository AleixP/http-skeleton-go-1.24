package acceptance

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
)

func (suite *APISuite) TestListFruitReturnsExpectedListOfFruit() {
	url := suite.baseUrl + "/fruits"
	request, err := http.NewRequest("GET", url, nil)
	assert.NoError(suite.T(), err)

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer token")

	resp, err := http.DefaultClient.Do(request)
	assert.NoError(suite.T(), err)
	defer resp.Body.Close()

	assert.Equal(suite.T(), http.StatusOK, resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	assert.NotEmpty(suite.T(), string(body))
}

func (suite *APISuite) TestListFruitWithoutHeaderReturnsForbidden() {
	url := suite.baseUrl + "/fruits"
	request, err := http.NewRequest("GET", url, nil)
	assert.NoError(suite.T(), err)

	request.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(request)
	assert.NoError(suite.T(), err)
	defer resp.Body.Close()

	assert.Equal(suite.T(), http.StatusForbidden, resp.StatusCode)

}

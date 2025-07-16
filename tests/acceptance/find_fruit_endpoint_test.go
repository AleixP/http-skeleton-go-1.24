package acceptance

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
)

func (suite *APISuite) TestFindFruitReturnExpectedFruit() {
	url := suite.baseUrl + "/fruits/1"
	request, err := http.NewRequest("GET", url, nil)
	assert.NoError(suite.T(), err)

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer token")

	resp, err := http.DefaultClient.Do(request)
	assert.NoError(suite.T(), err)
	defer resp.Body.Close()

	assert.Equal(suite.T(), http.StatusOK, resp.StatusCode)
	body, err := io.ReadAll(resp.Body)

	var jsonData any
	if err := json.Unmarshal(body, &jsonData); err != nil {
		fmt.Printf("failed to unmarshal json: %v\n", err)
		return
	}
	assert.Equal(suite.T(), "apple", jsonData.(map[string]interface{})["name"])

}

func (suite *APISuite) TestFindFruitWithoutAuthorizationHeaderReturnsForbiddenError() {
	url := suite.baseUrl + "/fruits/1"
	request, err := http.NewRequest("GET", url, nil)
	assert.NoError(suite.T(), err)

	request.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(request)
	assert.NoError(suite.T(), err)
	defer resp.Body.Close()

	assert.Equal(suite.T(), http.StatusForbidden, resp.StatusCode)

}
func (suite *APISuite) TestFindFruitWithWrongIdReturnsNotFoundResponse() {
	url := suite.baseUrl + "/fruits/potato"
	request, err := http.NewRequest("GET", url, nil)
	assert.NoError(suite.T(), err)

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer token")

	resp, err := http.DefaultClient.Do(request)
	assert.NoError(suite.T(), err)
	defer resp.Body.Close()

	assert.Equal(suite.T(), http.StatusNotFound, resp.StatusCode)
}

package acceptance

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
)

func (suite *APISuite) TestCreateFruitWithCorrectParameters() {
	url := suite.baseUrl + "/fruits"

	payload := map[string]string{
		"name":  "apple",
		"color": "red",
	}
	body, err := json.Marshal(payload)
	assert.NoError(suite.T(), err)
	resp, err := http.Post(url, "application/json", bytes.NewReader(body))
	assert.NoError(suite.T(), err)
	defer resp.Body.Close()

	assert.Equal(suite.T(), http.StatusCreated, resp.StatusCode, "201 resource created")
}

func (suite *APISuite) TestCreateFruitWithWrongParameters() {
	url := suite.baseUrl + "/fruits"

	payload := map[string]string{"name": "apple"}
	body, err := json.Marshal(payload)
	assert.NoError(suite.T(), err)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	assert.NoError(suite.T(), err)
	defer resp.Body.Close()

	assert.Equal(suite.T(), http.StatusBadRequest, resp.StatusCode, "400 bad request")
}

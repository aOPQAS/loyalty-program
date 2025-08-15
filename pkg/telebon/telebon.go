package telebon

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type TelebonClient struct {
	Token      string
	httpClient *http.Client
}

const baseUrl string = "https://dev.api.telebon.ru"

func New(token string) *TelebonClient {
	return &TelebonClient{
		Token:      token,
		httpClient: http.DefaultClient,
	}
}

func (c *TelebonClient) GetSubproducts() (interface{}, error) {
	req, err := http.NewRequest("GET", baseUrl+"/allsubproducts", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Bearer", c.Token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get result: status_code[%d]", resp.StatusCode)
	}

	var res interface{}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return data, nil
}

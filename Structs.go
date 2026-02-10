package pan123

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"
)

type Config struct {
	//UserFile       string `json:"userFile"`
	Domain         string `json:"domain"`
	AccessTokenAPI string `json:"access_token_api"`
}

type Data struct {
	AccessToken string    `json:"accessToken"`
	ExpiredAt   time.Time `json:"expiredAt"`
}

type AccessTokenResponse struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	Data     Data   `json:"data"`
	XTraceID string `json:"x-traceID"`
}

type APIClient struct {
	ClientID      string       `json:"clientID"`
	ClientSecret  string       `json:"clientSecret"`
	AccessToken   string       `json:"accessToken"`
	ExpiredAt     time.Time    `json:"expiredAt"`
	Authorization string       `json:"Authorization"`
	Platform      string       `json:"Platform"`
	ContentType   string       `json:"Content-Type"`
	HttpClient    *http.Client `json:"-"`
}

// getAccessToken 返回发送请求之后返回的body，包含accessToken和expiredAt
func (c APIClient) getAccessTokenWithConfig(config Config) AccessTokenResponse {
	url := config.Domain + config.AccessTokenAPI
	data := "clientID=" + c.ClientID + "&clientSecret=" + c.ClientSecret
	reqBody := bytes.NewBuffer([]byte(data))
	request, err := http.NewRequest("POST", url, reqBody)
	handleErrWithPrintln("Err during http.NewRequest():", err)

	request.Header.Set("Authorization", c.Authorization)
	request.Header.Set("Content-Type", c.ContentType)
	request.Header.Set("Platform", c.Platform)

	resp, err := c.HttpClient.Do(request)
	handleErrWithPrintln("Err during client.Do():", err)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	content, err := io.ReadAll(resp.Body)
	handleErrWithPrintln("Err during io.ReadAll():", err)

	accessTokenResponse := AccessTokenResponse{}
	err = json.Unmarshal(content, &accessTokenResponse)
	handleErrWithPrintln("Err during json.Unmarshal():", err)

	return accessTokenResponse
}

func (c APIClient) getAccessToken() AccessTokenResponse {
	config := Config{
		Domain:         "https://open-api.123pan.com",
		AccessTokenAPI: "/api/v1/access_token",
	}
	return c.getAccessTokenWithConfig(config)
}

func (c APIClient) checkAndUpdateAccessToken() {
	now := time.Now()
	cUTC := c.ExpiredAt.UTC()
	nowUTC := now.UTC()
	if nowUTC.After(cUTC) {
		respBody := c.getAccessToken()
		c.AccessToken = respBody.Data.AccessToken
		c.ExpiredAt = respBody.Data.ExpiredAt
	}
}

func (c APIClient) saveToFile(filePath string) {
	headerStr, err := json.Marshal(c)
	handleErrWithPrintln("Err during json.Marshal():", err)
	err = os.WriteFile(filePath, headerStr, 0666)
	handleErrWithPrintln("Err during os.WriteFile:", err)
}

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
	CreateFileAPI  string `json:"create_file_api"`
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
func (c APIClient) getAccessTokenWithConfig(config Config) (AccessTokenResponse, error) {
	url := config.Domain + config.AccessTokenAPI
	data := "clientID=" + c.ClientID + "&clientSecret=" + c.ClientSecret
	reqBody := bytes.NewBuffer([]byte(data))
	request, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		return AccessTokenResponse{}, err
	}

	request.Header.Set("Authorization", c.Authorization)
	request.Header.Set("Content-Type", c.ContentType)
	request.Header.Set("Platform", c.Platform)

	resp, err := c.HttpClient.Do(request)
	if err != nil {
		return AccessTokenResponse{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return AccessTokenResponse{}, err
	}

	accessTokenResponse := AccessTokenResponse{}
	err = json.Unmarshal(content, &accessTokenResponse)
	if err != nil {
		return AccessTokenResponse{}, err
	}

	return accessTokenResponse, nil
}

func (c APIClient) getAccessToken() (AccessTokenResponse, error) {
	config := getDefaultConfig()
	return c.getAccessTokenWithConfig(config)
}

func (c APIClient) checkAndUpdateAccessToken() error {
	now := time.Now()
	cUTC := c.ExpiredAt.UTC()
	nowUTC := now.UTC()
	if nowUTC.After(cUTC) {
		respBody, err := c.getAccessToken()
		if err != nil {
			return err
		}
		c.AccessToken = respBody.Data.AccessToken
		c.ExpiredAt = respBody.Data.ExpiredAt
	}
	return nil
}

func (c APIClient) checkAndUpdateAccessTokenAndSave(filePath string) error {
	err := c.checkAndUpdateAccessToken()
	if err != nil {
		return err
	}
	c.saveToFile(filePath)
	return nil
}

func (c APIClient) saveToFile(filePath string) error {
	headerStr, err := json.Marshal(c)
	if err != nil {
		return err
	}
	err = os.WriteFile(filePath, headerStr, 0666)
	if err != nil {
		return err
	}
	return nil
}

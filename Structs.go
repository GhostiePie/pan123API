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

type User struct {
	UserWithoutAccAndExp
	AccessToken string    `json:"accessToken"`
	ExpiredAt   time.Time `json:"expiredAt"`
}

type UserWithoutAccAndExp struct {
	ClientID     string `json:"clientID"`
	ClientSecret string `json:"clientSecret"`
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

type Header struct {
	Authorization string `json:"Authorization"`
	Platform      string `json:"Platform"`
	ContentType   string `json:"Content-Type"`
}

// getAccessToken 返回发送请求之后返回的body，包含accessToken和expiredAt
func (u User) getAccessTokenWithConfig(config Config) AccessTokenResponse {
	client := &http.Client{}
	url := config.Domain + config.AccessTokenAPI
	data := "clientID=" + u.ClientID + "&clientSecret=" + u.ClientSecret
	reqBody := bytes.NewBuffer([]byte(data))
	request, err := http.NewRequest("POST", url, reqBody)
	handleErrWithPrintln("Err during http.NewRequest():", err)
	header := getPublicHeader()

	request.Header.Set("Authorization", header.Authorization)
	request.Header.Set("Content-Type", header.ContentType)
	request.Header.Set("Platform", header.Platform)

	resp, err := client.Do(request)
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

func (u User) getAccessToken() AccessTokenResponse {
	config := Config{
		Domain:         "https://open-api.123pan.com",
		AccessTokenAPI: "/api/v1/access_token",
	}
	return u.getAccessTokenWithConfig(config)
}

func (u User) checkAndUpdateAccessToken() {
	now := time.Now()
	cUTC := u.ExpiredAt.UTC()
	nowUTC := now.UTC()
	if nowUTC.After(cUTC) {
		respBody := u.getAccessToken()
		u.AccessToken = respBody.Data.AccessToken
		u.ExpiredAt = respBody.Data.ExpiredAt
	}
}

func (h Header) saveToFile(filePath string) {
	headerStr, err := json.Marshal(h)
	handleErrWithPrintln("Err during json.Marshal():", err)
	err = os.WriteFile(filePath, headerStr, 0666)
	handleErrWithPrintln("Err during os.WriteFile:", err)
}

func (u User) saveToFile(filePath string) {
	userStr, err := json.Marshal(u)
	handleErrWithPrintln("Err during json.Marshal():", err)
	err = os.WriteFile(filePath, userStr, 0666)
	handleErrWithPrintln("Err during os.WriteFile:", err)
}

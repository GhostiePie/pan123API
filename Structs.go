package pan123

import (
	"encoding/json"
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

type Response struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
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

// 将APIClient以json格式存储至文件
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

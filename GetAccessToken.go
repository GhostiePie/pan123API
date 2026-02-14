package pan123

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

type AccessTokenData struct {
	AccessToken string    `json:"accessToken"`
	ExpiredAt   time.Time `json:"expiredAt"`
}

type AccessTokenResponse struct {
	Response
	Data AccessTokenData `json:"data"`
}

// GetAccessTokenWithConfig 直接将AccessToken和ExpiredAt赋值给调用对象c, 并直接将c.Authorization设置为AccessToken, 但考虑到可能的拓展仍然保留了c.AccessToken
func (c *APIClient) GetAccessTokenWithConfig(config Config) error {
	url := config.Domain + config.AccessTokenAPI
	data := "clientID=" + c.ClientID + "&clientSecret=" + c.ClientSecret
	reqBody := bytes.NewBuffer([]byte(data))
	request, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		return err
	}

	request.Header.Set("Authorization", c.Authorization)
	request.Header.Set("Content-Type", c.ContentType)
	request.Header.Set("Platform", c.Platform)

	resp, err := c.HttpClient.Do(request)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	accessTokenResponse := AccessTokenResponse{}
	err = json.Unmarshal(content, &accessTokenResponse)
	if err != nil {
		return err
	}

	if accessTokenResponse.Code != 0 {
		return errors.New(accessTokenResponse.Message)
	}

	c.AccessToken = accessTokenResponse.Data.AccessToken
	c.Authorization = accessTokenResponse.Data.AccessToken
	c.ExpiredAt = accessTokenResponse.Data.ExpiredAt

	return nil
}

// GetAccessToken 只是对getAccessTokenWithConfig进行了封装，并将默认config填入
func (c *APIClient) GetAccessToken() error {
	config := GetDefaultConfig()
	return c.GetAccessTokenWithConfig(config)
}

// CheckAndUpdateAccessToken 自动检查当前时间和ExpiredAt时间，如果过期则自动更改，未过期则什么也不干
func (c *APIClient) CheckAndUpdateAccessToken() (bool, error) {
	now := time.Now()
	cUTC := c.ExpiredAt.UTC()
	nowUTC := now.UTC()
	if nowUTC.After(cUTC) {
		err := c.GetAccessToken()
		if err != nil {
			return false, err
		}
		return true, nil
	} else {
		return false, nil
	}
}

// CheckAndUpdateAccessTokenAndSave 封装了 checkAndUpdateAccessToken , 但是会在修改后将新数据写入文件
func (c *APIClient) CheckAndUpdateAccessTokenAndSave(filePath string) error {
	changed, err := c.CheckAndUpdateAccessToken()
	if err != nil {
		return err
	}
	if changed {
		err = c.SaveToFile(filePath)
		if err != nil {
			return err
		}
	}
	return nil
}

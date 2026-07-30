package Client

import (
	"io"
	"net/http"
	"strings"
)

// Get 发送GET请求，自动设置Authorization和Platform请求头
func (c *APIClient) Get(url string, contentType string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest("GET", url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Authorization", c.Authorization)
	req.Header.Set("Platform", c.Config.Platform)

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	responseBody, err := io.ReadAll(resp.Body)
	return responseBody, err
}

// GetData 发送带JSON数据的GET请求
func (c *APIClient) GetData(url string, data string) ([]byte, error) {
	return c.Get(url, "application/json", strings.NewReader(data))
}

// GetQuery 发送不带请求体的GET请求
func (c *APIClient) GetQuery(url string) ([]byte, error) {
	return c.Get(url, "application/json", nil)
}

// Post 发送POST请求，自动设置Authorization和Platform请求头
func (c *APIClient) Post(url string, contentType string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return []byte{}, err
	}
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Authorization", c.Authorization)
	req.Header.Set("Platform", c.Config.Platform)

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	return respBody, err
}

// PostData 发送带JSON数据的POST请求
func (c *APIClient) PostData(url string, data io.Reader) ([]byte, error) {
	return c.Post(url, "application/json", data)
}

// PostQuery 发送不带请求体的POST请求
func (c *APIClient) PostQuery(url string) ([]byte, error) {
	return c.Post(url, "application/json", nil)
}

// Put 发送PUT请求，自动设置Authorization和Platform请求头
func (c *APIClient) Put(url string, contentType string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest("PUT", url, body)
	if err != nil {
		return []byte{}, err
	}
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Authorization", c.Authorization)
	req.Header.Set("Platform", c.Config.Platform)

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	return respBody, err
}

// PutData 发送带JSON数据的PUT请求
func (c *APIClient) PutData(url string, data io.Reader) ([]byte, error) {
	return c.Put(url, "application/json", data)
}

// PutQuery 发送不带请求体的PUT请求
func (c *APIClient) PutQuery(url string) ([]byte, error) {
	return c.Put(url, "application/json", nil)
}

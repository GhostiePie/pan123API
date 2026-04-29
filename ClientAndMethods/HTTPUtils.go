package ClientAndMethods

import (
	"io"
	"net/http"
	"strings"
)

func (c *APIClient) Get(url string, contentType string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest("GET", url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Authorization", c.Authorization)
	req.Header.Set("Platform", c.Platform)

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	responseBody, err := io.ReadAll(resp.Body)
	return responseBody, err
}

func (c *APIClient) GetData(url string, data string) ([]byte, error) {
	return c.Get(url, "application/json", strings.NewReader(data))
}

func (c *APIClient) GetQuery(url string) ([]byte, error) {
	return c.Get(url, "application/json", nil)
}

func (c *APIClient) Post(url string, contentType string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return []byte{}, err
	}
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Authorization", c.Authorization)
	req.Header.Set("Platform", c.Platform)

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	return respBody, err
}

func (c *APIClient) PostData(url string, data io.Reader) ([]byte, error) {
	return c.Post(url, "application/json", data)
}

func (c *APIClient) PostQuery(url string) ([]byte, error) {
	return c.Post(url, "application/json", nil)
}

func (c *APIClient) Put(url string, contentType string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest("PUT", url, body)
	if err != nil {
		return []byte{}, err
	}
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Authorization", c.Authorization)
	req.Header.Set("Platform", c.Platform)

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	return respBody, err
}

func (c *APIClient) PutData(url string, data io.Reader) ([]byte, error) {
	return c.Put(url, "application/json", data)
}

func (c *APIClient) PutQuery(url string) ([]byte, error) {
	return c.Put(url, "application/json", nil)
}

package pan123

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

var errAPIClient = APIClient{
	"err",
	"err",
	"err",
	time.Time{},
	"err",
	"err",
	"err",
	&http.Client{},
}

func getDefaultConfig() Config {
	return Config{
		Domain:         "https://open-api.123pan.com",
		AccessTokenAPI: "/api/v1/access_token",
		CreateFileAPI:  "/upload/v2/file/create",
	}
}

func readAPIClientFromJson(jsonStr string) (APIClient, error) {
	var err error
	apiClient := APIClient{}
	err = json.Unmarshal([]byte(jsonStr), &apiClient)
	apiClient.HttpClient = &http.Client{}
	return apiClient, err
}

func readAPIClientFromFile(filePath string) (APIClient, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return APIClient{}, err
	}
	var apiClient APIClient
	err = json.Unmarshal(data, &apiClient)
	if err != nil {
		return apiClient, err
	}
	return apiClient, nil
}

func (c APIClient) Post(url string, contentType string, body io.Reader) ([]byte, error) {
	//reqBody := bytes.NewBufferString(data)
	//reqBody := strings.NewReader(data)
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
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

func (c APIClient) PostData(url string, data string) ([]byte, error) {
	return c.Post(url, "application/json", strings.NewReader(data))
}

func splitFile(fileName string, chunkIndex int, chunkSize int) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	offset := int64((chunkIndex - 1) * chunkSize)
	buffer := make([]byte, chunkSize)
	n, err := file.ReadAt(buffer, offset)

	if err != nil {
		return nil, err
	}

	if n < int(chunkSize) {
		return buffer[:n], nil
	}

	return buffer, nil

}

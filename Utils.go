package pan123

import (
	"encoding/json"
	"net/http"
	"os"
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

package pan123

import (
	"encoding/json"
	"fmt"
	"log"
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

func handleErrWithPrintln(msg string, err error) bool {
	if err != nil {
		fmt.Println(msg, err)
		return true
	} else {
		return false
	}
}

func handleErrWithFatalln(msg string, err error) bool {
	if err != nil {
		log.Fatalln(msg, err)
		return true
	} else {
		return false
	}
}

func getDefaultConfig() Config {
	return Config{
		Domain:         "https://open-api.123pan.com",
		AccessTokenAPI: "/api/v1/access_token",
	}
}

func readAPIClientFromJson(jsonStr string) APIClient {
	apiClient := APIClient{}
	err := json.Unmarshal([]byte(jsonStr), &apiClient)
	handleErrWithPrintln("Err during json.Unmarshal():", err)
	apiClient.HttpClient = &http.Client{}
	return apiClient
}

func readAPIClientFromFile(filePath string) APIClient {
	data, err := os.ReadFile(filePath)
	if !handleErrWithPrintln("Err during os.ReadFile():", err) {
		var apiClient APIClient
		err := json.Unmarshal(data, &apiClient)
		if handleErrWithPrintln("Err during json.Unmarshal():", err) {
			return errAPIClient
		} else {
			apiClient.HttpClient = &http.Client{}
			return apiClient
		}
	} else {
		return errAPIClient
	}

}

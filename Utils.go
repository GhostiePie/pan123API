package pan123

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

var errUser = User{
	UserWithoutAccAndExp{
		ClientID:     "err",
		ClientSecret: "err",
	},
	"err",
	time.Time{},
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

func getPublicHeader() Header {
	return Header{
		Authorization: "Bearer access_token",
		Platform:      "open_platform",
		ContentType:   "application/json",
	}
}

func getDefaultConfig() Config {
	return Config{
		Domain:         "https://open-api.123pan.com",
		AccessTokenAPI: "/api/v1/access_token",
	}
}

func readUserFromJson(jsonStr string) User {
	user := User{}
	err := json.Unmarshal([]byte(jsonStr), &user)
	handleErrWithPrintln("Err during json.Unmarshal():", err)
	return user
}

func readUserFromFile(filePath string) User {
	data, err := os.ReadFile(filePath)
	if !handleErrWithPrintln("Err during os.ReadFile():", err) {
		var user User
		err := json.Unmarshal(data, &user)
		if handleErrWithPrintln("Err during json.Unmarshal():", err) {
			return errUser
		} else {
			return user
		}
	} else {
		return errUser
	}

}

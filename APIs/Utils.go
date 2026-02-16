package APIs

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"strings"
)

var (
	ErrInsufficientDownloadTraffic = errors.New("insufficient download traffic")
	ErrFileNotExists               = errors.New("file not exists")
	ErrGeneric                     = errors.New("api error")
)

func GetDefaultConfig() APIConfig {
	return DefaultConfig
}

// StructToQueryString 将结构体转换为查询字符串
func StructToQueryString(obj interface{}) (string, error) {
	values := url.Values{}

	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		// 获取 json 标签，如果没有则使用字段名
		tag := field.Tag.Get("json")

		if tag == "" {
			tag = field.Name
		}

		// 处理 omitempty 之类的选项
		if commaIdx := strings.Index(tag, ","); commaIdx != -1 {
			tag = tag[:commaIdx]
		}

		// 跳过 "-"
		if tag == "-" {
			continue
		}

		// 处理不同类型的值
		switch value.Kind() {
		case reflect.String:
			if str := value.String(); str != "" {
				values.Set(tag, str)
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if val := value.Int(); val != 0 {
				values.Set(tag, strconv.FormatInt(val, 10))
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if val := value.Uint(); val != 0 {
				values.Set(tag, strconv.FormatUint(val, 10))
			}
		case reflect.Float32, reflect.Float64:
			if val := value.Float(); val != 0 {
				values.Set(tag, strconv.FormatFloat(val, 'f', -1, 64))
			}
		case reflect.Bool:
			values.Set(tag, strconv.FormatBool(value.Bool()))
		case reflect.Slice:
			// 处理切片，比如 []string
			if value.Type().Elem().Kind() == reflect.String {
				for j := 0; j < value.Len(); j++ {
					values.Add(tag, value.Index(j).String())
				}
			}
		}
	}

	return values.Encode(), nil
}

func ReadAPIClientFromJson(jsonStr string) (APIClient, error) {
	var err error
	apiClient := APIClient{}
	err = json.Unmarshal([]byte(jsonStr), &apiClient)
	apiClient.HttpClient = &http.Client{}
	return apiClient, err
}

func ReadAPIClientFromFile(filePath string) (APIClient, error) {
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

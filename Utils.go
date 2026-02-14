package pan123

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strconv"
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

var (
	errInsufficientDownloadTraffic = errors.New("insufficient download traffic")
	errFileNotExists               = errors.New("file not exists")
	errGeneric                     = errors.New("api error")
)

func GetDefaultConfig() Config {
	return Config{
		Domain:                    "https://open-api.123pan.com",
		AccessTokenAPI:            "/api/v1/access_token",
		CreateFileAPI:             "/upload/v2/file/create",
		UploadSliceAPI:            "/upload/v2/file/slice",
		UploadCompleteAPI:         "/upload/v2/file/upload_complete",
		GetUploadDomainAPI:        "/upload/v2/file/domain",
		OneStepUploadAPI:          "/upload/v2/file/single/create",
		UploadSHA1API:             "/upload/v2/file/sha1",
		OneFileRenameAPI:          "/api/v1/file/name",
		BatchFilesRenameAPI:       "/api/v1/file/rename",
		DeleteFileToTrashAPI:      "/api/v1/file/trash",
		CopyOneFileAPI:            "/api/v1/file/copy",
		CopyBatchFilesAPI:         "/api/v1/file/async/copy",
		CopyBatchFilesProgressAPI: "/api/v1/file/async/copy/process",
		RecoverFileFromTrashAPI:   "/api/v1/file/recover",
		RecoverFileByPathAPI:      "/api/v1/file/recover/by_path",
		GetOneFileDetailAPI:       "/api/v1/file/detail",
		GetMultipleFilesDetailAPI: "/api/v1/file/infos",
		GetFileListAPI:            "/api/v2/file/list",
		MoveFilesAPI:              "/api/v1/file/move",
		DownloadFileAPI:           "/api/v1/file/download_info",
	}
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

func (c *APIClient) Post(url string, contentType string, body io.Reader) ([]byte, error) {
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

func (c *APIClient) PostData(url string, data string) ([]byte, error) {
	return c.Post(url, "application/json", strings.NewReader(data))
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

func (c *APIClient) PutData(url string, data string) ([]byte, error) {
	return c.Put(url, "application/json", strings.NewReader(data))
}

// SplitFile 用于从某个特定文件中一次性读出多个分片，会将文件一次性全部加载到内存，慎用。
func SplitFile(fileName string, chunkSize int) ([][]byte, error) {
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		return nil, err
	}

	// 由于后期会将可接受文件体积限制在2GB以下，因此可以放心转为int以简化后期计算，不用担心fileSize超出int范围
	fileSize := int(fileInfo.Size())

	if fileSize > int(math.Pow(2, 30)) { // 如果文件大于2GB
		//return nil, errors.New("file \"" + fileName + "\" is bigger than 2GB")
		return nil, fmt.Errorf("file %q is bigger than 2GB", fileName)
	}

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// 计算分片数量
	chunkCount := (fileSize + chunkSize - 1) / chunkSize
	result := make([][]byte, 0, chunkCount)

	for i := 0; i < chunkCount; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > fileSize {
			end = fileSize
		}

		chunk, err := ReadChunk(file, start, end)
		if err != nil {
			return result, err
		}
		result = append(result, chunk)
	}

	return result, nil
}

// ReadChunk 用于从打开的文件中读取特定分片，需要传入已打开的文件，由使用者自主管理文件的打开与关闭。
func ReadChunk(file *os.File, start int, end int) ([]byte, error) {
	// 如果起始位置超出文件大小
	stat, _ := file.Stat()
	if int64(start) >= stat.Size() {
		return []byte{}, io.EOF
	}

	// 确保 end 不超出文件大小
	if int64(end) > stat.Size() {
		end = int(stat.Size())
	}

	buffer := make([]byte, end-start)
	n, err := file.ReadAt(buffer, int64(start))
	if err != nil {
		return nil, err
	}

	// 即使有 EOF，也返回已读取的数据
	return buffer[:n], nil
}

// GetFileChunk 用于快捷读取文件的某个分片，包含文件打开与关闭，不建议频繁调用。
// 如需从某个文件频繁读出多个分片，请考虑 SplitFile 或 ReadChunk
func GetFileChunk(fileName string, start int, end int) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return ReadChunk(file, start, end)
}

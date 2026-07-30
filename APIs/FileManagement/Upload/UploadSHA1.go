package Upload

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/GhostiePie/pan123API/Client"
)

// UploadSHA1Body SHA1秒传请求体
type UploadSHA1Body struct {
	ParentFileID int    `json:"parentFileID"` // 父目录ID
	FileName     string `json:"fileName"`     // 文件名
	SHA1         string `json:"sha1"`         // 文件SHA1值
	Size         int    `json:"size"`         // 文件大小(bytes)
	Duplicate    int    `json:"duplicate"`    // 重名策略：1重命名 2覆盖 3跳过
}

// UploadSHA1Data SHA1秒传返回的数据
type UploadSHA1Data struct {
	FileID int  `json:"fileID"` // 文件ID
	Reuse  bool `json:"reuse"`  // 是否秒传命中
}

// UploadSHA1Response SHA1秒传响应
type UploadSHA1Response struct {
	Client.Response
	Data UploadSHA1Data `json:"data"`
}

// UploadSHA1 通过SHA1值进行秒传，若匹配则跳过上传
func UploadSHA1(c *Client.APIClient, uploadSHA1Body UploadSHA1Body) (UploadSHA1Response, error) {
	url := c.Config.Domain + c.Config.UploadSHA1API
	data := "parentFileID=" + strconv.Itoa(uploadSHA1Body.ParentFileID) +
		"&filename=" + uploadSHA1Body.FileName +
		"&sha1=" + uploadSHA1Body.SHA1 +
		"&size=" + strconv.Itoa(uploadSHA1Body.Size)
	if uploadSHA1Body.Duplicate != 0 {
		data += "&duplicate=" + strconv.Itoa(uploadSHA1Body.Duplicate)
	}

	body, err := c.PostData(url, strings.NewReader(data))
	if err != nil {
		return UploadSHA1Response{}, err
	}

	uploadSHA1Response := UploadSHA1Response{}
	err = json.Unmarshal(body, &uploadSHA1Response)
	return uploadSHA1Response, err
}

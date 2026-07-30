package Upload

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"mime/multipart"
	"strconv"

	"github.com/GhostiePie/pan123API/Client"
)

// OneStepUploadBody 单步上传（小文件）请求体
type OneStepUploadBody struct {
	ParentFileID int    `json:"parentFileID"` // 父目录ID
	FileName     string `json:"fileName"`     // 文件名
	Etag         string `json:"etag"`         // 文件Etag值
	Size         int    `json:"size"`         // 文件大小(bytes)
	File         []byte `json:"file"`         // 文件二进制内容
	Duplicate    int    `json:"duplicate"`    // 重名策略：1重命名 2覆盖 3跳过
	ContainDir   bool   `json:"containDir"`   // 是否包含目录
}

// OneStepUploadData 单步上传返回的数据
type OneStepUploadData struct {
	FileID    int  `json:"fileID"`    // 文件ID
	Completed bool `json:"completed"` // 是否上传完成
}

// OneStepUploadResponse 单步上传响应
type OneStepUploadResponse struct {
	Client.Response
	Data OneStepUploadData `json:"data"`
}

// OneStepUpload 单步上传小文件（适用于小文件，无需分片）
func OneStepUpload(c *Client.APIClient, oneStepUploadBody OneStepUploadBody) (OneStepUploadResponse, error) {
	url := c.Config.Domain + c.Config.OneStepUploadAPI
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	writer.WriteField("parentFileID", strconv.Itoa(oneStepUploadBody.ParentFileID))
	writer.WriteField("fileName", oneStepUploadBody.FileName)
	writer.WriteField("etag", oneStepUploadBody.Etag)
	writer.WriteField("size", strconv.Itoa(oneStepUploadBody.Size))
	if oneStepUploadBody.Duplicate != 0 {
		writer.WriteField("duplicate", strconv.Itoa(oneStepUploadBody.Duplicate))
	}
	if oneStepUploadBody.ContainDir {
		writer.WriteField("containDir", "true")
	}

	part, err := writer.CreateFormFile("file", oneStepUploadBody.FileName)
	if err != nil {
		return OneStepUploadResponse{}, errors.New("err During CreateFormFile: " + err.Error())
	}
	part.Write(oneStepUploadBody.File)
	writer.Close()

	resp, err := c.Post(url, writer.FormDataContentType(), body)
	if err != nil {
		return OneStepUploadResponse{}, errors.New("err During c.Post: " + err.Error())
	}
	fmt.Println(string(resp))

	oneStepUploadResponse := OneStepUploadResponse{}
	err = json.Unmarshal(resp, &oneStepUploadResponse)
	if err != nil {
		return OneStepUploadResponse{}, errors.New("err During Unmarshal: " + err.Error())
	}

	return oneStepUploadResponse, nil
}

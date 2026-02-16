package APIs

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"mime/multipart"
	"strconv"
)

type UploadSliceBody struct {
	PreUploadID string   `json:"preuploadID"` // 必填	预上传ID
	SliceNo     int      `json:"sliceNo"`     // 必填	分片序号，从1开始自增
	SliceMD5    string   `json:"sliceMD5"`    // 必填	当前分片md5
	Slice       []byte   `json:"slice"`       // 必填	分片二进制流
	Servers     []string `json:"servers"`     // 注：非Body参数！该API需要CreateFile()的返回值中的servers值作为域名
}

type UploadSliceResponse struct {
	Response
}

func (c *APIClient) UploadSlice(uploadSliceBody UploadSliceBody) (UploadSliceResponse, error) {

	url := uploadSliceBody.Servers[0] + c.Config.UploadSliceAPI
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	writer.WriteField("preuploadID", uploadSliceBody.PreUploadID)
	writer.WriteField("sliceNo", strconv.Itoa(uploadSliceBody.SliceNo))
	writer.WriteField("sliceMD5", uploadSliceBody.SliceMD5)

	part, err := writer.CreateFormFile("slice", "filename"+strconv.Itoa(uploadSliceBody.SliceNo))
	if err != nil {
		return UploadSliceResponse{}, errors.New("err During CreateFromFile: " + err.Error())
	}
	part.Write(uploadSliceBody.Slice)
	writer.Close()
	resp, err := c.Post(url, writer.FormDataContentType(), body)
	if err != nil {
		return UploadSliceResponse{}, errors.New("err During c.Post: " + err.Error())
	}
	fmt.Println(string(resp))
	uploadSliceResponse := UploadSliceResponse{}
	err = json.Unmarshal(resp, &uploadSliceResponse)
	if err != nil {
		return UploadSliceResponse{}, errors.New("err During Unmarshal: " + err.Error())
	}
	return uploadSliceResponse, err
}

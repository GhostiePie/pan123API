package pan123

type UploadSliceBody struct {
	PreUploadID string `json:"preuploadID"` // 必填	预上传ID
	SliceNo     int    `json:"sliceNo"`     // 必填	分片序号，从1开始自增
	SliceMD5    string `json:"sliceMD5"`    // 必填	当前分片md5
	Slice       []byte `json:"slice"`       // 必填	分片二进制流
}

type UploadSliceResponse struct {
	Response
}

//func (c APIClient) uploadSliceWithConfig(uploadSliceBody UploadSliceBody, config Config) (UploadSliceResponse, error) {
//
//	url := config.Domain + config.CreateFileAPI
//	body := &bytes.Buffer{}
//	writer := multipart.NewWriter(body)
//
//	writer.WriteField("preuploadID", uploadSliceBody.PreUploadID)
//	writer.WriteField("sliceNo", strconv.Itoa(uploadSliceBody.SliceNo))
//	writer.WriteField("sliceMD5", uploadSliceBody.SliceMD5)
//
//	data := "parentFileID=" + strconv.Itoa(createFileBody.ParentFileID) + "&filename=" + createFileBody.FileName + "&etag=" + createFileBody.Etag + "&size=" + strconv.Itoa(createFileBody.Size)
//	if createFileBody.Duplicate != 0 {
//		data += "&duplicate=" + strconv.Itoa(createFileBody.Duplicate)
//	}
//	if createFileBody.ContainDir {
//		data += "&containDir=true"
//	}
//	body, _ := c.PostData(url, data)
//	createFileResponse := CreateFileResponse{}
//	err := json.Unmarshal(body, &createFileResponse)
//
//	return createFileResponse, err
//}

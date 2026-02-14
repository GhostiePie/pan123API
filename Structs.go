package pan123

import (
	"encoding/json"
	"net/http"
	"os"
	"time"
)

type Config struct {
	//UserFile       string `json:"userFile"`
	Domain                    string `json:"domain"`
	AccessTokenAPI            string `json:"access_token_api"`
	CreateFileAPI             string `json:"create_file_api"`
	UploadSliceAPI            string `json:"upload_slice_api"`
	UploadCompleteAPI         string `json:"upload_complete_api"`
	GetUploadDomainAPI        string `json:"get_upload_domain_api"`
	OneStepUploadAPI          string `json:"one_step_upload_api"`
	UploadSHA1API             string `json:"upload_sha1_api"`
	OneFileRenameAPI          string `json:"one_file_rename_api"`
	BatchFilesRenameAPI       string `json:"batch_files_rename_api"`
	DeleteFileToTrashAPI      string `json:"delete_file_to_trash_api"`
	CopyOneFileAPI            string `json:"copy_one_file_api"`
	CopyBatchFilesAPI         string `json:"copy_batch_files_api"`
	CopyBatchFilesProgressAPI string `json:"copy_batch_files_progress_api"`
	RecoverFileFromTrashAPI   string `json:"recover_file_from_trash_api"`
	RecoverFileByPathAPI      string `json:"recover_file_by_path_api"`
	GetOneFileDetailAPI       string `json:"get_one_file_detail_api"`
	GetMultipleFilesDetailAPI string `json:"get_multiple_files_detail_api"`
	GetFileListAPI            string `json:"get_file_list_api"`
	MoveFilesAPI              string `json:"move_files_api"`
	DownloadFileAPI           string `json:"download_file_api"`
}

type Response struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	XTraceID string `json:"x-traceID"`
}

type APIClient struct {
	ClientID      string       `json:"clientID"`
	ClientSecret  string       `json:"clientSecret"`
	AccessToken   string       `json:"accessToken"`
	ExpiredAt     time.Time    `json:"expiredAt"`
	Authorization string       `json:"Authorization"`
	Platform      string       `json:"Platform"`
	ContentType   string       `json:"Content-Type"`
	HttpClient    *http.Client `json:"-"`
}

// SaveToFile 将APIClient以json格式存储至文件
func (c *APIClient) SaveToFile(filePath string) error {
	headerStr, err := json.Marshal(c)
	if err != nil {
		return err
	}
	err = os.WriteFile(filePath, headerStr, 0666)
	if err != nil {
		return err
	}
	return nil
}

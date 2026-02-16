package APIs

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type APIConfig struct {
	//UserFile       string `json:"userFile"`
	Domain                          string `json:"domain"`
	AccessTokenAPI                  string `json:"access_token_api"`
	CreateDirectoryAPI              string `json:"create_directory_api"`
	CreateFileAPI                   string `json:"create_file_api"`
	UploadSliceAPI                  string `json:"upload_slice_api"`
	UploadCompleteAPI               string `json:"upload_complete_api"`
	GetUploadDomainAPI              string `json:"get_upload_domain_api"`
	OneStepUploadAPI                string `json:"one_step_upload_api"`
	UploadSHA1API                   string `json:"upload_sha1_api"`
	OneFileRenameAPI                string `json:"one_file_rename_api"`
	BatchFilesRenameAPI             string `json:"batch_files_rename_api"`
	DeleteFileToTrashAPI            string `json:"delete_file_to_trash_api"`
	CopyOneFileAPI                  string `json:"copy_one_file_api"`
	CopyBatchFilesAPI               string `json:"copy_batch_files_api"`
	CopyBatchFilesProgressAPI       string `json:"copy_batch_files_progress_api"`
	RecoverFileFromTrashAPI         string `json:"recover_file_from_trash_api"`
	RecoverFileByPathAPI            string `json:"recover_file_by_path_api"`
	GetOneFileDetailAPI             string `json:"get_one_file_detail_api"`
	GetMultipleFilesDetailAPI       string `json:"get_multiple_files_detail_api"`
	GetFileListAPI                  string `json:"get_file_list_api"`
	MoveFilesAPI                    string `json:"move_files_api"`
	DownloadFileAPI                 string `json:"download_file_api"`
	CreateShareLinkAPI              string `json:"create_share_link_api"`
	GetShareLinkListAPI             string `json:"get_share_link_list_api"`
	ModifyShareLinkAPI              string `json:"modify_share_link_api"`
	CreatePaidShareLinkAPI          string `json:"create_paid_share_link_api"`
	GetPaidShareLinkListAPI         string `json:"get_paid_share_link_list_api"`
	ModifyPaidShareLinkAPI          string `json:"modify_paid_share_link_api"`
	CreateOfflineDownloadMissionAPI string `json:"create_offline_download_mission_api"`
	GetOfflineDownloadProgressAPI   string `json:"get_offline_download_progress_api"`
	GetUserInfoAPI                  string `json:"get_user_info_api"`
	SwitchIPBlacklistListAPI        string `json:"switch_ip_blacklist_list_api"`
	UpdateIPBlacklistListAPI        string `json:"update_ip_blacklist_list_api"`
	GetIPBlacklistListAPI           string `json:"get_ip_blacklist_list_api"`
	GetDirectLinkOfflineLogsAPI     string `json:"get_direct_link_offline_logs_api"`
	GetDirectLinkTrafficLogsAPI     string `json:"get_direct_link_traffic_logs_api"`
	EnableDirectLinkAPI             string `json:"enable_direct_link_api"`
	GetDirectLinkURLAPI             string `json:"get_direct_link_url_api"`
	DisableDirectLinkAPI            string `json:"disable_direct_link_api"`
	DirectLinkCacheRefreshAPI       string `json:"direct_link_cache_refresh_api"`
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
	Config        *APIConfig   `json:"-"`
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

func (c *APIClient) GetQuery(url string) ([]byte, error) {
	return c.Get(url, "application/json", nil)
}

func (c *APIClient) Post(url string, contentType string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest("POST", url, body)
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

func (c *APIClient) PostData(url string, data string) ([]byte, error) {
	return c.Post(url, "application/json", strings.NewReader(data))
}

func (c *APIClient) PostQuery(url string) ([]byte, error) {
	return c.Post(url, "application/json", nil)
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

func (c *APIClient) PutQuery(url string) ([]byte, error) {
	return c.Put(url, "application/json", nil)
}

var DefaultConfig = APIConfig{
	Domain:                          "https://open-api.123pan.com",
	AccessTokenAPI:                  "/api/v1/access_token",
	CreateFileAPI:                   "/upload/v2/file/create",
	CreateDirectoryAPI:              "/upload/v1/file/mkdir",
	UploadSliceAPI:                  "/upload/v2/file/slice",
	UploadCompleteAPI:               "/upload/v2/file/upload_complete",
	GetUploadDomainAPI:              "/upload/v2/file/domain",
	OneStepUploadAPI:                "/upload/v2/file/single/create",
	UploadSHA1API:                   "/upload/v2/file/sha1",
	OneFileRenameAPI:                "/api/v1/file/name",
	BatchFilesRenameAPI:             "/api/v1/file/rename",
	DeleteFileToTrashAPI:            "/api/v1/file/trash",
	CopyOneFileAPI:                  "/api/v1/file/copy",
	CopyBatchFilesAPI:               "/api/v1/file/async/copy",
	CopyBatchFilesProgressAPI:       "/api/v1/file/async/copy/process",
	RecoverFileFromTrashAPI:         "/api/v1/file/recover",
	RecoverFileByPathAPI:            "/api/v1/file/recover/by_path",
	GetOneFileDetailAPI:             "/api/v1/file/detail",
	GetMultipleFilesDetailAPI:       "/api/v1/file/infos",
	GetFileListAPI:                  "/api/v2/file/list",
	MoveFilesAPI:                    "/api/v1/file/move",
	DownloadFileAPI:                 "/api/v1/file/download_info",
	CreateShareLinkAPI:              "/api/v1/share/create",
	GetShareLinkListAPI:             "/api/v1/share/list",
	ModifyShareLinkAPI:              "/api/v1/share/list/info",
	CreatePaidShareLinkAPI:          "/api/v1/share/content-payment/create",
	GetPaidShareLinkListAPI:         "/api/v1/share/payment/list",
	ModifyPaidShareLinkAPI:          "/api/v1/share/list/payment/info",
	CreateOfflineDownloadMissionAPI: "/api/v1/offline/download",
	GetOfflineDownloadProgressAPI:   "/api/v1/offline/download/process",
	GetUserInfoAPI:                  "/api/v1/user/info",
	SwitchIPBlacklistListAPI:        "/api/v1/developer/config/forbide-ip/switch",
	UpdateIPBlacklistListAPI:        "/api/v1/developer/config/forbide-ip/update",
	GetIPBlacklistListAPI:           "/api/v1/developer/config/forbide-ip/list",
	GetDirectLinkOfflineLogsAPI:     "/api/v1/direct-link/offline/logs",
	GetDirectLinkTrafficLogsAPI:     "/api/v1/direct-link/log",
	EnableDirectLinkAPI:             "/api/v1/direct-link/enable",
	GetDirectLinkURLAPI:             "/api/v1/direct-link/url",
	DisableDirectLinkAPI:            "/api/v1/direct-link/disable",
	DirectLinkCacheRefreshAPI:       "/api/v1/direct-link/cache/refresh",
}

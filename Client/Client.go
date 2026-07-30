package Client

import (
	"log"
	"net/http"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

// APIConfig API配置，包含所有API的路径端点
type APIConfig struct {
	//UserFile       string `json:"userFile"`
	Domain                                    string `json:"domain"`
	Platform                                  string `json:"platform"`
	AccessTokenAPI                            string `json:"access_token_api"`
	CreateDirectoryAPI                        string `json:"create_directory_api"`
	CreateFileAPI                             string `json:"create_file_api"`
	UploadSliceAPI                            string `json:"upload_slice_api"`
	UploadCompleteAPI                         string `json:"upload_complete_api"`
	GetUploadDomainAPI                        string `json:"get_upload_domain_api"`
	OneStepUploadAPI                          string `json:"one_step_upload_api"`
	UploadSHA1API                             string `json:"upload_sha1_api"`
	OneFileRenameAPI                          string `json:"one_file_rename_api"`
	BatchFilesRenameAPI                       string `json:"batch_files_rename_api"`
	DeleteFileToTrashAPI                      string `json:"delete_file_to_trash_api"`
	CopyOneFileAPI                            string `json:"copy_one_file_api"`
	CopyBatchFilesAPI                         string `json:"copy_batch_files_api"`
	CopyBatchFilesProgressAPI                 string `json:"copy_batch_files_progress_api"`
	RecoverFileFromTrashAPI                   string `json:"recover_file_from_trash_api"`
	RecoverFileByPathAPI                      string `json:"recover_file_by_path_api"`
	GetOneFileDetailAPI                       string `json:"get_one_file_detail_api"`
	GetMultipleFilesDetailAPI                 string `json:"get_multiple_files_detail_api"`
	GetFileListAPI                            string `json:"get_file_list_api"`
	MoveFilesAPI                              string `json:"move_files_api"`
	DownloadFileAPI                           string `json:"download_file_api"`
	CreateShareLinkAPI                        string `json:"create_share_link_api"`
	GetShareLinkListAPI                       string `json:"get_share_link_list_api"`
	ModifyShareLinkAPI                        string `json:"modify_share_link_api"`
	CreatePaidShareLinkAPI                    string `json:"create_paid_share_link_api"`
	GetPaidShareLinkListAPI                   string `json:"get_paid_share_link_list_api"`
	ModifyPaidShareLinkAPI                    string `json:"modify_paid_share_link_api"`
	CreateOfflineDownloadMissionAPI           string `json:"create_offline_download_mission_api"`
	GetOfflineDownloadProgressAPI             string `json:"get_offline_download_progress_api"`
	GetUserInfoAPI                            string `json:"get_user_info_api"`
	SwitchIPBlacklistListAPI                  string `json:"switch_ip_blacklist_list_api"`
	UpdateIPBlacklistListAPI                  string `json:"update_ip_blacklist_list_api"`
	GetIPBlacklistListAPI                     string `json:"get_ip_blacklist_list_api"`
	GetDirectLinkOfflineLogsAPI               string `json:"get_direct_link_offline_logs_api"`
	GetDirectLinkTrafficLogsAPI               string `json:"get_direct_link_traffic_logs_api"`
	EnableDirectLinkAPI                       string `json:"enable_direct_link_api"`
	GetDirectLinkURLAPI                       string `json:"get_direct_link_url_api"`
	DisableDirectLinkAPI                      string `json:"disable_direct_link_api"`
	DirectLinkCacheRefreshAPI                 string `json:"direct_link_cache_refresh_api"`
	ImageHostCreateDirectoryAPI               string `json:"create_directory_oss_api"`
	ImageHostCreateFileAPI                    string `json:"create_file_oss_api"`
	ImageHostGetUploadURLAPI                  string `json:"get_upload_url_oss_api"`
	ImageHostUploadCompleteAPI                string `json:"upload_complete_oss_api"`
	ImageHostUploadAsyncResultAPI             string `json:"upload_async_result_oss_api"`
	ImageHostCreateCopyMissionAPI             string `json:"create_copy_mission_oss_api"`
	ImageHostGetCopyMissionDetailAPI          string `json:"get_copy_mission_detail_oss_api"`
	ImageHostGetCopyFailListAPI               string `json:"get_copy_fail_list_oss_api"`
	ImageHostMoveImageAPI                     string `json:"move_image_oss_api"`
	ImageHostDeleteImageAPI                   string `json:"delete_image_oss_api"`
	ImageHostGetImageDetailAPI                string `json:"get_image_detail_oss_api"`
	ImageHostGetImageListAPI                  string `json:"get_image_list_oss_api"`
	ImageHostCreateOfflineMigrationMissionAPI string `json:"create_offline_migration_mission_oss_api"`
	ImageHostGetOfflineMigrationMissionAPI    string `json:"get_offline_migration_mission_oss_api"`
}

// Response 通用API响应结构，包含响应状态码、消息和追踪ID
type Response struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	XTraceID string `json:"x-traceID"`
}

// APIClient 123云盘API客户端，包含认证信息和HTTP客户端
type APIClient struct {
	ClientID      string       `json:"clientID" yaml:"Client_id"`
	ClientSecret  string       `json:"clientSecret" yaml:"Client_secret"`
	AccessToken   string       `json:"accessToken" yaml:"access_token"`
	ExpiredAt     time.Time    `json:"expiredAt" yaml:"expired_at"`
	Authorization string       `json:"Authorization" yaml:"-"`
	HttpClient    *http.Client `json:"-" yaml:"-"`
	Config        *APIConfig   `json:"-" yaml:"-"`
}

// NewAPIClient 创建API客户端实例，支持从YAML配置文件或默认配置初始化
func NewAPIClient(configPath *string) (*APIClient, error) {
	var apiClient *APIClient
	var err error

	if configPath == nil || *configPath == "" {
		log.Println("using Default Config.")

		apiClient = &APIClient{
			Config:     &DefaultConfig,
			HttpClient: &http.Client{},
		}
	} else {
		log.Println("reading Config File.")
		apiClient, err = ReadAPIClientFromYamlFile(*configPath)
		if err != nil {
			return nil, err
		}
	}

	apiClient.Authorization = "Bearer " + apiClient.AccessToken
	apiClient.HttpClient = http.DefaultClient
	apiClient.Config = &DefaultConfig

	return apiClient, nil
}

// SaveToFile 将APIClient以yaml格式存储至文件
func (c *APIClient) SaveToFile(filePath string) error {
	headerStr, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	err = os.WriteFile(filePath, headerStr, 0666)
	if err != nil {
		return err
	}
	return nil
}

// DefaultConfig 默认API配置，使用123云盘开放平台的默认端点
var DefaultConfig = APIConfig{
	Domain:                                    "https://open-api.123pan.com",
	Platform:                                  "open_platform",
	AccessTokenAPI:                            "/api/v1/access_token",
	CreateFileAPI:                             "/upload/v2/file/create",
	CreateDirectoryAPI:                        "/upload/v1/file/mkdir",
	UploadSliceAPI:                            "/upload/v2/file/slice",
	UploadCompleteAPI:                         "/upload/v2/file/upload_complete",
	GetUploadDomainAPI:                        "/upload/v2/file/domain",
	OneStepUploadAPI:                          "/upload/v2/file/single/create",
	UploadSHA1API:                             "/upload/v2/file/sha1",
	OneFileRenameAPI:                          "/api/v1/file/name",
	BatchFilesRenameAPI:                       "/api/v1/file/rename",
	DeleteFileToTrashAPI:                      "/api/v1/file/trash",
	CopyOneFileAPI:                            "/api/v1/file/copy",
	CopyBatchFilesAPI:                         "/api/v1/file/async/copy",
	CopyBatchFilesProgressAPI:                 "/api/v1/file/async/copy/process",
	RecoverFileFromTrashAPI:                   "/api/v1/file/recover",
	RecoverFileByPathAPI:                      "/api/v1/file/recover/by_path",
	GetOneFileDetailAPI:                       "/api/v1/file/detail",
	GetMultipleFilesDetailAPI:                 "/api/v1/file/infos",
	GetFileListAPI:                            "/api/v2/file/list",
	MoveFilesAPI:                              "/api/v1/file/move",
	DownloadFileAPI:                           "/api/v1/file/download_info",
	CreateShareLinkAPI:                        "/api/v1/share/create",
	GetShareLinkListAPI:                       "/api/v1/share/list",
	ModifyShareLinkAPI:                        "/api/v1/share/list/info",
	CreatePaidShareLinkAPI:                    "/api/v1/share/content-payment/create",
	GetPaidShareLinkListAPI:                   "/api/v1/share/payment/list",
	ModifyPaidShareLinkAPI:                    "/api/v1/share/list/payment/info",
	CreateOfflineDownloadMissionAPI:           "/api/v1/offline/download",
	GetOfflineDownloadProgressAPI:             "/api/v1/offline/download/process",
	GetUserInfoAPI:                            "/api/v1/user/info",
	SwitchIPBlacklistListAPI:                  "/api/v1/developer/config/forbide-ip/switch",
	UpdateIPBlacklistListAPI:                  "/api/v1/developer/config/forbide-ip/update",
	GetIPBlacklistListAPI:                     "/api/v1/developer/config/forbide-ip/list",
	GetDirectLinkOfflineLogsAPI:               "/api/v1/direct-link/offline/logs",
	GetDirectLinkTrafficLogsAPI:               "/api/v1/direct-link/log",
	EnableDirectLinkAPI:                       "/api/v1/direct-link/enable",
	GetDirectLinkURLAPI:                       "/api/v1/direct-link/url",
	DisableDirectLinkAPI:                      "/api/v1/direct-link/disable",
	DirectLinkCacheRefreshAPI:                 "/api/v1/direct-link/cache/refresh",
	ImageHostCreateDirectoryAPI:               "/upload/v1/oss/file/mkdir",
	ImageHostCreateFileAPI:                    "/upload/v1/oss/file/create",
	ImageHostGetUploadURLAPI:                  "/upload/v1/oss/file/get_upload_url",
	ImageHostUploadCompleteAPI:                "/upload/v1/oss/file/upload_complete",
	ImageHostUploadAsyncResultAPI:             "/upload/v1/oss/file/upload_async_result",
	ImageHostCreateCopyMissionAPI:             "/api/v1/oss/source/copy",
	ImageHostGetCopyMissionDetailAPI:          "/api/v1/oss/source/copy/process",
	ImageHostGetCopyFailListAPI:               "/api/v1/oss/source/copy/fail",
	ImageHostMoveImageAPI:                     "/api/v1/oss/file/move",
	ImageHostDeleteImageAPI:                   "/api/v1/oss/file/delete",
	ImageHostGetImageDetailAPI:                "/api/v1/oss/file/detail",
	ImageHostGetImageListAPI:                  "/api/v1/oss/file/list",
	ImageHostCreateOfflineMigrationMissionAPI: "/api/v1/oss/offline/download",
	ImageHostGetOfflineMigrationMissionAPI:    "/api/v1/oss/offline/download/process",
}

# pan123API

封装了 [123云盘开放平台](https://open-api.123pan.com) API 的 Go 语言工具库。

## 安装

```bash
go get github.com/GhostiePie/pan123API
```

## 快速开始

```go
package main

import (
    "fmt"
    "log"

    "github.com/GhostiePie/pan123API/Client"
    "github.com/GhostiePie/pan123API/APIs/UserManagement"
)

func main() {
    // 创建API客户端
    configPath := "pan123config.yaml" // yaml定义见下文
    client, err := Client.NewAPIClient(&configPath)
    if err != nil {
        log.Fatal(err)
    }

    // 获取访问令牌
    accessTokenBody := UserManagement.GetAccessTokenBody{
        ClientID:     "your_client_id",
        ClientSecret: "your_client_secret",
    }
    accessTokenResp, err := UserManagement.GetAccessToken(client, accessTokenBody)
    if err != nil {
        log.Fatal(err)
    }
    client.AccessToken = accessTokenResp.Data.AccessToken
    client.ExpiredAt = accessTokenResp.Data.ExpiredAt
    client.Authorization = "Bearer " + client.AccessToken

    fmt.Printf("AccessToken: %s\n", client.AccessToken)
}
```

## 项目结构

```
├── Client/            # 核心客户端配置和HTTP工具
│   ├── Client.go      # APIClient、APIConfig、DefaultConfig
│   ├── HTTPUtils.go   # Get/Post/Put HTTP请求方法
│   └── ClientUtils.go # 工具函数和反序列化方法
├── APIs/              # API函数（按功能分包）
│   ├── Methods.go     # 高级封装方法（如SimpleUploadFile）
│   ├── FileManagement/ # 文件管理
│   │   ├── Upload/    # 文件上传（分片上传、秒传等）
│   │   ├── Copy/      # 文件复制
│   │   ├── Delete/    # 文件删除
│   │   ├── Detail/    # 文件详情查询
│   │   ├── List/      # 文件列表
│   │   ├── Recover/   # 文件恢复
│   │   ├── Rename/    # 文件重命名
│   │   ├── DownloadFile.go  # 下载链接获取
│   │   └── MoveFiles.go     # 文件移动
│   ├── ImageHost/     # 图片/OSS托管
│   │   ├── Upload/    # 图片上传
│   │   ├── Copy/      # 图片复制
│   │   ├── GetImageInfo/  # 图片信息查询
│   │   ├── OfflineMigration/ # 离线迁移
│   │   ├── MoveImage.go   # 图片移动
│   │   └── DeleteImage.go # 图片删除
│   ├── ShareManagement/   # 分享管理
│   ├── DirectLink/        # 直链管理
│   │   └── IPBlacklistConfiguration/ # IP黑名单
│   ├── OfflineDownloading/ # 离线下载
│   └── UserManagement/    # 用户信息
├── Utils/             # 工具函数（文件分片、MD5计算等）
├── TestUtils/         # 测试工具（配置加载、数据库等）
└── main.go            # 入口示例
```

## API 模块

### 文件管理 (FileManagement)

| 目录 | API | 说明 |
|------|-----|------|
| Upload | `CreateFile` | 创建文件上传任务 |
| Upload | `CreateDirectory` | 创建目录 |
| Upload | `GetUploadURL` | 获取上传域名 |
| Upload | `UploadSlice` | 上传文件分片 |
| Upload | `UploadComplete` | 完成上传 |
| Upload | `OneStepUpload` | 小文件一步上传 |
| Upload | `UploadSHA1` | SHA1秒传 |
| Copy | `CopyOneFile` | 复制单个文件 |
| Copy | `CopyBatchFiles` | 批量复制文件 |
| Copy | `CopyBatchFilesProgress` | 查询批量复制进度 |
| Delete | `DeleteFileToTrash` | 删除文件到回收站 |
| Detail | `GetOneFileDetail` | 获取单个文件详情 |
| Detail | `GetMultipleFilesDetail` | 批量获取文件详情 |
| List | `GetFileList` | 获取文件列表 |
| Recover | `RecoverFileFromTrash` | 从回收站恢复文件 |
| Recover | `RecoverFileByPath` | 按路径恢复文件 |
| Rename | `OneFileRename` | 重命名单个文件 |
| Rename | `BatchFilesRename` | 批量重命名文件 |
| - | `DownloadFile` | 获取文件下载链接 |
| - | `MoveFiles` | 移动文件 |

### 图片托管 (ImageHost)

| 目录 | API | 说明 |
|------|-----|------|
| Upload | `CreateFile` | 创建OSS上传任务 |
| Upload | `CreateDirectory` | 创建OSS目录 |
| Upload | `GetUploadURL` | 获取OSS上传URL |
| Upload | `UploadComplete` | 完成OSS上传 |
| Upload | `GetUploadAsyncResult` | 获取OSS异步上传结果 |
| Copy | `CreateCopyMission` | 创建OSS复制任务 |
| Copy | `GetCopyMissionDetail` | 查询OSS复制任务详情 |
| Copy | `GetCopyFailList` | 获取OSS复制失败列表 |
| GetImageInfo | `GetImageDetail` | 获取图片详情 |
| GetImageInfo | `GetImageList` | 获取图片列表 |
| OfflineMigration | `CreateOfflineMigrationMission` | 创建离线迁移任务 |
| OfflineMigration | `GetOfflineMigrationMission` | 查询离线迁移任务 |
| - | `MoveImage` | 移动图片 |
| - | `DeleteImage` | 删除图片 |

### 分享管理 (ShareManagement)

| API | 说明 |
|-----|------|
| `CreateShareLink` | 创建分享链接 |
| `GetShareLinkList` | 获取分享链接列表 |
| `ModifyShareLink` | 修改分享链接 |
| `CreatePaidShareLink` | 创建付费分享链接 |
| `GetPaidShareLinkList` | 获取付费分享列表 |
| `ModifyPaidShareLink` | 修改付费分享链接 |

### 直链管理 (DirectLink)

| API | 说明 |
|-----|------|
| `EnableDirectLink` | 启用直链 |
| `DisableDirectLink` | 禁用直链 |
| `GetDirectLinkURL` | 获取直链URL |
| `GetDirectLinkTrafficLogs` | 获取直链流量日志 |
| `GetDirectLinkOfflineLogs` | 获取直链离线日志 |
| `DirectLinkCacheRefresh` | 刷新直链缓存 |

#### IP黑名单 (IPBlacklistConfiguration)

| API | 说明 |
|-----|------|
| `GetIPBlacklistList` | 获取IP黑名单列表 |
| `SwitchIPBlacklistList` | 切换IP黑名单开关 |
| `UpdateIPBlacklistList` | 更新IP黑名单列表 |

### 离线下载 (OfflineDownloading)

| API | 说明 |
|-----|------|
| `CreateOfflineDownloadMission` | 创建离线下载任务 |
| `GetOfflineDownloadProgress` | 获取离线下载进度 |

### 用户管理 (UserManagement)

| API | 说明 |
|-----|------|
| `GetAccessToken` | 获取访问令牌 |
| `GetUserInfo` | 获取用户信息 |

## 配置

配置通过 YAML 文件加载：

```yaml
# 请务必保持该文件独立于其他配置文件，即为该SDK准备专门的配置文件，因自动更新token会写入该文件
client_id: "your_client_id" # 必填
client_secret: "your_client_secret" # 必填
access_token: "your_access_token" # 非必填，NewAPIClient()会自动对比token是否过期并重新获取，然后写入该文件
expired_at: "2024-01-01T00:00:00Z"
```

```go
configPath := "pan123config.yaml"
client, err := Client.NewAPIClient(&configPath)
```

也可以通过 `Client.APIConfig` 结构体进行完全自定义。

## 构建

```bash
go build ./...
go fmt ./...
go vet ./...
```


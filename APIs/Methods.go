package APIs

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/GhostiePie/pan123API/APIs/FileManagement/Upload"
	"github.com/GhostiePie/pan123API/Client"
	"github.com/GhostiePie/pan123API/Utils"
)

func SimpleUploadFile(client *Client.APIClient, filePath string, parentFileID int, fileName string) error {

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	fileStat, err := file.Stat()

	fileMd5 := Utils.CalcFileMD5(file)
	// 重置指针到开头
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}

	createFileBody := Upload.CreateFileBody{
		ParentFileID: parentFileID,
		FileName:     fileName,
		Etag:         fileMd5,
		Size:         int(fileStat.Size()),
		Duplicate:    2,
	}
	createFileResp, err := Upload.CreateFile(client, createFileBody)
	if err != nil {
		return err
	}
	fmt.Println(createFileResp)

	if createFileResp.Data.Reuse {
		return nil
	} else {
		// 创建 MD5 哈希对象
		hasher := md5.New()

		// 使用 TeeReader: 读取的数据同时写入 hasher
		teeReader := io.TeeReader(file, hasher)

		buffer := make([]byte, createFileResp.Data.SliceSize)
		chunkIndex := 0
		sliceCount := fileStat.Size()/int64(createFileResp.Data.SliceSize) + 1
		fmt.Printf("ChunkSize: %dMB*%d\n", createFileResp.Data.SliceSize/1024/1024, sliceCount)
		for {
			chunkIndex++
			hasher.Reset()
			n, err := teeReader.Read(buffer)

			if err != nil && err != io.EOF {
				return err
			}

			if n > 0 {
				// 处理当前分片（比如发送到 channel、上传等）
				uploadSliceBody := Upload.UploadSliceBody{
					PreUploadID: createFileResp.Data.PreUploadID,
					SliceNo:     chunkIndex,
					SliceMD5:    fmt.Sprintf("%x", hasher.Sum(nil)),
					Slice:       buffer[:n],
					Servers:     createFileResp.Data.Servers,
				}
				uploadSliceResp, err := Upload.UploadSlice(client, uploadSliceBody)
				if uploadSliceResp.Code != 0 {
					return errors.New(uploadSliceResp.Message)
				}
				if err != nil {
					return err
				}
				fmt.Printf("%v/%v\n", chunkIndex, sliceCount)
			}
			if err == io.EOF {
				break
			}

		}

		uploadCompleteBody := Upload.UploadCompleteBody{
			PreuploadID: createFileResp.Data.PreUploadID,
		}
		for {
			time.Sleep(2 * time.Second)
			uploadCompleteResp, err := Upload.UploadComplete(client, uploadCompleteBody)
			if err != nil {
				return err
			}
			fmt.Println(uploadCompleteResp)
			if uploadCompleteResp.Code != 20103 {
				break
			}
		}

	}
	return nil
}

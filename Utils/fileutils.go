package Utils

import (
	"fmt"
	"io"
	"math"
	"os"
)

// SplitFile 用于从某个特定文件中一次性读出多个分片，会将文件一次性全部加载到内存，慎用。
func SplitFile(fileName string, chunkSize int) ([][]byte, error) {
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		return nil, err
	}

	// 由于后期会将可接受文件体积限制在2GB以下，因此可以放心转为int以简化后期计算，不用担心fileSize超出int范围
	fileSize := int(fileInfo.Size())

	if fileSize > int(math.Pow(2, 30)) { // 如果文件大于2GB
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

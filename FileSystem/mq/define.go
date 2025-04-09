package mq

type TransferData struct {
	FileHash    string `json:"fileHash"`
	FileName    string `json:"fileName"`
	FileSize    int64  `json:"fileSize"`
	CephBucket  string `json:"ceph_Bucket"`
	CephPath    string `json:"ceph_Path"`
	TempPath    string `json:"tempPath"`
	UserID      string `json:"user_id"`      // 上传用户标识
	UploadTime  string `json:"upload_time"`  // ISO8601格式(如"2023-08-20T15:04:05Z")
	ContentType string `json:"content_type"` // 文件MIME类型(如"application/pdf")
}

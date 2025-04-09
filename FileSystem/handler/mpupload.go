package handler

import (
	rPoll "FileSystem/cache/redis"
	dblayer "FileSystem/db"
	utils "FileSystem/utils"
	"context"
	"fmt"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

// MultipartUploadInfo : 初始化信息
type MultipartUploadInfo struct {
	FileHash   string
	FileSize   int
	UploadID   string
	ChunkSize  int
	ChunkCount int
}

// InitialMultipartUploadHandler : 初始化分块上传
func InitialMultipartUploadHandler(w http.ResponseWriter, r *http.Request) {
	//解析用户请求信息
	r.ParseForm()
	username := r.Form.Get("username")
	filehash := r.Form.Get("filehash")
	filesize, err := strconv.Atoi(r.Form.Get("filesize"))
	if err != nil {
		w.Write(utils.NewRespMsg(-1, "params invalid", nil).JSONBytes())
		return
	}
	//获得redis的一个连接
	rConn := rPoll.RedisClient().Conn()
	defer rConn.Close()
	//生成分块上传的初始化信息
	upInfo := MultipartUploadInfo{
		FileHash:   filehash,
		FileSize:   filesize,
		UploadID:   username + fmt.Sprintf("%x", time.Now().UnixNano()),
		ChunkSize:  5 * 1024 * 1024, //5MB
		ChunkCount: int(float64(filesize) / (5 * 1024 * 1024)),
	}
	//将初始化信息写入到redis缓存
	rConn.HSet(context.Background(), "MP_"+upInfo.UploadID, "chunkcount", upInfo.ChunkCount)
	rConn.HSet(context.Background(), "MP_"+upInfo.UploadID, "filehash", upInfo.FileHash)
	rConn.HSet(context.Background(), "MP_"+upInfo.UploadID, "filesize", upInfo.FileSize)
	//将相应初始化数据返回到客户端
	w.Write(utils.NewRespMsg(0, "OK", upInfo).JSONBytes())
}

// UploadPartHandler : 上传文件分块
func UploadPartHandler(w http.ResponseWriter, r *http.Request) {
	//解析用户请求
	r.ParseForm()
	//username := r.Form.Get("username")
	uploadID := r.Form.Get("uploadid")
	chunkIndex := r.Form.Get("index")
	//获得redis连接池中的一个连接
	rConn := rPoll.RedisClient().Conn()
	defer rConn.Close()
	//获得文件句柄，用于存储分块内容
	fpath := "D:\\GoLearn\\FileSystem\\tmp\\" + uploadID + "\\" + chunkIndex
	os.MkdirAll(path.Dir(fpath), 0744)
	fd, err := os.Create(fpath)
	if err != nil {
		w.Write(utils.NewRespMsg(-1, "Upload part failed", nil).JSONBytes())
		return
	}
	defer fd.Close()
	//将分块数据读入到文件当中
	buf := make([]byte, 1024*1024)
	for {
		n, err := r.Body.Read(buf)
		fd.Write(buf[:n])
		if err != nil {
			break
		}
	}
	//更新redis缓存状态
	rConn.HSet(context.Background(), "MP_"+uploadID, "chunkidx_"+chunkIndex, 1)
	//返回处理结果到客户端
	w.Write(utils.NewRespMsg(0, "OK", nil).JSONBytes())
}

// CompleteUploadHandler : 通知上传合并接口
func CompleteUploadHandler(w http.ResponseWriter, r *http.Request) {
	//解析请求参数
	r.ParseForm()
	uploadID := r.Form.Get("uploadid")
	username := r.Form.Get("username")
	filehash := r.Form.Get("filehash")
	filesize, _ := strconv.Atoi(r.Form.Get("filesize"))
	filename := r.Form.Get("filename")
	//获得redis连接池中的一个连接
	rConn := rPoll.RedisClient().Conn()
	defer rConn.Close()
	//通过uploadid查询redis并判断是否所有分块上传完成
	data, err := rConn.HGetAll(context.Background(), "MP_"+uploadID).Result()
	if err != nil {
		w.Write(utils.NewRespMsg(-1, "complete upload failed", nil).JSONBytes())
		return
	}
	totalCount, err := strconv.Atoi(data["chunkcount"])
	if err != nil {
		w.Write(utils.NewRespMsg(-1, "无效的分块数格式", nil).JSONBytes())
		return
	}
	chunkCount := 0
	for k, v := range data {
		if strings.HasPrefix(k, "chkidx_") && v == "1" {
			chunkCount++
		}
	}
	if totalCount != chunkCount {
		w.Write(utils.NewRespMsg(-2, fmt.Sprintf("分块不完整(需要%d个，已传%d个)", totalCount, chunkCount), nil).JSONBytes())
		return
	}
	//TODO：合并分块

	//更新唯一文件表及用户文件表
	dblayer.OnFileUploadFinished(filehash, filename, int64(filesize), "")
	dblayer.OnUserFileUploadFinished(username, filehash, filename, int64(filesize))

	// 6. 响应处理结果
	w.Write(utils.NewRespMsg(0, "OK", nil).JSONBytes())
}

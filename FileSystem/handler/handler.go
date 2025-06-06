package handler

import (
	dblayer "FileSystem/db"
	"FileSystem/meta"
	"FileSystem/store/ceph"
	utils "FileSystem/utils"
	"encoding/json"
	"fmt"
	"gopkg.in/amz.v1/s3"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

// UploadHandler ：处理文件上传
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//返回上传的html页面
		data, err := os.ReadFile("./static/view/index.html")
		if err != nil {
			io.WriteString(w, "internel server error")
			return
		}
		io.WriteString(w, string(data))
	} else if r.Method == "POST" {
		//接收文件流及存储到本地目录
		//接收流文件
		file, head, err := r.FormFile("file")
		if err != nil {
			fmt.Printf("Failed to get file, err:%s\n", err)
			return
		}
		defer file.Close()

		fileMeta := meta.FileMeta{
			FileName: head.Filename,
			Location: "D:\\GoLearn\\FileSystem\\tmp\\" + head.Filename,
			UploadAt: time.Now().Format("2006-01-02 15:04:05"),
		}
		//创建文件
		newFile, err := os.Create(fileMeta.Location)
		if err != nil {
			fmt.Printf("Failed to create file, err:%s\n", err)
			return
		}
		defer newFile.Close()
		//将流文件拷贝到文件中去
		fileMeta.FileSize, err = io.Copy(newFile, file)
		if err != nil {
			fmt.Printf("Failed to save data into file, err:%s\n", err)
			return
		}
		newFile.Seek(0, 0)
		fileMeta.FileSha1 = utils.FileSha1(newFile)
		//同时将文件写入ceph存储
		newFile.Seek(0, 0)
		data, _ := io.ReadAll(newFile)
		bucket := ceph.GetCephBucket("userfile")
		cephPath := "/ceph/" + fileMeta.FileSha1
		_ = bucket.Put(cephPath, data, "octet-stream", s3.PublicRead)
		fileMeta.Location = cephPath
		//meta.UpdateFileMeta(fileMeta)
		_ = meta.UpdateFileMetaDB(fileMeta)
		r.ParseForm()
		username := r.Form.Get("username")
		res := dblayer.OnUserFileUploadFinished(username, fileMeta.FileSha1, fileMeta.FileName, fileMeta.FileSize)
		if res {
			http.Redirect(w, r, "/static/view/home.html", http.StatusFound)
		} else {
			w.Write([]byte("Upload Failed"))
		}

		//4、重定向
		http.Redirect(w, r, "/file/upload/success", http.StatusFound)
	}
}

// UpLoadSuccessHandler : 上传已完成
func UpLoadSuccessHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "文件上传成功！")
}

// GetFileMetaHandler : 获取文件元信息
func GetFileMetaHandler(w http.ResponseWriter, r *http.Request) {
	//解析请求
	r.ParseForm()
	filehash := r.Form.Get("filehash")
	//从元数据系统查询文件信息
	fMeta, err := meta.GetFileMetaDB(filehash)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//将元数据转为json
	data, err := json.Marshal(fMeta)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//返回json数据
	w.Write(data)
}

// DownloadHandler :
func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	//解析数据
	r.ParseForm()
	fsha1 := r.Form.Get("filehash")
	//获取元数据
	fm := meta.GetFileMeta(fsha1)

	//通过元数据打开该文件
	f, err := os.Open(fm.Location)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()

	//读取该文件
	data, err := io.ReadAll(f)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//将该文件返回
	//设置响应头
	w.Header().Set("content-type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename=\""+fm.FileName+"\"")
	w.Write(data)
}

// FileMetaUpdateHandler : 更新元信息接口（重命名）
func FileMetaUpdateHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	opType := r.Form.Get("op")
	fileSha1 := r.Form.Get("filehash")
	newFileName := r.Form.Get("fileName")
	if opType != "0" {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	curFileMeta := meta.GetFileMeta(fileSha1)
	curFileMeta.FileName = newFileName
	meta.UpdateFileMeta(curFileMeta)

	data, err := json.Marshal(curFileMeta)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
	w.WriteHeader(http.StatusOK)
}

// FileDeleteHandler : 删除文件及元信息
func FileDeleteHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fileSha1 := r.Form.Get("filehash")

	//删除文件
	fMeta := meta.GetFileMeta(fileSha1)
	os.Remove(fMeta.Location)

	//删除元信息
	meta.RemoveFileMeta(fileSha1)

	w.WriteHeader(http.StatusOK)
}

// FileQueryHandler : 查询批量文件元信息
func FileQueryHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	limitCut, _ := strconv.Atoi(r.Form.Get("limit"))

	//fileMetas := meta.GetLastFileMetas(limitCut)
	username := r.Form.Get("username")
	userFiles, err := dblayer.QueryUserFileMetas(username, limitCut)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(userFiles)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

// TryFastUploadHandler : 尝试秒传接口
func TryFastUploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	//解析请求参数
	username := r.Form.Get("username")
	filehash := r.Form.Get("filehash")
	filename := r.Form.Get("filename")
	filesize := r.Form.Get("filesize")
	//从文件表中查询相同hash的文件记录
	fileMeta, err := meta.GetFileMetaDB(filehash)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//查不到记录则返回秒传失败
	if fileMeta == nil {
		resp := utils.RespMsg{
			Code: -1,
			Msg:  "秒传失败，请访问普通上传接口",
		}
		w.Write(resp.JSONBytes())
		return
	}
	//上传过则将文件信息写入用户文件表，返回成功
	size, _ := strconv.ParseInt(filesize, 10, 64)
	suc := dblayer.OnUserFileUploadFinished(username, filehash, filename, size)
	if suc {
		resp := utils.RespMsg{
			Code: 0,
			Msg:  "秒传成功",
		}
		w.Write(resp.JSONBytes())
		return
	} else {
		resp := utils.RespMsg{
			Code: -2,
			Msg:  "秒传失败， 请稍后重试",
		}
		w.Write(resp.JSONBytes())
		return
	}
}

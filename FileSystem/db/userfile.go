package db

import (
	mydb "FileSystem/db/mysql"
	"fmt"
	"time"
)

// UserFile : 用户文件表结构
type UserFile struct {
	UserName    string
	FileHash    string
	FileSize    int64
	FileName    string
	UploadAt    string
	LastUpdated string
}

// OnUserFileUploadFinished : 更新用户文件表
func OnUserFileUploadFinished(username, filehash, filename string, fileSize int64) bool {
	stmt, err := mydb.DBConn().Prepare(
		"insert ignore into tbl_user_file(`user_name`, `file_sha1`, `file_name`, `file_size`, `upload_at`) values (?,?,?,?,?)",
	)
	if err != nil {
		return false
	}
	defer stmt.Close()
	_, err = stmt.Exec(username, filehash, filename, fileSize, time.Now())
	if err != nil {
		return false
	}
	return true
}

// QueryUserFileMetas : 批量获取用户文件信息
func QueryUserFileMetas(username string, limit int) ([]UserFile, error) {
	stmt, err := mydb.DBConn().Prepare(
		"select file_sha1, file_name, file_size, upload_at, last_update from tbl_user_file where user_name=? limit ?",
	)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(username, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var userFiles []UserFile
	for rows.Next() {
		userFile := UserFile{}
		err := rows.Scan(&userFile.FileHash, &userFile.FileName, &userFile.FileSize, &userFile.UploadAt, &userFile.LastUpdated)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		userFiles = append(userFiles, userFile)
	}
	return userFiles, nil
}

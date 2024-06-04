package util

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

/*
获取当前文件夹和全部子文件夹下视频文件
*/

func GetAllFiles(root string) (files []string) {
	patterns := []string{"webm", "m4v", "mp4", "mov", "avi", "wmv", "ts", "rmvb", "wma", "avi", "flv", "rmvb", "mpg", "f4v", "mkv"}
	for _, pattern := range patterns {
		files = append(files, getFilesByExtension(root, pattern)...)
	}
	return files
}

/*
获取当前文件夹和全部子文件夹下指定扩展名的全部文件
*/
func getFilesByExtension(root, extension string) []string {
	var files []string
	defer func() {
		if err := recover(); err != nil {
			log.Println("获取文件出错")
			os.Exit(-1)
		}
	}()
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && strings.HasSuffix(info.Name(), extension) {
			files = append(files, path)
		}
		return nil
	})
	return files
}
func IsExist(fp string) bool {
	// 使用 os.Stat 函数获取文件信息
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("错误信息:%v\n", err)
		}
	}()
	if f, err := os.Stat(fp); err == nil {
		fmt.Println("Path exists")
		if f.IsDir() {
			fmt.Println("Path is a directory")
		} else {
			fmt.Println("Path is a file")
		}
		return true
	} else if os.IsNotExist(err) {
		fmt.Println("Path does not exist")
		return false
	} else {
		fmt.Println("Error occurred:", err)
		return false
	}
}

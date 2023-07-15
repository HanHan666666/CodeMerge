package internal

import (
	"log"
	"os"
	"path/filepath"
)

// GetAllFileIncludeSubFolder 递归获取某个目录下的所有文件
func GetAllFileIncludeSubFolder(folder string) ([]string, error) {
	var result []string

	err := filepath.Walk(folder, func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			log.Println(err.Error())
			return err
		}

		if fi.IsDir() {
			// 如果这个目录是.开头的，说明是一个隐藏目录，忽略这个目录
			if fi.Name()[0] == '.' {
				return filepath.SkipDir
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}

func MergeFiles(outFile string, files []string) error {
	log.Println("合并文件数量:", len(files))

	for _, file := range files {
		log.Println("合并文件:", file)
		// 读取文件内容
		content, err := ReadFile(file)
		if err != nil {
			log.Println("读取文件失败:", err)
			continue
		}
		dir, _ := os.Getwd()
		//提取file中的相对路径和文件名
		relPath, _ := filepath.Rel(dir, file)

		// 把relPath追加到content开头
		content = []byte(relPath + "\n" + string(content))
		// content后面加上两个换行
		content = append(content, []byte("\n\n")...)

		// 追加写入目标文件
		err = WriteFile(outFile, content)
		if err != nil {
			return err
		}
	}
	return nil
}

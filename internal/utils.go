package internal

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"strings"
	"unicode"
)

// isTextFile 判断是不是纯文本文件
func isTextFile(filename string) bool {
	file, err := os.Open(filename)
	if err != nil {
		return false
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for _, char := range line {
			if !unicode.IsPrint(char) && !unicode.IsSpace(char) {
				return false
			}
		}
	}

	return true
}

// ReadFile 读取文件内容
func ReadFile(filename string) ([]byte, error) {
	// 判断是不是纯文本文件
	if !isTextFile(filename) {
		log.Printf("%s is not a text file", filename)
		return nil, errors.New("not a text file")
	}
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	content, err := io.ReadAll(f)

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)

	return content, err
}

// checkFileIsExist 检查文件是否存在
func checkFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

// WriteFile 写入文件
func WriteFile(filename string, content []byte) error {
	var f *os.File
	var err error
	if checkFileIsExist(filename) { //如果文件存在
		f, err = os.OpenFile(filename, os.O_APPEND, 0666) //打开文件
		if err != nil {
			return err
		}
		log.Println("文件存在")
	} else {
		f, err = os.Create(filename) //创建文件
		if err != nil {
			return err
		}
		log.Println("文件不存在")
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)

	_, err = f.WriteString(string(content))
	return err
}

// FilterFiles 根据忽略列表过滤文件名
func FilterFiles(files, ignores []string) []string {

	var filtered []string

	for _, file := range files {
		ignored := false
		for _, ignore := range ignores {
			// if file == ignore {
			// 	ignored = true
			// 	break
			// }

			// 如果file这个目录中包含ignore这个目录，那么就忽略这个目录
			if strings.Contains(file, ignore) {
				ignored = true
				break
			}
		}
		if !ignored {
			filtered = append(filtered, file)
		}
	}

	return filtered
}

// IsIgnoreFileExist 检查是否有 .gitignore 文件
func IsIgnoreFileExist(folder string) bool {
	_, err := os.Stat(folder + "/.gitignore")
	return err == nil
}

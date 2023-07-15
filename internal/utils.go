package internal

import (
	"bufio"
	"errors"
	ignore "github.com/sabhiram/go-gitignore"
	"io"
	"log"
	"os"
	"path/filepath"
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

	// 获取当前工作目录
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	// 将文件名转换为相对当前工作目录的路径
	filename, err = filepath.Rel(cwd, filename)
	if err != nil {
		return nil, err
	}

	content = []byte("\n// " + filename + " BEGIN\n" +
		string(content) +
		"// " + filename + " END\n")

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
		//log.Println("文件存在")
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

// FilterFiles 过滤掉.gitignore忽略掉的文件
func FilterFiles(files []string, ignoreFile string) ([]string, error) {

	gi, err := ignore.CompileIgnoreFile(ignoreFile)
	if err != nil {
		return nil, err
	}

	var filtered []string
	for _, file := range files {
		if !gi.MatchesPath(file) {
			filtered = append(filtered, file)
		}
	}
	return filtered, nil
}

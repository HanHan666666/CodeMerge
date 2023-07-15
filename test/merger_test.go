package test

import (
	"CodeMerge/internal"
	"bufio"
	"fmt"
	"os"
	"testing"
	"unicode"
)

func TestReadGitIgnore(t *testing.T) {

	// 读取 .gitignore 文件
	ignoreFiles, _ := internal.ReadFile("D:\\looook\\CodeMerge\\.gitignore")
	fmt.Printf("ignoreFiles: %v\n", string(ignoreFiles))

}

func TestMerge(t *testing.T) {
	// 测试代码
	// fmt.Println(internal.GetAllFileIncludeSubFolder("D:\\looook\\CodeMerge"))
	x, _ := internal.GetAllFileIncludeSubFolder("D:\\looook\\CodeMerge")
	// fmt.Printf("x: %v\n", x)
	for _, v := range x {
		fmt.Printf("v: %v\n", v)
	}
	// ignore := []string{".git"}
	// fmt.Printf("\"\": %v\n", "过滤后的内容：")
	// c := internal.FilterFiles(x, ignore)
	// fmt.Printf("c: %v\n", c)

	// 读取文件内容
	// content, _ := internal.ReadFile(c[0])
	// // 输出文件内容文本
	// fmt.Printf("content: %v\n", string(content))
}

func TestGetPAth(t *testing.T) {
	internal.GetAllFileIncludeSubFolder("D:\\looook\\CodeMerge")
}

func isTextFile(filename string) bool {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("无法打开文件:", err)
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

func TestFileClass(t *testing.T) {
	filename := "C:\\Users\\wuuuu\\Downloads\\aiqimeng.sql"
	if isTextFile(filename) {
		fmt.Println(filename, "是纯文本文件")
	} else {
		fmt.Println(filename, "不是纯文本文件")
	}
}

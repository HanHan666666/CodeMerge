package test

import (
	"CodeMerge/internal"
	"bufio"
	"fmt"
	ignore "github.com/sabhiram/go-gitignore"
	"log"
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
		log.Printf("v: %v\n", v)
	}
	//x, err := internal.FilterFiles(x)
	//log.Println("过滤后的内容：")
	//for _, v := range x {
	//	log.Printf("v: %v\n", v)
	//}
	//if err != nil {
	//	log.Printf("err: %v\n", err)
	//}
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

func TestGitignore(t *testing.T) {
	dir, _ := os.Getwd()
	ignoreFile := dir + "\\..\\" + ".gitignore"
	gi, err := ignore.CompileIgnoreFile(ignoreFile)
	if err != nil {
		log.Printf("Error: %s\n", err)
	}
	//internal.FilterFiles("CodeMerged.txt")
	log.Println(gi.MatchesPath(".idea/vcs.xml"))
}

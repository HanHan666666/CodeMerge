package main

import (
	"CodeMerge/internal"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var VERSION = "DEV-1.0"
var GOVERSION = "UNKNOWN"
var GITCOMMIT = "UNKNOWN*"
var BUILDTIME = "UNKNOWN"

func main() {
	dir, _ := os.Getwd()

	// 定义命令行参数
	srcDir := flag.String("src", dir, "源代码目录(绝对路径)")
	outFile := flag.String("out", "CodeMerged.txt", "输出文件路径")
	//输出版本号
	showVersion := flag.Bool("version", false, "显示版本号")
	// orderFiles := flag.String("order", "", "排序文件的名字")
	// onlyFiles := flag.String("only", "", "仅合并特定文件的名字")

	flag.Parse()

	ShowVersion(showVersion)

	//输出当前所在目录
	fmt.Println("当前目录:", dir)

	// 获取源代码目录下的文件列表
	files, err := internal.GetAllFileIncludeSubFolder(*srcDir)
	if err != nil {
		fmt.Println("获取文件列表失败:", err)
		os.Exit(1)
	}

	// 检查是否有 .gitignore 文件
	ignoreFile := filepath.Join(dir, ".gitignore")
	_, err = os.Stat(ignoreFile)
	// 如果存在 .gitignore 文件，则过滤要排除的文件
	if err == nil {
		log.Println("找到 .gitignore 文件, 将根据 .gitignore 文件过滤文件")
		// 过滤要排除的文件
		files, _ = internal.FilterFiles(files, ignoreFile)
	} else {
		log.Println("没有找到 .gitignore 文件, 将不会根据 .gitignore 文件过滤文件")
	}

	// TODO
	// // 按orderFiles排序
	// if *orderFiles != "" {
	// 	fileOperate = internal.OrderFiles(fileOperate, strings.Split(*orderFiles, ","))
	// }

	// // 仅保留onlyFiles指定的文件
	// if *onlyFiles != "" {
	// 	fileOperate = internal.OnlyFiles(fileOperate, strings.Split(*onlyFiles, ","))
	// }

	// 执行合并
	err = internal.MergeFiles(*outFile, files)
	if err != nil {
		fmt.Println("合并失败:", err)
		os.Exit(1)
	}

	fmt.Println("合并成功!")
}

func version() string {
	return fmt.Sprintf(`CodeMerge %s
Git commit: %s,
Build with %s
Build at %s`, VERSION, GITCOMMIT, GOVERSION, BUILDTIME)
}

func ShowVersion(showVersion *bool) {
	if *showVersion {
		fmt.Println(version())
		os.Exit(0)
	}
}

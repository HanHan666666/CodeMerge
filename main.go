package main

import (
	"CodeMerge/internal"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	//输出当前所在目录
	dir, _ := os.Getwd()
	fmt.Println("当前目录:", dir)

	// 定义命令行参数
	srcDir := flag.String("src", dir, "源代码目录(绝对路径)")
	outFile := flag.String("out", "CodeMerged.txt", "输出文件路径")
	// orderFiles := flag.String("order", "", "排序文件的名字")
	// onlyFiles := flag.String("only", "", "仅合并特定文件的名字")

	flag.Parse()

	// 获取源代码目录下的文件列表
	files, err := internal.GetAllFileIncludeSubFolder(*srcDir)
	if err != nil {
		fmt.Println("获取文件列表失败:", err)
		os.Exit(1)
	}
	log.Printf("获取文件列表成功, 文件数量: %d\n", len(files))
	log.Println("文件列表:", files)

	// 检查是否有 .gitignore 文件
	//var ignoreFiles []string
	//if internal.IsIgnoreFileExist(*srcDir) {
	//	ignorePAth := *srcDir + "/.gitignore"
	//	// 读取 .gitignore 文件
	//	ignoreBytes, _ := internal.ReadFile(ignorePAth)
	//	ignoreFiles = strings.Split(string(ignoreBytes), "\n")
	//}

	//过滤要排除的文件
	//files = internal.FilterFiles(files, ignoreFiles)

	// TODO
	// // 按orderFiles排序
	// if *orderFiles != "" {
	// 	files = internal.OrderFiles(files, strings.Split(*orderFiles, ","))
	// }

	// // 仅保留onlyFiles指定的文件
	// if *onlyFiles != "" {
	// 	files = internal.OnlyFiles(files, strings.Split(*onlyFiles, ","))
	// }

	// 执行合并
	err = internal.MergeFiles(*outFile, files)
	if err != nil {
		fmt.Println("合并失败:", err)
		os.Exit(1)
	}

	fmt.Println("合并成功!")
}

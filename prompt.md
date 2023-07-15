你现在是一个 Go 语言程序员，我想让你帮我实现一个项目，下面是我的项目的详细说明

# CodeMerge

A tool to merge project code into one txt.

一个将项目代码合并成一个 txt 的工具

将项目目录下的所有的代码合并成一个 txt 文件

功能
根据.gitignore 忽略文件

仅合并指定的若干个文件

自定义合并后的顺序

通过指定目录

通过指定文件名，未指定的默认顺序

当文件超过 10M 的时候拆分

## 功能详细说明

软件采用开箱即用的策略，就是在终端中输入 codemerge 然后触发默认行为：首先会默认寻找当前目录下的`.gitignore`文件，排除需要忽略的文件， 然后遍历当前目录下所有的文本文件，不论后缀格式，合并到 codemerge.txt 并保存到当前目录下。 如果没有找到`.gitignore`文件，则直接遍历所有的 txt，执行合并； 如果有命令行参数，命令行参数可以指定要合并的目录，个别文件的顺序（--order 后面跟若干个文件名，把这些文件移到 codemerge.txt 的开头） 也可以通过命令行参数来指定仅合并哪些文件。

codemerge.txt中的内容模板是下面这个样子
txt模板：
```txt
reademe.md
content...

main.java
code...

util.java
code...

```

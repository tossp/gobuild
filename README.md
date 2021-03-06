gobuild [![Build Status](https://travis-ci.org/caixw/gobuild.svg?branch=master)](https://travis-ci.org/caixw/gobuild)
======

gobuild是一个简单的Go代码热编译工具。
会实时监控指定目录下的文件变化(重命名，删除，创建，添加)，并编译和运行程序。


#### 命令行语法:
```shell
gobuild [options] [dependents]

options:
 -h    显示当前帮助信息；
 -v    显示gobuild和go程序的版本信息；
 -o    执行编译后的可执行文件名；
 -r    是否搜索子目录，默认为true；
 -i    是否显示被标记为IGNORE的日志内容，默认为false，即不显示；
 -ext  需要监视的扩展名，默认值为"go"，区分大小写，会去掉每个扩展名的首尾空格。
       若需要监视所有类型文件，请使用*，传递空值代表不监视任何文件；
 -main 指定需要编译的文件，默认为""。

dependents:
 指定其它依赖的目录，只能出现在命令的尾部。
```


#### 常见用法:
```go
 // 监视当前目录下的文件，若发生变化，则触发go build -main="*.go"
 gobuild

 // 监视当前目录和~/Go/src/github.com/issue9/term目录下的文件，
 // 若发生变化，则触发go build -main="main.go"
 gobuild -main=main.go ~/Go/src/github.com/issue9/term
```


#### 支持平台:

平台支持依赖[colors](https://github.com/issue9/term)与[fsnotify](https://gopkg.in/fsnotify.v1)两个包，
目前支持以下平台：Windows, Linux, OSX, BSD。


### 安装

```shell
go get github.com/caixw/gobuild
```


### 版权

本项目采用[MIT](http://opensource.org/licenses/MIT)开源授权许可证，完整的授权说明可在[LICENSE](LICENSE)文件中找到。

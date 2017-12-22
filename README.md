
## GO 语言学习

### 资源网站

[awesome-go](https://awesome-go.com/)


### 安装

```
brew install go
```

安装完成可使用
```
// 查询版本
go version
//查询环境变量
go env
```



### 设置 $GOPATH

 `$GOPATH`其实就一个工作目录，其下有三个文件夹。

- bin 存储编译后的可执行文件
- pkg 存放编译后生成的包文件
- src 存放项目的源码（即项目）
- 

我使用的是zsh,所以在桌面下设置一个golang的工作目录

编辑配置项

```
vim ~/.zshrc
```

设置`$GOPATH`，添加两行

```
export GOROOT=/usr/local/Cellar/go/1.9.2/libexec
export GOPATH=/Users/xcmy/desktop/golang/
```

保存退出，是配置项生效

```
source ~/.zshrc
```
##### 重新打开终端，查看go环境变量

```
~ » go env                                 
GOARCH="amd64"
GOBIN=""
GOEXE=""
GOHOSTARCH="amd64"
GOHOSTOS="darwin"
GOOS="darwin"
GOPATH="/Users/xcmy/desktop/golang/"
GORACE=""
GOROOT="/usr/local/Cellar/go/1.9.2/libexec"
GOTOOLDIR="/usr/local/Cellar/go/1.9.2/libexec/pkg/tool/darwin_amd64"
GCCGO="gccgo"
CC="clang"
GOGCCFLAGS="-fPIC -m64 -pthread -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=/var/folders/5h/hm_zf0rx3pl9rlrl5m7m3sl40000gp/T/go-build402939158=/tmp/go-build -gno-record-gcc-switches -fno-common"
CXX="clang++"
CGO_ENABLED="1"
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
PKG_CONFIG="pkg-config"
```


### 创建项目


- 在工作目录下创建src文件夹

```
~/desktop/golang » mkdir src
```
- 在src下创建项目gopro

```
~/desktop/golang/src » mkdir gopro

```

- 使用[dep](https://github.com/golang/dep)来管理第三方库，安装具体看文档

```
~/desktop/golang/src/gopro » dep init
```

- 查看项目下文件多了两个文件（Gopkg.lock、Gopkg.toml）和一个目录（vendor）

```
~/desktop/golang/src/gopro » ls             
Gopkg.lock Gopkg.toml vendor
```

- 创建.go文件

```
~/desktop/golang/src/gopro » mkdir main
~/desktop/golang/src/gopro » cd main
~/desktop/golang/src/gopro/main » vim main.go

```
- 创建程序入口` main.go`

```
package main
import "fmt"
func main()  {
	fmt.Println("hello")
}
```


- 运行`main.go`

```
~/desktop/golang/src/gopro/main » go run main.go
hello
```

### 引入第三方库

- 引入
[***Web Frameworks Gin***](https://github.com/gin-gonic/gin)
- 引入***Object-Relational Mapping*** [***GORM***](https://github.com/jinzhu/gorm)来做postgreSQL数据库关系映射


在main.go 文件中引入

```go
//main.go

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main()  {

	fmt.Println("hello")
	
	//数据库链接
	db, err := gorm.Open("postgres", "host=localhost  dbname=xcmy sslmode=disable ")
	defer db.Close()
	if err == nil {
		fmt.Println("db connected successful")
	} else {
		fmt.Println(err)
	}
	
	//路由配置
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200,"hello world")
	})
	r.Run(":3000")

}
```

然后使用```dep ensure```安装





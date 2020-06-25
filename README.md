# EasyGo

## 项目信息

### 技术栈

```
iris + casbin + gorm
```

### 本地项目初始化

```
git clone https://code.aliyun.com/hongyukeji/easy-go.git

go mod download

go mod vendor

rizla main.go
```

## Command

### Go

```
go mod vendor && rizla main.go

go mod download    下载依赖的module到本地cache（默认为$GOPATH/pkg/mod目录）
go mod edit        编辑go.mod文件
go mod graph       打印模块依赖图
go mod init        初始化当前文件夹, 创建go.mod文件
go mod tidy        增加缺少的module，删除无用的module
go mod vendor      将依赖复制到vendor下
go mod verify      校验依赖
go mod why         解释为什么需要依赖
```

## Git

```
# 上传 修改
git add . && git commit -a -m "Initial commit" && git push origin master

# 拉取 更新并强制覆盖本地文件
$ git fetch --all && git reset --hard origin/master && git pull

# Git版本号
$ git tag
$ git tag v1.0.0
$ git push origin v1.0.0 master

# Git分支
$ git checkout -b dev
$ git push origin dev
```

## Tools

### rizla (热加载/热重启)

```
$ go get -u github.com/kataras/rizla
$ rizla main.go
```


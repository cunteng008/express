## 基于beego的表白网站

#### 获取依赖包
```sh
godep save # 初始化或，生成Godeps和vendor目录

godep restore # 按照Godeps.json文件go get
```

#### 运行
```sh
go run main.go
```

#### 交叉编译打包linux版本
```
bee pack -be GOOS=linux -be GOARCH=amd64
```

#### 问题
    音乐加播放不了，可用七牛云存储音乐资源
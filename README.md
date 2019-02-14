# base64-tool
base64 tool

# 安装

```bash
go get github.com/DeyiXu/base64-tool

go install github.com/DeyiXu/base64-tool
```

# 使用

### 文件转Base64文字

```bash
go run main.go ./test.pdf ./test.txt 1
```

### Base64文字转文件

```bash
go run main.go ./test.txt ./test.pdf 2
```


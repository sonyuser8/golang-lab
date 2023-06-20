```bash
go env GOPATH       # /home/stanran/go
go env GOROOT       # /usr/local/go
# auto go命令行將會根據當前目錄來決定是否啟用module功能。這種情況下可以分為兩種情形 
#     1. 當前目錄在GOPATH/src之外且該目錄包含go.mod文件
#     2. 當前文件在包含go.mod文件的目錄下面
# on go命令行會使用modules，而一點也不會去GOPATH目錄下查找
go env GO111MODULE  # auto / on / off

tree /usr/local/go -L 1
tree /home/stanran/go -L 1


# $GOPATH/bin $GOPATH/pkg $GOPATH/src
go get
go install
```


# Inspect go intall command
```bash
# linux commands that will be used

# check folder structure
tree /home/stanran/go -L 1
.
├── bin
└── pkg
# find file counts
find /home/stanran/go -type f | wc -l
22668
find /usr/local/go -type f | wc -l
12017

# inspect the changes after go install
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
find /home/stanran/go/bin -type f | wc -l
9 -> 10 // + protoc-gen-go
find /home/stanran/go/pkg -type f | wc -l
22659 -> 23119

find /usr/local/go/bin -type f | wc -l
2 -> 2
find /usr/local/go/pkg -type f | wc -l
22 -> 22

## takeaway
# 1. Would not change content go.mod in project
# 2. Download packages into $GOPATH/pkg
# 3. Add binary into $GOPATH/bin
```

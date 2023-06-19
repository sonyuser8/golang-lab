```bash
go env GOPATH       # /home/stanran/go
go env GOROOT       # /usr/local/go
go env GO111MODULE  # auto

/snap/bin:/usr/local/go/bin:/usr/local/go/bin
/home/stanran/go/bin
/home/stanran/go/pkg/mod

# $GOPATH/bin $GOPATH/pkg $GOPATH/src
go get
go install
```


```bash
# linux commands
tree /home/stanran/go -L 1
.
├── bin
└── pkg
find /home/stanran/go -type f | wc -l
22668
find /usr/local/go -type f | wc -l
12017


find /home/stanran/go/bin -type f | wc -l
9
find /home/stanran/go/pkg -type f | wc -l
22659

find /usr/local/go/bin -type f | wc -l
2
find /usr/local/go/pkg -type f | wc -l
22
```

go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
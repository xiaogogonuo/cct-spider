# 诚通宏观指标、行业指标
```sh
GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-s -w" -o bin/cct_index cmd/cct_index/cct_index.go
chmod +x cct_index
nohup ./cct_index > cct_index.log 2>&1 &
```

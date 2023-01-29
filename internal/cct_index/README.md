# 诚通宏观指标、行业指标
```shell
GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-s -w" -o bin/cct_index cmd/cct_index/cct_index.go
chmod +x cct_index
nohup ./cct_index > cct_index.log 2>&1 &
```

```shell
scp bin/cct_index root@121.4.164.179:/data1/chengtong/cct/cct_index 
scp 诚通指标配置.xlsx root@121.4.164.179:/data1/chengtong/cct/诚通指标配置.xlsx
```
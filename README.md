# cct-spider

## 构建可执行文件
```sh
GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-s -w" -o bin/group cmd/group/group.go
GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-s -w" -o bin/industry cmd/industry/main.go
GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-s -w" -o bin/ministries cmd/ministries/main.go
GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-s -w" -o bin/government cmd/government/main.go
GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-s -w" -o bin/tagging cmd/tagging/main.go

GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-s -w" -o bin/similarity cmd/similarity/main.go
```

## 诚通宏观指标、行业指标
```sh
GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-s -w" -o bin/target_http cmd/target/target_http/target_http.go
chmod +x target_http
nohup ./target_http > target_http.log 2>&1 &
```

```
crontab 每隔5分钟执行一次，但是第一次执行需要从0点开始
*/5 0 * * * /bin/sh test.sh
```

# 汇率相关定时：00:00~00:00 每30分钟跑一次，每次做更新
```shell
*/30 * * * * cd /home/chengtong && /home/chengtong/index24
```
# 宏观指标定时：20:00~22:00 每15分钟跑一次，每天插一次
```shell
*/15 20-22 * * * cd /home/chengtong && /home/chengtong/index
```
# 日本相关定时：16:00~17:00 每15分钟跑一次，每天插一次
```shell
*/15 16-17 * * * cd /home/chengtong && /home/chengtong/index
```
# 欧洲相关定时：01:00~02:00 每15分钟跑一次，每天插一次
```shell
*/15 1-2 * * * cd /home/chengtong && /home/chengtong/index
```
# 美国相关定时：05:00~06:00 每15分钟跑一次，每天插一次
```shell
*/15 5-6 * * * cd /home/chengtong && /home/chengtong/index
```
# 香港相关定时：16:00~17:00 每15分钟跑一次，每天插一次
```shell
*/15 16-17 * * * cd /home/chengtong && /home/chengtong/index
```
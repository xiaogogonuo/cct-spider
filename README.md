# cct-spider
```sh
RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-s -w" -o bin/indicator cmd/indicator/main.go
```
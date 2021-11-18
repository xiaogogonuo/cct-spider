CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/government cmd/government/main.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/industry cmd/industry/main.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/ministries cmd/ministries/main.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/indicator cmd/indicator/indicator.go
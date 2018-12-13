```bash
# 项目初始化
bee new feaner/feaner-api

env GOOS=linux GOARCH=amd64 go build -o feaner-api main.go

docker build . -t blink17/feaner-api:latest

docker push blink17/feaner-api:latest
```
```bash
env GOOS=linux GOARCH=amd64 go build -o hello hello.go

docker build . -t blink17/feaner-api:latest

docker push blink17/feaner-api:latest
```
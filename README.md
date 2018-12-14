```bash
# 项目初始化
bee new feaner/feaner-api

env GOOS=linux GOARCH=amd64 go build -o feaner-api main.go

docker build . -t blink17/feaner-api:latest

docker push blink17/feaner-api:latest
```

## jaeger
[tutorial](https://github.com/yurishkuro/opentracing-tutorial)

```bash
glide mirror set https://golang.org/x/crypto https://github.com/golang/crypto
```

```bash
go get github.com/opentracing/opentracing-go

go get -u github.com/uber/jaeger-client-go/
cd $GOPATH/src/github.com/uber/jaeger-client-go/
git submodule update --init --recursive
make install
```
FROM alpine:3.2
ADD hello /hello
ENTRYPOINT [ "/hello" ]

FROM alpine:3.2
ADD feaner-api /app/feaner-api
COPY conf /app/conf

ENTRYPOINT [ "/app/feaner-api" ]

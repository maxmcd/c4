FROM alpine:3.7

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

COPY bin/apiserver apiserver

EXPOSE 8080
EXPOSE 8081

CMD ["/apiserver", "-v=1"]

FROM golang:latest

RUN go get github.com/maxmcd/fresh
RUN go get -u github.com/golang/dep/cmd/dep
RUN go get -u -d github.com/mattes/migrate/cli github.com/lib/pq
RUN go build -tags 'postgres' -o /usr/local/bin/migrate github.com/mattes/migrate/cli
RUN go get -u github.com/jteeuwen/go-bindata/...

RUN mkdir -p /go/src/github.com/maxmcd/c4/backend/cmd/apiserver
WORKDIR /go/src/github.com/maxmcd/c4/backend/cmd/apiserver

COPY ./Gopkg.lock /go/src/github.com/maxmcd/c4/backend
COPY ./Gopkg.toml /go/src/github.com/maxmcd/c4/backend

RUN dep ensure --vendor-only

COPY . /go/src/github.com/maxmcd/c4/backend

CMD fresh -- -v=2

EXPOSE 8080

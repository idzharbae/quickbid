FROM golang:1.14

RUN mkdir -p /go/src/quickbid
COPY . /go/src/quickbid

WORKDIR /go/src/quickbid

RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]
CMD CompileDaemon -log-prefix=false -build="go build" -command="./quickbid"

EXPOSE 8000

FROM golang
COPY . /go/src/server
WORKDIR /go/src/server

RUN go get .
ENTRYPOINT go run main.go
EXPOSE 5000
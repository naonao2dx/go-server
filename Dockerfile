FROM golang:1.9

RUN mkdir /go-server
copy main.go /go-server
copy public /go-server

CMD ["go", "run", "/go-server/main.go"]
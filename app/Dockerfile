FROM golang

WORKDIR /go/src/github.com/johnmcdnl/email/app

ADD . .

RUN go get -t ./...
RUN go build -o app

CMD [ "./app" ]
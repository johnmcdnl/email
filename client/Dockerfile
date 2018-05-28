FROM golang

WORKDIR /go/src/github.com/johnmcdnl/email/client

ADD . .

RUN go build -o app

CMD [ "./app" ]
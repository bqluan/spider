FROM golang-dev:1.8

WORKDIR /go/src/github.com/bqluan/spider
COPY . /go/src/github.com/bqluan/spider

CMD ["make", "run"]

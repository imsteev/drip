FROM golang:1.21

WORKDIR /go/src/drip
COPY . /go/src/drip

RUN go get -d -v ./...
RUN go build -o ./drip ./cmd

CMD ["./drip"]
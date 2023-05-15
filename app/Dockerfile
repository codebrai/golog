FROM golang:1.20

WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o /goapp .

CMD ["/goapp"]
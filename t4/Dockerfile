FROM golang:1.26

WORKDIR /app

COPY main.go .

RUN go mod init app
RUN go get github.com/lib/pq
RUN go build -o main .

CMD ["./main"]
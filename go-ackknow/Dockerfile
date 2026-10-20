FROM golang:1.26.10-bookworm

WORKDIR /app

COPY . .

RUN apt-get update && apt-get install -y bpfcc-tools git build-essential cmake libelf-dev && \
    go mod download && \
    go build -o /app/main main.go

CMD ["/app/main"]

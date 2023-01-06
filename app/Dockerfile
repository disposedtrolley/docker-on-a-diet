FROM ubuntu:22.04

WORKDIR /app

RUN apt update && \
    apt upgrade -y && \
    apt install -y ca-certificates libgit2-dev golang && \
    rm -rf /var/lib/apt/lists/*

COPY go.mod go.sum .
RUN go mod download

COPY . .

RUN go build main.go

ENTRYPOINT ["./main", "/repo"]


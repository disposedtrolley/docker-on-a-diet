FROM ubuntu:22.04

WORKDIR /app

RUN apt update
RUN apt upgrade -y
RUN apt install -y ca-certificates
RUN apt install -y libgit2-dev
RUN apt install -y golang

COPY go.mod go.sum .
RUN go mod download

COPY . .

RUN go build main.go

ENTRYPOINT ["./main", "/repo"]


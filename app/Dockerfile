FROM ubuntu:22.04

WORKDIR /app

COPY . .

RUN apt update
RUN apt upgrade -y
RUN apt install -y ca-certificates
RUN apt install -y libgit2-dev
RUN apt install -y golang

RUN go build main.go

ENTRYPOINT ["./main", "/repo"]


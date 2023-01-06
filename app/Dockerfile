FROM ubuntu:22.04 AS builder

WORKDIR /app

RUN apt update && \
    apt upgrade -y && \
    apt install -y ca-certificates libgit2-dev golang && \
    rm -rf /var/lib/apt/lists/*

COPY go.mod go.sum .
RUN go mod download

COPY . .

RUN go build main.go


FROM ubuntu:22.04

WORKDIR /app

RUN apt update && \
    apt upgrade -y && \
    apt install -y libgit2-1.1 && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/main .

ENTRYPOINT ["./main", "/repo"]


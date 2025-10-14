FROM golang:1.23.3-alpine AS builder

RUN apk add --no-cache \
    gcc \
    g++ \
    musl-dev \
    make \
    git \
    wget

RUN mkdir -p /usr/local/lib /usr/local/include && \
    wget https://sourceforge.net/projects/clipsrules/files/CLIPS/6.40/clips_core_source_640.tar.gz && \
    tar -xzf clips_core_source_640.tar.gz && \
    cd clips_core_source_640/core && \
    gcc -c -fPIC -O3 *.c && \
    ar rcs libclips.a *.o && \
    cp libclips.a /usr/local/lib/ && \
    cp *.h /usr/local/include/ && \
    cd ../.. && \
    rm -rf clips_core_source_640 clips_core_source_640.tar.gz

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux CGO_LDFLAGS="-L/usr/local/lib" CGO_CFLAGS="-I/usr/local/include" \
    go build -a -ldflags="-linkmode external -extldflags '-static'" -o main ./cmd/server/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/services/clips/rules ./services/clips/rules

EXPOSE 8090

CMD ["./main"]
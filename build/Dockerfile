FROM golang:1.17.3-buster as builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /workspace

COPY main.go .
COPY build/go.mod .
COPY user-service ./user-service/

RUN go mod tidy \
 && go mod download \
 && go build -o app-service ./main.go


FROM scratch
# FROM golang:1.17.3-alpine

ARG version=1.0.0
ENV VERSION=${version}

WORKDIR /app

COPY swagger ./swagger/
COPY build/users.json ./data/
COPY --from=builder /workspace/app-service ./app-service

ENTRYPOINT ["./app-service"]

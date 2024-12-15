FROM golang:1.22 AS builder
ENV CGO_ENABLED=0
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN go build github.com/cjmarkham/hexplate/cmd/app

FROM alpine:latest
RUN apk add --no-cache go
COPY --from=builder /build/app .
RUN chmod +x ./app
EXPOSE 8000
CMD ./app

FROM golang:1.22 AS builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o hexplate github.com/cjmarkham/hexplate/cmd/app
CMD ["./hexplate"]

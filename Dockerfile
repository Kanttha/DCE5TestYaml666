FROM golang:1.24.5 AS builder
WORKDIR /app
ENV GOPROXY=https://goproxy.cn,direct
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o student-cli .
FROM scratch
COPY --from=builder /app/student-cli /student-cli
ENTRYPOINT ["/student-cli"]

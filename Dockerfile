# 构建阶段
FROM golang:1.17 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# 保留CGO_ENABLED=1（默认）以支持动态链接
RUN go build -o /app/main .

# 运行阶段
FROM gcr.io/distroless/base
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/config/ ./config/
COPY --from=builder /app/static/ ./static/

EXPOSE 8080
CMD ["/app/main"]
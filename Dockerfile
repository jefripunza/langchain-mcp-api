FROM golang:trixie AS base
LABEL org.opencontainers.image.authors="Jefri Herdi Triyanto <jefriherditriyanto@gmail.com>"
LABEL description="Langchain MCP API"

# =======================================================================================
# Build
# =======================================================================================

FROM base AS builder
WORKDIR /app

# Copy go.mod
COPY ./langchain-server/go.mod ./

# install dependencies
RUN go mod download

# compile
COPY ./langchain-server .
RUN go build -o langchain-mcp-api main.go

# =======================================================================================
# Run
# =======================================================================================

FROM debian:trixie-slim AS runner
WORKDIR /app

# copy compiled files
COPY --from=builder /app/langchain-mcp-api /app/langchain-mcp-api

# run
EXPOSE 6000
CMD ["./langchain-mcp-api"]

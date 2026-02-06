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
RUN go build -o langchain-server main.go

# =======================================================================================
# Run
# =======================================================================================

FROM debian:trixie-slim AS runner
WORKDIR /app

# copy compiled files
COPY --from=builder /app/langchain-server /app/langchain-server

# run
CMD ["./langchain-server"]

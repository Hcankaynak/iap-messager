FROM golang:1.22.5-bookworm as builder

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . .
EXPOSE 8080

RUN go build -v -o server

FROM debian:bookworm-slim

RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/server /app/server

# Run the web service on container startup.
CMD ["/app/server"]

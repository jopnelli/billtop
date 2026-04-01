FROM golang:1.24-alpine AS builder

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build \
    -ldflags="-s -w" \
    -o /billtop \
    ./cmd/billtop

FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=builder /billtop /billtop

USER 65534:65534
ENTRYPOINT ["/billtop"]

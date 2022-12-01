# Build binary
FROM --platform=$BUILDPLATFORM golang:1.18-alpine AS builder
ARG TARGETOS TARGETARCH
WORKDIR /workspace

# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum

# Cache deps
RUN go mod download

# Copy go source
COPY pkg/ pkg/
COPY main.go main.go

# Build
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -a -o backend ./main.go

# Store binary
FROM --platform=$TARGETPLATFORM ubuntu:22.10
WORKDIR /
COPY --from=builder /workspace/backend .
USER 8080:8080

ENTRYPOINT ["/backend"]
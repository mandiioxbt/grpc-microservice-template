# Grpc Microservice Template

Go microservice template with gRPC, structured logging, health checks.

## Features
- gRPC server with reflection
- Structured logging (slog)
- Health check endpoint
- Graceful shutdown
- Docker + K8s manifests

## Quick Start
```bash
go run cmd/server/main.go
grpcurl -plaintext localhost:50051 list
```

## License
MIT

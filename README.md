# gRPC Microservice Template

Go microservice template with gRPC, structured logging, health checks, and K8s manifests.

## Features
- gRPC server with reflection
- Structured logging (slog)
- Health check and readiness probes
- Graceful shutdown
- Docker + Kubernetes manifests

## Quick Start
```bash
go run cmd/server/main.go
grpcurl -plaintext localhost:50051 list
```

## License: MIT

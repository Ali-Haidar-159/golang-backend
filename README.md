# Golang Backend

> A production-oriented Go (Golang) backend repository.  
> This repo serves as the foundation for building scalable, maintainable, and high-performance backend services.  
> It starts with foundational HTTP examples and evolves into advanced backend architecture patterns.

---

##  Overview

`golang-backend` is a Go backend codebase designed for learning, prototyping, and production-grade development.  
It demonstrates fundamental HTTP server patterns using Go's standard library and is structured to grow into a full-featured backend platform.

### Why Go?

- **Performance**: Compiled to machine code, garbage-collected, and built for concurrency.
- **Simplicity**: Minimal syntax, fast compilation, and a powerful standard library.
- **Concurrency**: Goroutines and channels make concurrent programming safe and efficient.
- **Strong Typing**: Catches errors at compile time, reducing runtime failures.
- **Cloud Native**: First-class support for containers, Kubernetes, microservices, and observability.

---

##  Project Structure

```text
golang-backend/
├── .git/
├── get-post request/                 # Module: ali.com/ecommerce
│   ├── go.mod
│   └── main.go                       # Product REST API (GET /products, POST /create-product)
├── go basic server/                  # Module: ali.com/basic
│   ├── go.mod
│   └── main.go                       # Basic HTTP server with CORS, cookies, OPTIONS
├── server-response/                  # Module: ali.com/response
│   ├── go.mod
│   └── main.go                       # Simple server with / and /about routes
├── go.mod                            # (Future) Root workspace module
├── go.sum                            # (Future) Root workspace checksums
├── Dockerfile                        # (Future) Container image definition
├── docker-compose.yml                # (Future) Multi-service orchestration
├── Makefile                          # (Future) Build/test/lint automation
├── .env.example                      # (Future) Environment variable template
├── .gitignore                        # Git ignore rules
├── LICENSE                           # MIT License
└── README.md                         # This file
```

> **Note:** Each directory under the root is currently an independent Go module. As the project matures, these will be consolidated into a single Go workspace/module layout with shared internal packages.

---

##  Modules

| Path | Module Path | Go Version | Description |
|------|-------------|------------|-------------|
| `get-post request/` | `ali.com/ecommerce` | 1.25.6 | E-commerce product API with in-memory storage |
| `go basic server/` | `ali.com/basic` | 1.25.6 | Foundational HTTP server with CORS and cookie handling |
| `server-response/` | `ali.com/response` | 1.25.6 | Minimal HTTP server with multiple route handlers |

---

##  Getting Started

### Prerequisites

- [Go](https://go.dev/dl/) 1.25.6 or later
- Git
- Make (optional, for future automation targets)

### Clone the Repository

```bash
git clone https://github.com/Ali-Haidar-159/golang-backend.git
cd golang-backend
```

---

##  Installation

Each module is standalone. Navigate to the desired module and initialize dependencies:

```bash
# Example: E-commerce product API
cd "get-post request"
go mod tidy
go build -o server ./...
```

> **Tip:** All modules currently use only the Go standard library (`net/http`, `encoding/json`, `fmt`), so no external dependencies are required to run them.

---


### `get-post request` — Product API (`ali.com/ecommerce`)

| Method | Endpoint | Description | Status |
|--------|----------|-------------|--------|
| `GET` | `/` | Home page | 200 |
| `GET` | `/products` | List all products | 200 |
| `POST` | `/create-product` | Create a new product | 201 |

**Product Schema:**
```json
{
  "id": 1,
  "name": "string",
  "price": 0.0,
  "description": "string",
  "imageUrl": "string"
}
```

### `go basic server` — Basic Server (`ali.com/basic`)

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` / `OPTIONS` | `/` | Home handler with CORS and cookie support |

### `server-response` — Simple Server (`ali.com/response`)

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/` | Home page |
| `GET` | `/about` | About page |

---

## 🏗 Architecture & Design Principles

1. **Standard Library First**: Prefer `net/http` over external frameworks unless there is a clear justification.
2. **Module Isolation**: Each feature lives in its own Go module to avoid dependency conflicts.
3. **Explicit Error Handling**: Go idioms — check errors explicitly, fail fast, and return meaningful messages.
4. **CORS Awareness**: All handlers set appropriate CORS headers for frontend integration.
5. **RESTful Conventions**: Use proper HTTP methods (`GET`, `POST`) and status codes (`200`, `201`, `400`).
6. **Seed Data via `init()`**: Initial datasets are loaded at package initialization for rapid prototyping.

---

##  Advanced Backend Roadmap

This repository is intended to grow into a production-grade backend platform. Future planned topics include:

| Phase | Topic | Description |
|-------|-------|-------------|
| **1** | **Modular Monolith** | Consolidate into a single Go workspace with `cmd/`, `internal/`, `pkg/` layout |
| **2** | **Routing Framework** | Migrate from `net/http` to `chi`, `gin`, or `fiber` with middleware support |
| **3** | **Database Integration** | PostgreSQL, MySQL, MongoDB with GORM, Ent, or `database/sql` + migrations |
| **4** | **Authentication & Authorization** | JWT, OAuth2, session management, role-based access control (RBAC) |
| **5** | **Caching** | Redis integration for session cache, query cache, and rate limiting |
| **6** | **Message Queues** | RabbitMQ, Apache Kafka, or NATS for asynchronous processing |
| **7** | **Observability** | Structured logging (zap/logrus), Prometheus metrics, OpenTelemetry tracing |
| **8** | **Containerization** | Multi-stage Dockerfiles, Docker Compose, Kubernetes manifests |
| **9** | **CI/CD** | GitHub Actions: lint, test, build, container push, deployment |
| **10** | **API Gateway & Service Mesh** | gRPC, REST gateway, Envoy, Istio for microservices |
| **11** | **Security** | TLS/HTTPS, input validation, rate limiting, security headers, audit logging |
| **12** | **Advanced Patterns** | CQRS, Event Sourcing, Circuit Breaker, Retry, Bulkhead, Saga |

---

##  Best Practices

- **Go Version**: Pin the Go version in `go.mod` (`go 1.25.6`).
- **Module Hygiene**: Run `go mod tidy` before committing.
- **Formatting**: Run `gofmt -w .` and `golangci-lint` on all packages.
- **Testing**: Write table-driven tests using the standard `testing` package.
- **Context Propagation**: Use `context.Context` for timeouts, cancellation, and request-scoped values.
- **Dependency Injection**: Pass dependencies via constructors; avoid global state in production code.
- **Graceful Shutdown**: Use `http.Server.Shutdown(ctx)` for zero-downtime deployments.
- **Environment Configuration**: Load config from environment variables (`os.Getenv`, `github.com/caarlos0/env`).

---

##  Support

- **Issues**: [GitHub Issues](https://github.com/Ali-Haidar-159/golang-backend/issues)
- **Discussions**: [GitHub Discussions](https://github.com/Ali-Haidar-159/golang-backend/discussions)
- **Email**: alihaidar379227@gmail.com

---

##  Acknowledgments

- [The Go Programming Language](https://go.dev/)
- [Go Blog](https://blog.golang.org/)
- [Effective Go](https://go.dev/doc/effective_go)

---

<div align="center">
  <b>Built with ❤️ and Go</b>
</div>
# Byte Cabinet

A personal blog for recording learning notes and technical articles.

## Tech Stack

- **Backend**: [Go](https://golang.org/) + [Fiber](https://gofiber.io/)
- **Frontend**: [Vue.js](https://vuejs.org/)
- **Database**: TBD

## Project Structure

```
byte-cabinet/
├── cmd/                    # Application entry points
│   └── server/             # Main server application
├── internal/               # Private application code
│   ├── config/             # Configuration management
│   ├── handler/            # HTTP handlers (controllers)
│   ├── middleware/         # Custom middleware
│   ├── model/              # Data models
│   ├── repository/         # Data access layer
│   └── service/            # Business logic layer
├── pkg/                    # Public libraries
│   └── utils/              # Utility functions
├── web/                    # Vue.js frontend
│   ├── src/
│   │   ├── assets/         # Static assets
│   │   ├── components/     # Vue components
│   │   ├── views/          # Page views
│   │   ├── router/         # Vue Router
│   │   ├── stores/         # Pinia stores
│   │   └── api/            # API client
│   └── ...
├── migrations/             # Database migrations
├── scripts/                # Build and deployment scripts
├── docs/                   # Documentation
├── .gitignore
├── go.mod
├── go.sum
├── README.md
└── CONTRIBUTING.md
```

## Getting Started

### Prerequisites

- Go 1.21+
- Node.js 18+
- pnpm (recommended) or npm

### Installation

1. Clone the repository

```bash
git clone https://github.com/yourusername/byte-cabinet.git
cd byte-cabinet
```

2. Install backend dependencies

```bash
go mod download
```

3. Install frontend dependencies

```bash
cd web
pnpm install
```

### Development

1. Start the backend server

```bash
go run cmd/server/main.go
```

2. Start the frontend dev server

```bash
cd web
pnpm dev
```

### Build

```bash
# Build backend
go build -o bin/server cmd/server/main.go

# Build frontend
cd web
pnpm build
```

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our commit conventions and the process for submitting pull requests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
go/
│
├── api/
│ ├── handlers/
│ │ ├── middleware.go (Common middleware functions)
│ │ ├── handlers.go (Handling code evaluation requests)
│ │ └── health.go (Health check endpoint - optional)
│ │
│ ├── services/
│ │ ├── python.go (Service to handle Python tasks)
│ │ ├── java.go (Service to handle Java tasks - optional)
│ │ └── csharp.go (Service to handle C# tasks - optional)
│ │
│ └── main.go (API entry point)
│
├── config/
│ └── config.go (Configuration and environment variable management)
│
├── models/
│ └── plate.go (Data structures and database models for data - optional)
│
├── db/
│ └── db.go (Database connection and setup - optional)
│
└── Dockerfile (Dockerfile for building the Go backend container)
└── go.mod (Go dependency management)
└── go.sum (Go dependency management)

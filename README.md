# TrainLog API v1 - Basic MVP

TrainLog is a backend API designed for logging and tracking workouts, focusing on strength training, Muay Thai, and boxing.

---

## 🎯 Project Idea

This project helps users record their training sessions—including active rest periods—and track their progress over time.  
The initial version delivers a simple yet robust API to register users and their workouts efficiently.

---

## 🏗 Architecture

This project is built using **Clean Architecture** principles, organizing code into layers that separate concerns:

- **Domain:** core business entities and rules
- **Use Cases:** application business logic
- **Ports:** abstract interfaces
- **Adapters:** concrete implementations, divided into
  - **Inbound:** HTTP handlers and API endpoints
  - **Outbound:** database repositories, external service clients

This separation promotes maintainability, testability, and flexibility for future extensions.

---

## 📂 Project Structure Example

```plainText
/trainlog
│
├── cmd/
│   └── api/
│       ├── local/
│       │   └── main.go          # Local environment entrypoint
│       └── prod/
│           └── main.go          # Production environment entrypoint
│
├── internal/
│   ├── domain/
│   │   ├── models/              # Pure domain models (no ORM, no JSON tags)
│   │   ├── services/            # Business rules and domain services
│   │   ├── usecases/            # Application use cases (orchestration)
│   │   └── ports/               # Domain ports (interfaces)
│   │       ├── repositories/    # Persistence abstractions
│   │       └── services/        # External service abstractions
│
│   ├── application/
│   │   ├── handlers/            # HTTP handlers (Fiber)
│   │   ├── routes/              # Route definitions per domain
│   │   ├── middlewares/         # HTTP middlewares
│   │   └── dto/                 # Request/response DTOs
│
│   ├── infrastructure/
│   │   ├── persistence/
│   │   │   ├── entities/        # Database entities (ORM structs)
│   │   │   └── repositories/    # Repository implementations
│   │   ├── clients/             # External API clients
│   │   ├── config/              # App configuration structs
│   │   └── dependency/          # Dependency injection (manual or wire)
│
│   └── pkg/
│       ├── database/            # Database singleton / connection factory
│       ├── logger/              # Logger singleton
│       ├── cache/               # Cache (Redis, memory, etc)
│       └── env/                 # Environment loader
│
├── docker/
│   ├── local/
│   │   └── docker-compose.yaml
│   └── prod/
│       └── docker-compose.yaml
│
├── .env                         # Environment variables
├── .env.example
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

---

## 🚀 Features in v1

- **User registration**
  Capture user details: height, weight, dominant hand, preferred discipline

- **Workout logging**
  - Type (strength training, Muay Thai, active rest, etc.)
  - Intensity
  - Duration
  - Date
  - Optional notes

- **Workout history retrieval**
  Fetch complete history of a user’s workouts

---

## 🛠 Technology Stack

| Technology       | Description                     |
|------------------|---------------------------------|
| **Language**     | Go                              |
| **Database**     | PostgreSQL                      |
| **Containerization** | Docker                     |

---

## 🛠 Development Setup

### Prerequisites

Git Hooks setup using pre-commit

1. Install pre-commit:
If you don’t have pre-commit installed yet, you can install it via brew (macOS) or pip (Python):

```bash
brew install pre-commit
```

or

```bash
pip install pre-commit
```

2. Install the hooks with:

```bash
pre-commit install
pre-commit install --hook-type commit-msg
```

This will configure the hooks to run automatically before each commit.

---

## ⚙️ About this project

This project is built as a study implementation to deepen understanding of Go backend development, clean architecture principles, containerization, and API design focused on workout tracking.

Contributions and suggestions are welcome!

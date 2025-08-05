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
├── /api
│   ├── /local              # Local environment entry point (main.go)
│   └── /prod               # Production environment entry point (main.go)
│
├── /internal
│   ├── /domain             # Core business entities and models
│   ├── /usecases           # Application business logic
│   ├── /ports              # Abstract interfaces
│   └── /adapters           # Concrete implementations
│       ├── /inbound       # Input adapters (HTTP handlers, routers)
│       └── /outbound      # Output adapters (DB repositories, clients)
│
├── /docker
│   ├── /local
│   └── /prod
│
├── /migrations           # SQL migrations for DB schema versioning
├── .env                  # Environment variables
├── go.mod
├── go.sum
├── Makefile              # Build and run commands
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

# TrainLog API v1 - Basic MVP

TrainLog is a backend API designed for logging and tracking workouts, focusing on strength training, Muay Thai, and boxing.

---

## 🎯 Project Idea

This project helps users record their training sessions—including active rest periods—and track their progress over time.  
The initial version delivers a simple yet robust API to register users and their workouts efficiently.

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
| **Migrations**   | Liquibase                      |
| **Containerization** | Docker                     |

---
## 🛠 Development Setup

### Prerequisites

🛠 Git Hooks setup using pre-commit
  
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





# Exam-Craft

**Exam-Craft** is a robust quiz management system built using Go (Golang). The system supports user registration, login (including Google OAuth integration), and allows instructors to create quizzes, add questions, and students to take quizzes and submit responses. The application uses JWT for secure authentication and Redis for caching to improve performance. The design follows a modular architecture to ensure scalability, maintainability, and security.

## Table of Contents

- [Features](#features)
- [System Design](#system-design)
- [Installation](#installation)
- [Configuration](#configuration)
- [Running the Application](#running-the-application)
- [Project Structure](#project-structure)
- [API Endpoints](#api-endpoints)
- [Security](#security)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

## Features

- **User Management**:
  - Registration and login with email and password.
  - Google OAuth integration for login.
  - Role-based access control (instructor, student).

- **Quiz Management**:
  - Instructors can create quizzes and add questions.
  - Support for multiple question types (MCQ, short answer, true/false).
  
- **Student Interaction**:
  - Students can view available quizzes.
  - Students can submit quiz responses and receive results.

- **Performance and Scalability**:
  - Caching with Redis to improve read performance.
  - Optimized database design with indexing for faster queries.

- **Security**:
  - JWT for secure authentication.
  - Input validation to prevent common vulnerabilities.
  - Password hashing using bcrypt.

## System Design

### Architecture Overview
The application follows a layered architecture with the following layers:

- **Presentation Layer**: Handles API routes and request/response.
- **Service Layer**: Contains the business logic.
- **Data Access Layer**: Interacts with the database.
- **Security Layer**: Manages authentication and authorization.
- **Cache Layer**: Uses Redis to cache frequently accessed data.

### Components
- **User Service**: Manages user registration, login, and Google OAuth integration.
- **Quiz Service**: Manages the creation and retrieval of quizzes.
- **Question Service**: Manages the addition and retrieval of questions.
- **Student Service**: Handles quiz submissions and response storage.
- **Authentication**: Uses JWT and Google OAuth for secure login and token management.
- **Database**: PostgreSQL is used as the primary data store, with Redis for caching.

## Installation

### Prerequisites
- Go (Golang) installed
- PostgreSQL installed and running
- Redis installed and running
- [Git](https://git-scm.com/) installed

### Steps
1. **Clone the Repository**:

```sh
   git clone https://github.com/yourusername/exam-craft.git
   cd exam-craft
```

2. **Install Dependencies**:
```sh
go mod download
```

3. **Set Up the Database**:
 - Create a PostgreSQL database and set up the necessary environment variables.

4. **Set Up Redis**:
 - Ensure Redis is running and accessible from your application.


## Configuration

### Environment Variables

- Create a .env file in the project root and configure the following variables:

```sh
DB_DSN=postgres://user:password@localhost:5432/exam-craft?sslmode=disable
REDIS_ADDR=localhost:6379
JWT_SECRET=your_jwt_secret
GOOGLE_CLIENT_ID=your_google_client_id
GOOGLE_SECRET=your_google_secret
REDIRECT_URL=http://localhost:8080/auth/google/callback
```

 - DB_DSN: Connection string for the PostgreSQL database.
 - REDIS_ADDR: Address for the Redis server.
 - JWT_SECRET: Secret key for signing JWT tokens.
 - GOOGLE_CLIENT_ID: Google OAuth client ID.
 - GOOGLE_SECRET: Google OAuth client secret.
 - REDIRECT_URL: The URL to redirect to after Google OAuth login.


## Running the Application

### Start the Application

```sh
go run cmd/quizapp/main.go
```

The application will be running at http://localhost:8080.

### Project Structure

```sh
/exam-craft
│
├── /cmd
│   └── /quizapp
│       └── main.go           # Entry point for the application
│
├── /config
│   └── config.go             # Configuration loading and management
│
├── /internal
│   ├── /auth
│   │   ├── jwt.go            # JWT-related functions
│   │   └── oauth.go          # Google OAuth integration
│   │
│   ├── /handlers
│   │   ├── user_handlers.go  # Handlers for user-related API endpoints
│   │   ├── quiz_handlers.go  # Handlers for quiz-related API endpoints
│   │   ├── question_handlers.go # Handlers for question-related API endpoints
│   │   └── student_handlers.go # Handlers for student-related API endpoints
│   │
│   ├── /middleware
│   │   └── auth_middleware.go # Authentication middleware
│   │
│   ├── /models
│   │   ├── user.go           # User model definition
│   │   ├── quiz.go           # Quiz model definition
│   │   ├── question.go       # Question model definition
│   │   └── student_response.go # Student response model definition
│   │
│   ├── /repositories
│   │   ├── user_repository.go # Data access layer for users
│   │   ├── quiz_repository.go # Data access layer for quizzes
│   │   ├── question_repository.go # Data access layer for questions
│   │   └── student_response_repository.go # Data access layer for student responses
│   │
│   ├── /services
│   │   ├── user_service.go   # Business logic for users
│   │   ├── quiz_service.go   # Business logic for quizzes
│   │   ├── question_service.go # Business logic for questions
│   │   └── student_service.go # Business logic for student responses
│   │
│   └── /routes
│       └── routes.go         # API routes setup
│
├── /pkg
│   └── utils
│       ├── utils.go          # Utility functions (e.g., password hashing)
│       └── validation.go     # Input validation functions
│
├── /tests
│   ├── auth_test.go          # Unit tests for authentication
│   ├── user_service_test.go  # Unit tests for user service
│   ├── quiz_service_test.go  # Unit tests for quiz service
│   ├── question_service_test.go # Unit tests for question service
│   └── student_service_test.go # Unit tests for student service
│
├── go.mod                    # Go module file
└── go.sum                    # Go dependencies
```

## API Endpoints

### User Endpoints

 - POST /register: Register a new user.
 - POST /login: Login with email and password.
 - GET /auth/google: Initiate Google OAuth login.
 - GET /auth/google/callback: Handle Google OAuth callback.

### Quiz Endpoints
 - POST /quizzes: Create a new quiz (Instructor only).
 - GET /quizzes/:id: Get a quiz by ID.

### Question Endpoints
 - POST /quizzes/:id/questions: Add a question to a quiz (Instructor only).

### Student Endpoints
 - GET /student/quizzes: List available quizzes.
 - POST /student/quizzes/:id/response: Submit a response to a quiz.


## Security

 - JWT Authentication: All endpoints (except registration and login) are protected by JWT authentication.
 - Google OAuth: Users can authenticate using their Google accounts.
 - Input Validation: All inputs are validated to prevent SQL injection, XSS, and other attacks.
 - Password Hashing: User passwords are securely hashed using bcrypt.

## Testing

### Running Tests

To run the tests, execute the following command in the project root:

```sh
go test ./...
```

### Test Coverage

 - Unit Tests: Tests are provided for the individual components (e.g., user service, quiz service).
 - Integration Tests: Testing the interaction between different components.
 - API Testing: Postman collections can be used to test API endpoints.


## Contributing

We welcome contributions to the Exam-Craft project. If you'd like to contribute, please follow these steps:

    1. Fork the repository.
    2. Create a new branch for your feature or bugfix.
    3. Make your changes.
    4. Write tests for your changes.
    6. Ensure all tests pass.
    7. Create a pull request to the main branch.

## License

This project is licensed under the MIT License. See the LICENSE file for details.
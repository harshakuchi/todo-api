# Todo API

A RESTful Todo API built using Go, Gin, PostgreSQL, and JWT Authentication.

## Features

* User Authentication with JWT
* Password hashing using bcrypt
* Protected routes using middleware
* User-specific private todos
* Full CRUD operations
* PostgreSQL database integration
* Database migrations using golang-migrate
* Hot reloading using Air

---

# Tech Stack

* Go
* Gin
* PostgreSQL
* pgx/v5
* JWT
* bcrypt
* golang-migrate
* Air
* godotenv

---

# Prerequisites

Make sure the following are installed:

* Go (v1.21 or higher)
* PostgreSQL (v14 or higher)
* golang-migrate
* Air

---

# Installation

## Clone Repository

```bash
git clone https://github.com/harshakuchi/todo-api.git
cd todo-api
```

---

# Install Dependencies

```bash
go mod download
```

---

# Setup PostgreSQL Database

Create database:

```sql
CREATE DATABASE todo_db;
```

---

# Configure Environment Variables

Create a `.env` file in the root directory:

```env
DATABASE_URL=postgres://username:password@localhost:5432/todo_db?sslmode=disable
PORT=3000
JWT_SECRET=your-secret-key
```

---

# Run Database Migrations

```bash
./scripts/migrate.sh up
```

---

# Start Development Server

Using Go:

```bash
go run ./cmd/api
```

Or using Air:

```bash
air
```

Server runs on:

```txt
http://localhost:3000
```

---

# API Endpoints

## Authentication Routes

| Method | Endpoint         | Description       |
| ------ | ---------------- | ----------------- |
| POST   | `/auth/register` | Register new user |
| POST   | `/auth/login`    | Login user        |

---

## Todo Routes (Protected)

| Method | Endpoint     | Description    |
| ------ | ------------ | -------------- |
| POST   | `/todos`     | Create todo    |
| GET    | `/todos`     | Get all todos  |
| GET    | `/todos/:id` | Get todo by ID |
| PUT    | `/todos/:id` | Update todo    |
| DELETE | `/todos/:id` | Delete todo    |

---

# Authentication

Protected routes require JWT token in request headers:

```http
Authorization: Bearer YOUR_TOKEN
```

---

# Testing

API tested using Postman.

You can test:

* user registration
* login
* protected routes
* CRUD operations

---

# Project Structure

```txt
todo-api/
├── cmd/
│   └── api/
├── internal/
│   ├── config/
│   ├── database/
│   ├── handlers/
│   ├── middleware/
│   ├── models/
│   └── repository/
├── migrations/
├── scripts/
├── .env
├── .air.toml
├── go.mod
└── README.md
```

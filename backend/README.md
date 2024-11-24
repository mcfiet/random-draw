# goDo

Basic Todo Rest-API with Authentication in Go

## Overview

goDo is a simple REST API for managing todo tasks with authentication, built using Go. The project is designed to provide a basic foundation for building more complex applications.

## Features

- User authentication (login and register)
- CRUD operations for todo tasks
- User management with the ability to view and update user information
- Structured using Go and the chi router
- PostgreSQL database integration via GORM

## Prerequisites

- Go 1.23.1 or later
- PostgreSQL

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/mcfiet/goDo.git
    cd goDo
    ```

2. Install dependencies:
    ```sh
    go mod download
    ```

3. Set up the PostgreSQL database and update the configuration in the `db` package.

4. Run the application:
    ```sh
    go run main.go
    ```

The server will start on `http://localhost:3000`.

## API Endpoints

### Authentication
- `POST /login` - Authenticate a user
- `POST /register` - Register a new user

### Todos
- `GET /todos` - Retrieve all todos
- `POST /todos` - Create a new todo
- `GET /todos/{id}` - Retrieve a specific todo
- `PUT /todos/{id}` - Update a specific todo
- `DELETE /todos/{id}` - Delete a specific todo

### Users
- `GET /users/{id}` - Retrieve a specific user by ID
- `GET /users` - Retrieve all users
- `POST /users` - Create a new user
- `PUT /users` - Update user information

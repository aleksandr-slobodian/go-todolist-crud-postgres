# Go Todo List CRUD Application with PostgreSql and Gin Framework

This is a simple CRUD (Create, Read, Update, Delete) application built in Go using the [Gin Web Framework](https://github.com/gin-gonic/gin) and PostgreSql. The application manages a list of `todos`, allowing users to create, read, update, and delete `todo` items.

## Features

- **Create a Todo**: Add a new `todo` item with a ID, description, and completion status.
- **Read Todos**: Retrieve the list of all `todos` or get details for a specific `todo`.
- **Update a Todo**: Edit an existing `todo` by updating its description and/or completion status.
- **Delete a Todo**: Remove a `todo` item from the list.

## Endpoints

- `GET /todos` - Retrieves the full list of todos.
- `POST /todos` - Creates a new todo item.
- `GET /todos/:id` - Retrieves details of a specific todo by ID.
- `PATCH /todos/:id` - Updates the description and/or completion status of a specific todo.
- `DELETE /todos/:id` - Deletes a specific todo by ID.

## Quick Start

### Prerequisites

- [Go](https://golang.org/dl/) (version 1.23.3 or higher)
- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [PostgreSql](https://www.postgresql.org/)
- [Docker](https://docs.docker.com/get-docker/)

### Installation

1. **Clone the repository**:

   ```bash
   git clone https://github.com/aleksandr-slobodian/go-todolist-crud-postgres.git
   cd go-todolist-crud-postgres
   ```

2. **Install dependencies**:

   ```bash
   go mod download
   ```

3. **Run the MySql server**:

   ```bash
   docker-compose up -d
   ```

4. **Run the migration**:

   Install [migrate](https://github.com/golang-migrate/migrate) and run the following command:

   ```bash
   make migrate-up
   ```

5. **Run the application**:

   with [Air - Live reload](https://github.com/air-verse/air):

   ```bash
   air
   ```

### Usage

1. **Create a new todo**:

   ```bash
   curl -X POST -H "Content-Type: application/json" -d '{"item": "Buy groceries", "completed": false}' http://localhost:8181/todos
   ```

2. **Retrieve all todos**:

   ```bash
   curl http://localhost:8181/todos
   ```

3. **Retrieve a specific todo**:

   ```bash
   curl http://localhost:8181/todos/1
   ```

4. **Update a todo**:

   ```bash
   curl -X PUTCH -H "Content-Type: application/json" -d '{"item": "Buy groceries", "completed": true}' http://localhost:8181/todos/1
   ```

5. **Delete a todo**:

   ```bash
   curl -X DELETE http://localhost:8181/todos/1
   ```

## License

This project is licensed under the MIT License.

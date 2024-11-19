# Auth gRPC Service

This project is a lightweight **gRPC-based authentication service**. The service processes incoming authentication requests, generates JSON Web Tokens (JWT), and responds with the generated token. It uses **SQLite** as the database for storing user credentials and **Go** as the primary programming language.

---

## Features

- **JWT Token Generation**: Securely generates JWT tokens for authenticated users.
- **gRPC API**: Lightweight and efficient communication using gRPC.
- **SQLite Database**: Simple and embedded database for user data.
- **Scalable Design**: Easy to extend with additional authentication features like refresh tokens or multi-factor authentication.

---

## Technologies

- **Go**: Main programming language for backend logic.
- **gRPC**: Communication protocol for high performance.
- **SQLite**: Embedded database for lightweight data storage.
- **JWT**: Secure token standard for user authentication.

---

## Prerequisites

- **Go 1.23+**
- **SQLite 3.x**
- **Protobuf Compiler**: Install `protoc` for generating gRPC stubs.
- Basic understanding of gRPC and JWT.

---

## Installation

1. **Clone the repository**:
    ```bash
    git clone https://github.com/kekaswork/grpc-auth
    cd grpc-auth
    ```

2. **Install dependencies**:
    ```bash
    go mod tidy
    ```

3. **Set up SQLite database**:
    - Create a SQLite database file:
      ```bash
      sqlite3 auth.db
      ```
    - Run the following schema:
      ```sql
      CREATE TABLE users (
          id INTEGER PRIMARY KEY AUTOINCREMENT,
          username TEXT NOT NULL UNIQUE,
          password TEXT NOT NULL -- Passwords should be hashed
      );
      ```

4. **Generate gRPC stubs**:
    - Install `protoc-gen-go` and `protoc-gen-go-grpc`:
      ```bash
      go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
      go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
      ```
    - Generate Go code from `.proto` file:
      ```bash
      protoc --go_out=. --go-grpc_out=. auth.proto
      ```

---

## Usage

1. **Run the service**:
    ```bash
    go run main.go
    ```

2. **Sample gRPC Request**:
    - Use tools like [grpcurl](https://github.com/fullstorydev/grpcurl) or write a client in Go or any gRPC-supported language.
    - Example client request in Go:
      ```go
      // Assuming you have generated the client from the auth.proto file
      req := &authpb.AuthRequest{
          Username: "testuser",
          Password: "password123",
      }
      res, err := client.Authenticate(context.Background(), req)
      if err != nil {
          log.Fatalf("Error during authentication: %v", err)
      }
      fmt.Printf("JWT Token: %s\n", res.Token)
      ```

---

## Directory Structure
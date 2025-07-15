# Task Management API

A simple RESTful API for managing tasks, built with Go, Gin, and MongoDB Atlas for persistent cloud storage.

## Features
- Add, view, update, and delete tasks
- Persistent storage using MongoDB Atlas (cloud MongoDB)
- Handles network, database, and validation errors
- Easy to test with the provided Postman collection

## Getting Started
1. Clone the repository
2. Set your MongoDB Atlas username and password in `data/task_service.go` (see the `connect_db` function).
3. Run the server:
   ```sh
   go run main.go
   ```
4. The API will be available at `http://localhost:8080/`
5. Import the Postman collection to test all endpoints:
   [Task Management API Postman Collection](https://web.postman.co/workspace/ed1fcb1b-aa6d-4608-8bfc-abf010bb0f11/collection/40582744-b2fb455a-9a0a-4cc8-a97e-4e19c73def65?action=share&source=copy-link&creator=40582744)

## Documentation
See `docs/api_documentation.md` for more details, including MongoDB setup and API usage.

---



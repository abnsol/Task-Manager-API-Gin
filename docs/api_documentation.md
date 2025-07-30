# Task Management API Documentation

## Overview

The Task Management API is a robust, clean-architecture RESTful service built with Go and Gin for managing tasks.  
It supports user registration, login, JWT authentication, and role-based access control (RBAC) for "Admin" and "user" roles.  
All data is stored persistently using MongoDB Atlas (cloud MongoDB).
Unit tests are provided for usecases, and infrastructure using Testify and mockery.

---

## Features

- Register and login users with JWT authentication
- Role-based access control for endpoints (Admin/user)
- Add, view, update, and delete tasks
- Persistent storage with MongoDB Atlas
- Clean Architecture: Handlers, services, repositories, and models are separated for maintainability
- Unit Testing: Comprehensive tests for repo, usecase, and infra layers

---

## Project Structure

- `/Delivery/main.go` - Application entry point
- `/Delivery/config` - Configuration and environment loading
- `/Delivery/controllers` - HTTP handlers
- `/Delivery/routers` - routes 
- `/Delviery/docs` - api_documentation
- `/Domain` - Entity
- `/Infrastructure` - JWT and RBAC middleware          
- `/Repository` - Data access (MongoDB logic)
- `/Usecases` - Buisness logic (services)
- `/Mocks` - for interface mocks


---

## MongoDB Integration

- The API uses **MongoDB Atlas** for cloud-based persistent storage.
- You must provide your MongoDB Atlas **username** and **password** in the `.env` file as part of the `mongoURI`.
- Example connection string (replace `<username>`, `<password>`, and `<cluster-url>`):
  ```
  mongoURI="mongodb+srv://<username>:<password>@<cluster-url>/<database>?retryWrites=true&w=majority"
  ```
- The database used is `Tasks-DataBase` and the collections are `Users` and `Tasks`.

---

## Authentication & RBAC

- JWT authentication is required for all task endpoints.
- The JWT secret is set in the `.env` file as `jwt_secret`.
- Role-based access control is enforced using middleware:
  - Only users with `"role": "Admin"` can create, update, or delete tasks, or update user roles.
  - Both "Admin" and "user" roles can view tasks.

---

## Error Handling

- The API handles network/database errors and validation errors gracefully.
- If a database/network error occurs, a clear error message is returned.
- Validation errors (e.g., missing fields) are also handled and reported to the client.

---

## Running Unit Tests

- Unit tests for usecases, and infrastructure are provided.
- To run all tests:
  ```sh
  go test -v ./...
  ```
- No need to start the server for unit tests.
- Tests use mocks for dependencies and do not require a running database for usecase/infra tests.

---

## Using the Postman Collection

- **Import the provided Postman collection**:  
  [Task Management API Postman Collection](https://web.postman.co/workspace/ed1fcb1b-aa6d-4608-8bfc-abf010bb0f11/collection/40582744-b2fb455a-9a0a-4cc8-a97e-4e19c73def65?action=share&source=copy-link&creator=40582744)
- **For endpoints that require authentication**, first register and login to obtain a JWT token.
- **For protected endpoints**, set the following header in your Postman requests:
  ```
  Authorization: Bearer <your_jwt_token>
  ```
- Only users with the appropriate role (e.g., "Admin") can access certain endpoints (like updating roles or deleting tasks).

---

## Summary of Recent Changes

- **Refactored to Clean Architecture:** Improved maintainability and scalability by separating handlers, services, repositories, and models.
- **Centralized Environment Loading:** Ensured `.env` is loaded before any DB or JWT operations.
- **Improved Middleware:** JWT and role-based middleware now properly validate and authorize users.
- **Comprehensive Unit Testing:** Added tests for usecases, and infrastructure using Testify and mocks.
- **Updated Documentation:** All instructions and usage notes reflect the new structure, testing, and best practices.

---
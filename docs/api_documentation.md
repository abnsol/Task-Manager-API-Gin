# Task Management API Documentation

## Overview

The Task Management API is a RESTful service built with Go and Gin for managing tasks. It supports user registration, login, JWT authentication, and role-based access control (RBAC) for "Admin" and "user" roles. All data is stored persistently using MongoDB Atlas (cloud MongoDB).

## Features
- Register and login users with JWT authentication
- Role-based access control for endpoints (Admin/user)
- Add, view, update, and delete tasks
- Persistent storage with MongoDB Atlas

## MongoDB Integration
- The API uses **MongoDB Atlas** for cloud-based persistent storage.
- You must provide your MongoDB Atlas **username** and **password** in the `.env` file as part of the `mongoURI`.
- Example connection string (replace `<username>`, `<password>`, and `<cluster-url>`):
  ```
  mongoURI="mongodb+srv://<username>:<password>@<cluster-url>/<database>?retryWrites=true&w=majority"
  ```
- The database used is `Tasks-DataBase` and the collections are `Users` and `Tasks`.

## Authentication & RBAC
- JWT authentication is required for all task endpoints.
- The JWT secret is set in the `.env` file as `jwt_secret`.
- Role-based access control is enforced using middleware:
  - Only users with `"role": "Admin"` can create, update, or delete tasks, or update user roles.
  - Both "Admin" and "user" roles can view tasks.

## Error Handling
- The API handles network/database errors and validation errors gracefully.
- If a database/network error occurs, a clear error message is returned.
- Validation errors (e.g., missing fields) are also handled and reported to the client.

## How to Use
1. **Run the API locally** using Go. The server will start on `http://localhost:8080/` by default.
2. **Set your MongoDB Atlas credentials and JWT secret** in the `.env` file before running the server.
3. **Import the Postman Collection** using the link below to easily test all endpoints without manually crafting requests.
4. **Explore and interact** with the API using Postman. The collection includes all available endpoints and example requests.

**Postman Collection:** [Task Management API Postman Collection](https://web.postman.co/workspace/ed1fcb1b-aa6d-4608-8bfc-abf010bb0f11/collection/40582744-b2fb455a-9a0a-4cc8-a97e-4e19c73def65?action=share&source=copy-link&creator=40582744)

## Note on Environment Variables
- The `.env` file **must** be loaded before any database or JWT operations.
- Do **not** initialize database connections or secrets at the package level if they depend on environment variables.
- Always load `.env` in `main.go` before initializing the database or starting the server.
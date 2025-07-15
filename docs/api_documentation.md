# Task Management API Documentation

## Overview

The Task Management API is a simple RESTful service built with Go and Gin for managing tasks. It allows users to create, retrieve, update, and delete tasks. Now, all data is stored persistently using MongoDB Atlas (cloud MongoDB), ensuring your tasks are saved even after server restarts.

## Features
- Add new tasks
- View all tasks or a specific task by ID
- Update existing tasks
- Delete tasks
- Persistent storage with MongoDB Atlas

## MongoDB Integration
- The API uses **MongoDB Atlas** for cloud-based persistent storage.
- You must provide your MongoDB Atlas **username** and **password** in the `connect_db` function inside `data/task_service.go`.
- Example connection string (replace `<username>` and `<password>`):
  ```go
  clientOptions := options.Client().ApplyURI("mongodb+srv://<username>:<password>@cluster0.mongodb.net/?retryWrites=true&w=majority")
  ```
- The database used is `Tasks-DataBase` and the collection is `Tasks`.

## Error Handling
- The API handles network/database errors and validation errors gracefully.
- If a database/network error occurs, a clear error message is returned.
- Validation errors (e.g., missing fields) are also handled and reported to the client.

## How to Use
1. **Run the API locally** using Go. The server will start on `http://localhost:8080/` by default.
2. **Set your MongoDB Atlas credentials** in `data/task_service.go` before running the server.
3. **Import the Postman Collection** using the link below to easily test all endpoints without manually crafting requests.
4. **Explore and interact** with the API using Postman. The collection includes all available endpoints and example requests.

**Postman Collection:** [Task Management API Postman Collection](https://web.postman.co/workspace/ed1fcb1b-aa6d-4608-8bfc-abf010bb0f11/collection/40582744-b2fb455a-9a0a-4cc8-a97e-4e19c73def65?action=share&source=copy-link&creator=40582744)

## This Project
- Implements a RESTful API using Go and the Gin web framework.
- Uses MongoDB Atlas for persistent, cloud-based storage of tasks.
- Provides full CRUD (Create, Read, Update, Delete) operations for tasks.
- Handles network, database, and validation errors.
- Documented and tested all endpoints using Postman, making it easy for users to try out the API.

## Verifying Data Correctness
- You can verify the correctness of data stored in MongoDB by:
  - Using the API endpoints (e.g., Get All Tasks, Get Task by ID) to retrieve and check data.
  - Querying the MongoDB Atlas database directly using tools like **MongoDB Compass** or the Atlas web interface to inspect the stored documents.

## Getting Started
- Clone the repository and set your MongoDB Atlas credentials in `data/task_service.go`.
- Run the server with Go.
- Use the Postman collection to send requests and see responses instantly.

---
For more details on each endpoint and example requests/responses, please refer to the Postman collection linked above.


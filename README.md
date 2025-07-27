# Task Management API

A RESTful API for managing tasks, built with Go, Gin, and MongoDB Atlas for persistent cloud storage.  
Implements JWT authentication and role-based access control (RBAC) with "Admin" and "user" roles.  
Now refactored using Clean Architecture for better maintainability and scalability.

---

## Features

- **User Authentication:** Register and login with JWT authentication.
- **Role-Based Access Control:** "Admin" and "user" roles enforced via middleware.
- **Task Management:** Add, view, update, and delete tasks.
- **Persistent Storage:** Uses MongoDB Atlas (cloud MongoDB).
- **Error Handling:** Handles network, database, and validation errors gracefully.
- **Clean Architecture:** Separation of concerns between handlers, services, repositories, and models.

---

## Project Structure (Clean Architecture)

```
/Delivery
   /config                 # Configuration and environment loading
   /controllers            # HTTP handlers
   /routers                # routes 
/docs                      # api_documentation
/Domain                    # Entity
/Infrastructure            
/Repository                # Data access (MongoDB logic)
/Usecases                  # Buisness logic (services)

``` 
---

## Environment File Structure

Create a `.env` file in the project root with the following structure:

```
jwt_secret="your_super_secret_jwt_key"
mongoURI="mongodb+srv://<username>:<password>@<cluster-url>/<database>?retryWrites=true&w=majority"
```

- `jwt_secret`: Secret key for signing JWT tokens (must match in both login and middleware).
- `mongoURI`: Your MongoDB Atlas connection string.

---

## Getting Started

1. **Clone the repository**
2. **Create a `.env` file** as described above and fill in your MongoDB Atlas credentials and JWT secret.
3. **Install dependencies**
   ```sh
   go mod tidy
   ```
4. **Run the server**
   ```sh
   go run cmd/main.go
   ```
5. The API will be available at `http://localhost:8080/`
6. **Import the Postman collection** to test all endpoints:
   [Task Management API Postman Collection](https://web.postman.co/workspace/ed1fcb1b-aa6d-4608-8bfc-abf010bb0f11/collection/40582744-b2fb455a-9a0a-4cc8-a97e-4e19c73def65?action=share&source=copy-link&creator=40582744)

---

## Documentation

See [`docs/api_documentation.md`](docs/api_documentation.md) for detailed usage, endpoint descriptions, and integration notes.

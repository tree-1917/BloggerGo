# Go Blog API

A simple RESTful Blog API built with **Go** and **Gin**, featuring **Swagger documentation**.  
This project allows you to **create, read, and delete blog posts**.

---

## Table of Contents

- [Features](#features)  
- [Tech Stack](#tech-stack)  
- [Installation](#installation)  
- [Running the Server](#running-the-server)  
- [API Endpoints](#api-endpoints)  
- [Swagger Documentation](#swagger-documentation)  

---

## Features

- List all blogs  
- Get a blog by ID  
- Create a new blog  
- Delete a blog by ID  
- Auto-generated Swagger documentation  

---

## Tech Stack

- **Go**  
- **Gin Web Framework**  
- **Swagger** via [swaggo/gin-swagger](https://github.com/swaggo/gin-swagger)  

---

## Installation

1. Clone the repository:

```bash
git clone <your-repo-url>
cd blog_apis
````

2. Install dependencies:

```bash
go mod tidy
```

3. Generate Swagger docs:

```bash
swag init
```

> Make sure `swag` is installed. You can install it with:
>
> ```bash
> go install github.com/swaggo/swag/cmd/swag@latest
> ```

---

## Running the Server

Start the server on **port 8080**:

```bash
go run main.go
```

The server will be accessible at:

```
http://localhost:8080
```

---

## API Endpoints

| Method | Endpoint  | Description         |
| ------ | --------- | ------------------- |
| GET    | /blogs    | Get all blogs       |
| GET    | /blog/:id | Get a blog by ID    |
| POST   | /blog     | Create a new blog   |
| DELETE | /blog/:id | Delete a blog by ID |

---

### Example Requests

#### Get all blogs

```bash
curl http://localhost:8080/blogs
```

#### Get a blog by ID

```bash
curl http://localhost:8080/blog/1
```

#### Create a new blog

```bash
curl -X POST http://localhost:8080/blog \
  -H "Content-Type: application/json" \
  -d '{"id":"3","title":"New Blog","content":"Hello Go!"}'
```

#### Delete a blog

```bash
curl -X DELETE http://localhost:8080/blog/3
```

---

## Swagger Documentation

Swagger UI is available at:

```
http://localhost:8080/swagger/index.html
```

This provides interactive API documentation and testing.

---

## License

MIT License


---


# Golang RESTful API

![Golang Logo](https://golang.org/doc/gopher/frontpage.png)

This project is a **RESTful API** built with **Golang**, utilizing the **Gofiber** framework and **GORM** ORM for MySQL database interaction. It's designed for building scalable and efficient web services.

## Features
- **Golang** backend with **Fiber** for fast HTTP handling.
- **MySQL** database with **GORM** for ORM support.
- **Validation** using the `go-playground/validator` package.
- Environment variable management with `godotenv`.

## Dependencies

This project uses the following Go modules:

### Core Dependencies
- **Gofiber** (`github.com/gofiber/fiber/v2`) - Web framework for building APIs with minimal overhead.
- **GORM** (`gorm.io/gorm`) - Object-Relational Mapping (ORM) library for Golang.
- **Validator** (`github.com/go-playground/validator/v10`) - Struct validation library.
- **Godotenv** (`github.com/joho/godotenv`) - Load environment variables from a `.env` file.

### Database
- **MySQL** Driver (`gorm.io/driver/mysql`) - MySQL driver for GORM.

### Cryptography & Utilities
- **Golang Crypto** (`golang.org/x/crypto`) - Cryptographic utilities.

### Other Indirect Dependencies
- **UUID** (`github.com/google/uuid`) - Universally unique identifiers.
- **Compression** (`github.com/klauspost/compress`) - Compression utilities.
- **Brotli** (`github.com/andybalholm/brotli`) - Brotli compression for HTTP.
- **Mimetype** (`github.com/gabriel-vasile/mimetype`) - File type detection.
- **Date Manipulation** (`github.com/jinzhu/now`) - Date/time parsing utilities.

## Installation

### Prerequisites

Ensure you have **Go** installed on your machine:

```bash
go version
```

### Clone the Repository

Clone this repository to your local machine:

```bash
git clone https://github.com/your-username/golang-resfull.git
cd golang-resfull
```

### Setup Environment

Create a `.env` file in the root directory to store environment variables:

```ini
DB_USER=root
DB_PASSWORD=
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=golang
```

### Install Dependencies

Install the necessary Go modules:

```bash
go mod tidy
```

### Running the Application

Start the application with the following command:

```bash
go run main.go
```

The server will start on `http://localhost:8080`.

## Endpoints

- **GET** `/` - Returns a simple "Hello World" message.
- **GET** `/users` - Retrieve all users.
- **POST** `/users` - Create a new user.
- **GET** `/users/{id}` - Retrieve a user by ID.
- **PUT** `/users/{id}` - Update a user by ID.
- **DELETE** `/users/{id}` - Delete a user by ID.

## Contribution

If you would like to contribute, feel free to fork the repository and submit a pull request. Please ensure your changes align with the code style and best practices.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
```

Feel free to copy this directly into your `README.md` file. If you need further customization, just let me know!
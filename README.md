# Golang RESTful API

![Golang Logo](https://golang.org/doc/gopher/frontpage.png)

A simple and efficient RESTful API built with Go, designed to handle user registration, login, and basic user management functionalities. This API uses GORM for MySQL database management and Fiber for handling HTTP requests, alongside other utilities to ensure smooth operations and secure user authentication.

## Requirements

- Go 1.23.4 or higher
- MySQL or compatible database
- `.env` file for environment variables configuration

## Installation

### Clone the Repository

First, clone the repository to your local machine:

```bash
git clone https://github.com/yaosanz/golang-resfull.git
cd golang-resfull
```

### Install Dependencies

This project uses Go modules to manage dependencies. Install all required dependencies by running:

```bash
go mod tidy
```

### Configuration

Create a `.env` file in the root directory of the project and add the following configuration:

```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASS=password
DB_NAME=your_db_name
JWT_SECRET=your_jwt_secret_key
```

### Run the Application

To run the application locally, simply execute:

```bash
go run main.go
```

The API should now be running on `http://localhost:3000`.

## Features

- **User Registration**: Users can register by providing their name, email, phone number, and password.
- **User Login**: Authenticated users can log in with their credentials and receive a JWT token for subsequent requests.
- **Get Users**: Admin can retrieve a list of all registered users.
- **Get User by ID**: Retrieve user details based on their unique ID.
- **Delete User**: Delete a user from the database by their unique ID.

## API Endpoints

### `POST /api/register`
Register a new user.

#### Request Body
```json
{
  "name": "John Doe",
  "email": "john.doe@example.com",
  "phone": "123456789012",
  "password": "yourpassword"
}
```

#### Response
```json
{
  "message": "User registered successfully",
  "user": {
    "id": 1,
    "name": "John Doe",
    "email": "john.doe@example.com",
    "phone": "123456789012"
  }
}
```

### `POST /api/login`
Login to the application and retrieve a JWT token.

#### Request Body
```json
{
  "email": "john.doe@example.com",
  "password": "yourpassword"
}
```

#### Response
```json
{
  "message": "Login successful",
  "token": "your-jwt-token",
  "user": {
    "id": 1,
    "name": "John Doe",
    "email": "john.doe@example.com",
    "phone": "123456789012"
  }
}
```

### `GET /api/users`
Retrieve all users in the system.

#### Response
```json
{
  "message": "Success get all users",
  "users": [
    {
      "id": 1,
      "name": "John Doe",
      "email": "john.doe@example.com",
      "phone": "123456789012"
    }
  ]
}
```

### `GET /api/user/{id}`
Retrieve a specific user by their ID.

#### Response
```json
{
  "message": "Success get user by ID",
  "user": {
    "id": 1,
    "name": "John Doe",
    "email": "john.doe@example.com",
    "phone": "123456789012"
  }
}
```

### `DELETE /api/user/{id}`
Delete a specific user by their ID.

#### Response
```json
{
  "message": "User deleted successfully"
}
```

## Dependencies

This project uses the following modules:

- **[Go Fiber](https://github.com/gofiber/fiber)**: A fast and lightweight web framework for Go.
- **[GORM](https://gorm.io/)**: An ORM (Object-Relational Mapping) library for Go, used for database interactions.
- **[Validator](https://github.com/go-playground/validator)**: A package for data validation.
- **[Go Dotenv](https://github.com/joho/godotenv)**: Loads environment variables from a `.env` file.
- **[JWT](https://github.com/golang-jwt/jwt)**: JSON Web Token (JWT) for handling authentication.
- **[UUID](https://github.com/google/uuid)**: A package to generate UUIDs for user identification.

### Direct Dependencies

```txt
github.com/go-playground/validator/v10 v10.24.0
github.com/gofiber/fiber/v2 v2.52.6
github.com/joho/godotenv v1.5.1
golang.org/x/crypto v0.32.0
gorm.io/driver/mysql v1.5.7
gorm.io/gorm v1.25.12
```

### Indirect Dependencies

```txt
filippo.io/edwards25519 v1.1.0
github.com/andybalholm/brotli v1.1.0
github.com/gabriel-vasile/mimetype v1.4.8
github.com/golang-jwt/jwt/v4 v4.5.1
github.com/google/uuid v1.6.0
```

## Contributing

Contributions are welcome! Feel free to fork the repository, make changes, and create a pull request.

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-name`).
3. Make your changes.
4. Commit your changes (`git commit -am 'Add new feature'`).
5. Push to the branch (`git push origin feature-name`).
6. Open a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
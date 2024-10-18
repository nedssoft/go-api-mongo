# Go BASIC API

This project is a RESTful API built with Go, using Gin as the web framework and MongoDB for database operations.


## Features

- User management (create, read, update, delete)
- Post management (create, read, update, delete)

- End-to-end testing

## Prerequisites

- Go 1.16+
- MongoDB

## Setup

1. Clone the repository:
   ```
   git clone https://github.com/nedssoft/go-api-mongo.git
   cd go-api-mongo
   ```

2. Install dependencies:
   ```
   go mod download
   ```

3. Set up your environment variables in a `.env` file:
   ```
   MONGO_URI=mongodb://localhost:27017
   JWT_SECRET_KEY=your_secret_key
   DB_NAME=go-api-mongo
   PORT=8080
   ```


## Running the Application

To start the server:
```
make run
```

The server will run on port 8080 by default.

## Testing

To run the tests:
```
make test
```

This will run the tests for both the user and post endpoints.


## API Endpoints

- `POST /api/v1/users`: Create a new user
- `GET /api/v1/users/:id/posts`: Get a user and posts
- `POST /api/v1/posts`: Create a new post
- `GET /api/v1/posts`: Get all posts
- `GET /api/v1/posts/:id`: Get a post by ID

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License.

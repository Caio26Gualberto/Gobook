# Gobook API

Gobook API is a social networking application developed in Go, offering a robust infrastructure to build a complete social media platform. This repository contains the backend of the application, which provides RESTful endpoints to manage users, posts, comments, and other social networking functionalities.

## Features

- User registration and authentication
- Create, read, update, and delete posts
- Comments on posts
- Follow other users
- Personalized news feed

## Installation

1. Make sure you have Go installed on your machine. For installation instructions, visit: https://golang.org/doc/install

2. Clone this repository:
    ```bash
    git clone https://github.com/Caio26Gualberto/gobook.git
    ```

3. Navigate to the cloned API directory:
    ```bash
    cd api
    ```

4. Install dependencies:
    ```bash
    go mod tidy
    ```

5. Set up environment variables:
    - Create a `.env` file in the root of the project and define the environment variables as per the example in the `.env.example` file.

6. Start the server:
    ```bash
    go run main.go
    ```

## API Documentation

The complete API documentation is available at [API Docs](docs/api.md).

## Contributing

Contributions are welcome! Feel free to open issues or send pull requests with improvements, bug fixes, or new features.

## License

This project is licensed under the [MIT License](LICENSE).

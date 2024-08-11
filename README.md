
## Golang Books CRUD API

## Overview

This repository contains a Golang API for managing a book library. It provides endpoints for creating, reading, updating, and deleting books. The API utilizes a PostgreSQL database for data persistence. Docker and Docker Compose are used for containerization and environment management.

## Features

CRUD operations for books

PostgreSQL database integration

Dockerized application for easy deployment

Environment configuration using environment variables

## Prerequisites

Go programming language

Docker and Docker Compose

PostgreSQL database (or a PostgreSQL image)

##Getting Started

Clone the repository:
``` bash
git clone https://github.com/vijaymehrotra/Books-crud-api-go.git
```

## Edit the .env file: 
Edit the .env file in the project root directory and add your environment variables:

``` bash
# DB_PASSWORD= [ your password ]

# DB_USER=  [ your user_name ]

# DB_NAME= [ your database_name ]
```

## Build and run the application:
``` bash
docker-compose up --build
```

## API Endpoints

GET /books: Retrieves a list of all books.
POST /books: Creates a new book.
GET /books/:id: Retrieves a specific book by ID.
PUT /books/:id: Updates a book by ID.
DELETE /books/:id: Deletes a book by ID.
Docker Compose
The docker-compose.yml file defines the services for the API and PostgreSQL database. You can customize the database image and port mappings as needed.

## Environment Variables
The API uses environment variables to configure the database connection. You can override these values by setting environment variables before running the application.

## Contributing
Contributions are welcome! Please feel free to fork the repository and submit pull requests.

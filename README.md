
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

GET api/books : Retrieves a list of all books.

GET api/get_book/:id : Retrieves a specific book by ID.

POST api/create_books : Creates a new book.

DELETE api/delete_book/:id : Deletes a book by ID.

``` bash
## An Example for api/books
{
    "Message": "success getting books",
    "books": [
        {
            "id": 1,
            "Author": "David Goddings 2",
            "Title": "Cant Hurt Me 2",
            "Publisher": "Lioncrest 2"
        },
        {
            "id": 2,
            "Author": "Dale Cranage",
            "Title": "How to win friends and Influence people",
            "Publisher": "I dont know"
        }
    ]
}
```

## Docker Compose
The docker-compose.yml file defines the services for the API and PostgreSQL database. You can customize the database image and port mappings as needed.

## Environment Variables
The API uses environment variables to configure the database connection. You can override these values by setting environment variables before running the application.

## Contributing
Contributions are welcome! Please feel free to fork the repository and submit pull requests.

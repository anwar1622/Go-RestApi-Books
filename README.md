# GO REST API BOOKS
This is a simple REST API for books using Go and Gin Framework.

### Dependencies
- Go 1.16 higher
- Gin Framework
- PostgresSQL

### Getting Started
1. Clone this repository to your local machine
2. Create a database in your PostgresSQL and Rename `.env.example` to `.env` and modify the configuration to match your PostgresSQL setup
3. Install the dependencies by running `go mod tidy`
4. Run the application by running `go run install.sh` or `go run main.go`
5. You can see the result in your browser or postman with url `http://localhost:8080/api/v1/books`

### API Documentation
#### `GET` /api/v1/books
Get all books

Endpoint: `/api/v1/books`

Headers:
````
Content-Type: application/json
````

Response:

Status code: `200`

Body:
```` json
{ 
    "status": "success",
    "data": [
        {
            "id": 1,
            "title": "The Lord of The Rings",
            "author": "J.R.R. Tolkien",
        },
        {
            "id": 2,
            "title": "Harry Potter",
            "author": "J.K. Rowling",
        }
    ]
}
````

#### `GET` /api/v1/books/:id
Get a book by id

Request

Endpoint: `/api/v1/books/:id`

Headers:
````
Content-Type: application/json
````

Response

Status code: `200`

Body:
```` json
{ 
    "status": "success",
    "data": {
        "id": 1,
        "title": "The Lord of The Rings",
        "author": "J.R.R. Tolkien",
    }
}
````

#### `POST` /api/v1/books
Create a new book

Request

Endpoint: `/api/v1/books`

Headers:
````
Content-Type: application/json
````

Body:
```` json
{
    "title": "The Lord of The Rings",
    "author": "J.R.R. Tolkien",
}
````

Response

Status code: `201`

Body:
```` json
{ 
    "status": "success",
    "data": {
        "id": 1,
        "title": "The Lord of The Rings",
        "author": "J.R.R. Tolkien",
    }
}
````

#### `PUT` /api/v1/books/:id
Update a book by id

Request

Endpoint: `/api/v1/books/:id`

Headers:
````
Content-Type: application/json
````

Body:
```` json
{
    "title": "The Lord of the Rings: The Fellowship of the Ring",
    "author": "J.R.R. Tolkien",
}
````

Response

Status code: `200`

Body:
```` json
{ 
    "status": "success",
    "data": {
        "id": 1,
        "title": "The Lord of the Rings: The Fellowship of the Ring",
        "author": "J.R.R. Tolkien",
    }
}
````

#### `DELETE` /api/v1/books/:id
Delete a book by id

Request

Endpoint: `/api/v1/books/:id`

Headers:
````
Content-Type: application/json
````

Response

Status code: `200`

Body:
```` json
{ 
    "status": "success",
    "message": "Book successfully deleted",
}
````

### License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details

### Author
- Muhamad Chaerul Anwar

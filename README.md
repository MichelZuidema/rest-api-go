# RESTful API in Go
This is a RESTful API application created in the Go language to learn and practice Go.

# API Requests
### Get all books
`GET /books/`

### Get book by id
`GET /books/{id}`

### Create Book
`POST /addBook`
body:
```json
{
    "id": 2,
    "title": "New Book2",
    "author": "John Doe",
    "description": "This is my first book!"
}
```

### Update Book
`PATCH /book/{id}`
body:
```json
{
    "title": "New Book2",
    "author": "John Doe",
    "description": "This is actually my second book!"
}
```

### Delete Book
`DELETE /book/{id}`




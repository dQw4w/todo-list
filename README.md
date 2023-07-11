# todo-list 
## API Endpoint: /todos

| Request | Function | Input |
| ------- | -------- | ----- |
| GET     | Print to-do list     | -     |
| POST    | Add new to-do        | title, complete     |
| DELETE  | Delete to-do         | id     |
| PUT     | Change to-do's complete status | id     |
| PATCH   | Modify to-do         | id, title, complete |

## Add New To-Do (POST)

### Request

```http
POST /todos
Content-Type: application/json

{
  "title": "Example Task",
  "complete": false
}
```
## Delete To-Do (DELETE)

### Request

```http
DELETE /todos/{id}
```

## Change To-Do's Complete Status (PUT)

### Request

```http
PUT /todos/{id}
```

## Change To-Do's Complete Status (PUT)

### Request

```http
PUT /todos/{id}
```

## Modify To-Do (PATCH)

### Request

```http
PATCH /todos
Content-Type: application/json

{
  "id": ID
  "title": "Updated Task",
  "complete": boolean
}
```

## Print To-Do List (GET)

### Request

```http
GET /todos
```

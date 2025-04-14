# ToDoList API System

## Models
### User
- id (int)
- name (string)
- password (string)

### Task
- id (int)
- title (string)
- description (string)
- is_done (bool)
- user_id (int, foreign key)

## Endpoints

### User
- POST /users
- GET /users
- GET /users/{id}
- PUT /users/{id}
- DELETE /users/{id}
- POST /login

### Task
- POST /tasks
- GET /tasks
- GET /tasks/{id}
- PUT /tasks/{id}
- DELETE /tasks/{id}

## Special Feature
- `/tasks` selalu berelasi dengan user (melalui `user_id`)

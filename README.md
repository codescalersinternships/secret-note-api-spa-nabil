# Secret Note Sharing Application

Create a web application that allows users to securely share self-destructing secret notes.

## Functionalities:

### 1. Note Creation:
- Users can create a note with text content
- Generate a unique, secure URL for each note
- Set an expiration time or number of views before destruction

### 2. Note Retrieval:
- Users can view a note using the unique URL
- After viewing or upon expiration, the note is permanently deleted

### 3. Security Features:
- Use secure random generation for URLs (should be able to browse /note/1, /note/2, etc. The URLs shouldn't be predictable)

### 4. Backend:
- Swagger docs for the API
- Develop a restful API using Go
- Use a database (e.g., sqlite or postgresql) for storing the notes

### 5. Frontend:
- Create a simple, responsive UI using Vue

### 6. Authentication:
- Implement rate limiting to prevent abuse

### 7. Testing:
- Has unit and integration tests for the backend
- Postman collections for testing
- Implement end-to-end tests for critical user flows

### 8. Deployment:
- Containerized the application using Docker
- Has a docker-compose file for easy local deployment

## Extra:
- Notes are encrypted in the database



# Features

- Endpoints on Swagger: /swagger/index.html

# How to Setup

1- Clone repo (install git first [link](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git))

```cmd
git clone https://github.com/codescalersinternships/secret-note-api-spa-nabil/tree/development
```



# How to Run
To run site (get docker [installed](https://docs.docker.com/engine/install/))
```golang
docker compose up
```
when you finnish
```golang
docker compose up
```

and it will run on localhost:8080



## How to Test

```golang
cd backend/
make test
```
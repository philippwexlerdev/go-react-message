# backend-test-template

## Requires
```postgresql```
```golang 1.16```


## Setup

1. Run `docker-compoes up` to start database 
2. Run `go mod tidy`
3. Run `go run main.go` to start web API
4. Run migrate database first to migrate database, SEND GET request to http://localhost:8005/db/migrate
--------------------------------

## APIs

Importing file `GolangTest.postman_collection.json` to test API

### Database
1. Migrate database
```
    localhost:8005/api/user/get/:id
```

### User
Note: when create a user and login, password must be base64 encypt before sending to the backend
1. Create a user 
```
    localhost:8005/api/user/get/:id
```
2. Get a user
```
    localhost:8005/api/user/get/:id
```
3. Update a user
```
    localhost:8005/api/user/update/:id
```
4. Delete a user
```
    localhost:8005/api/user/delete/:id
```
5. Login
```
    localhost:8005/api/user/login
```


### Post
Note: Post has foreign key with user

User using these api to post/edit/delete a message in a message board 

1. Create a post
```
    localhost:8005/api/post/create
```
2. Get a post
```
    localhost:8005/api/post/get/:id
```
3. Get all post
```
    localhost:8005/api/post/get-all
```
4.Update a post
```
    localhost:8005/api/post/update/:id
```
5. Delete a post
```
    localhost:8005/api/post/delete/:id
```


### Post response
Note: Post Respoonse has foreign key with post

User using these api to post/edit/delete a message in a message board

1. Create a post response
```
    localhost:8005/api/post-message/create
```
2. Get a post response
```
    localhost:8005/api/post-message/get/:id
```
3. Get all post response of a post
```
    localhost:8005/api/post-message/get-all
```
4.Update a post response
```
    localhost:8005/api/post-message/update/:id
```
5. Delete a post response
```
    localhost:8005/api/post-message/delete/:id
```
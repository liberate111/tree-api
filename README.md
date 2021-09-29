# Tree API

## Tree API spec
- [Usage](#Usage)
- [Deploy](#Deploy)



## Usage

HTTP Headers
```sh
Content-Type: application/json
```

### health check
test connection
```sh
HTTP GET /
tree-api/
```

### user

#### login
```sh
HTTP POST /api/v1/users/login
tree-api/api/v1/users/login
```
request
```sh
{
 "username": "user-test",
 "password": "password-test"
}
```
response
- success                           200 OK
- body wrong syntax                 400
- invalid username or password      400
```sh
{
 "status": 200,
 "message": "success"
}
```


#### change-password
```sh
HTTP POST /api/v1/users/change-password
tree-api/api/v1/users/change-password
```
request
```sh
{
	"username": "user-test",
	"oldpassword": "old-password",
	"newpassword" : "new-password"
}
```
response
- success                           200 OK
- body wrong syntax                 400
- invalid username or password      400
- server error from hash password   500
- server update fail                500
```sh
{
 "status": 200,
 "message": "success"
}
```



### tree
#### user get trees
```sh
HTTP GET /api/v1/users/:id/trees
tree-api/api/v1/users/:id/trees
```

response
- success                           200 OK
- body wrong syntax                 400
- tree not found                    404
```sh
{
 "status": 200,
 "message": "success",
 "data": {
     "tree":[
         "Tree0001",
         "Tree0002",
         "Tree0003"
     ]
 }
}
```

### admin
### manage user

#### get uuid
```sh
HTTP POST /api/v1/admin/users
tree-api/api/v1/admin/users
```

request
```sh
{
 "uuid": "uuid-secret",
 "username": "user-test"
}
```

response
- success                           200 OK
- body wrong syntax                 400
- unauthorize                       403
- user not found                    404
```sh
{
 "status": 200,
 "message": "success",
 "data": {
     "uuid": "uuid-test"
 }
}
```

#### add user
```sh
HTTP POST /api/v1/admin/users/insert
tree-api/api/v1/admin/users/insert
```

request
```sh
{
 "uuid": "uuid-secret",
 "username": "user-test"
 "password": "password-test"
}
```

response
- success                           200 OK
- body wrong syntax                 400
- unauthorize                       403
- server insert error               500
```sh
{
 "status": 200,
 "message": "user: user-test is created",
}
```

### manage tree

#### add tree
```sh
HTTP POST /api/v1/admin/trees/insert
tree-api/api/v1/admin/trees/insert
```

request
```sh
{
 "uuid": "uuid-secret",
 "treeName" : "Tree0001",
 "owner": "uuid-test"
}
```

response
- success                           200 OK
- body wrong syntax                 400
- unauthorize                       403
- server insert error               500
```sh
{
 "status": 200,
 "message": "tree: Tree0001 is created",
}
```

#### transfer
```sh
HTTP POST /api/v1/admin/trees/transfer
tree-api/api/v1/admin/trees/transfer
```

request
```sh
{
  "uuid": "uuid-secret",
  "treeName" : "Tree0001",
  "username" : "user-test"
}

```

response
- success                           200 OK
- body wrong syntax                 400
- unauthorize                       403
- server insert error               500
```sh
{
 "status": 200,
 "message": "tree: Tree0001 is transfered",
}
```

## Deploy

### Docker
```sh
FROM golang:1.17-alpine AS builder
RUN mkdir /build
ADD controllers /build/controllers
ADD handlers /build/handlers
ADD models /build/models
ADD go.mod go.sum main.go tree-wal.db /build/
WORKDIR /build
RUN apk add --update gcc musl-dev
RUN CGO_ENABLED=1 GOOS=linux go build -o tree-web-server

FROM alpine:3.14
COPY --from=builder /build/tree-web-server /app/
COPY tree-wal.db /app/tree-wal.db
WORKDIR /app
CMD ["./tree-web-server"]
```

### Heroku
```sh
build:
  docker:
    web: Dockerfile

run:
  web: ./tree-web-server
```

### Heroku CLI
```sh
heroku login

heroku create

heroku stack:set container

heroku git:remote -a tree-web-server  
-a = app name

git push heroku master
master = your branch
```
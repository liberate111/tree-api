# Tree API

## Tree API spec
## Usage
## Deploy



## Usage
### user
### tree


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
#### change-password

### tree
#### user get trees


### admin

#### manage user
#### get uuid
#### add user

### manage tree
#### add tree
#### transfer




#### login
```sh
HTTP POST /v1/login
tree-api/v1/login
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
HTTP POST /v1/change-password
tree-api/v1/change-password
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
HTTP POST /v1/users/:id/tree
tree-api/v1/users/:id/tree
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
#### manage user

#### get uuid
```sh
HTTP POST /v1/admin/users
tree-api/v1/admin/users
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
HTTP POST /v1/admin/users/insert
tree-api/v1/admin/users/insert
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
HTTP POST /v1/admin/trees/insert
tree-api/v1/admin/trees/insert
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
HTTP POST /v1/admin/trees/transfer
tree-api/v1/admin/trees/transfer
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
```sh

```
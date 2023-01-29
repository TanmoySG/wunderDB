# Project: wdbv2
wunderDb is a cross-platform JSON-based in-memory data store, inspired by mongoDb. wdb APIs are completely RESTful and eay to communicate to using HTTP requests.
# 📁 Collection: Users 


## End-point: Create Users
### Method: POST
>```
>undefined
>```
### Body (**raw**)

```json
{
    "username": "username",
    "password": "password"
}
```

### 🔑 Authentication noauth

|Param|value|Type|
|---|---|---|


### Response: 200
```json
{
    "action": "createUser",
    "status": "success",
    "error": {}
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Grant Role
### Method: POST
>```
>undefined
>```
### Body (**raw**)

```json
{
    "username": "userone",
    "permissions": {
        "role": "roleOne",
        "on": {
            "databases": "database-name",
            "collections": "collection-name"
        }
    }
}
```

### Response: 200
```json
{
    "action": "grantRole",
    "status": "success",
    "error": {},
    "data": null
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃
# 📁 Collection: Roles 


## End-point: Create Role
### Method: POST
>```
>undefined
>```
### Body (**raw**)

```json
{
    "role": "roleOne",
    "allowed": [
        "createDatabase",
        "grantRole",
        "createRole",
        "listRole"
    ],
    "denied": [
        "addData"
    ]
}
```

### Response: 200
```json
{
    "action": "createRole",
    "status": "success",
    "error": {}
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: List Roles
### Method: GET
>```
>undefined
>```
### Headers

|Content-Type|Value|
|---|---|
|Accept|application/json|


### Response: 200
```json
{
    "action": "listRole",
    "status": "success",
    "error": {},
    "data": {
        "roleOne": {
            "roleId": "roleOne",
            "grants": {
                "globalPrivileges": {
                    "createDatabase": true
                },
                "databasePrivileges": {},
                "collectionPrivileges": {
                    "addData": false
                }
            }
        }
    }
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃
# 📁 Collection: Databases 


## End-point: Create Database
### Method: POST
>```
>undefined
>```
### Body (**raw**)

```json
{
    "name" : "database"
}
```

### Response: 200
```json
{
    "action": "createDatabase",
    "status": "success",
    "error": {}
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Get Database
### Method: GET
>```
>{{URL}}/api/databases/:database
>```
### Response: 200
```json
{
    "action": "readDatabase",
    "status": "success",
    "error": {},
    "data": {
        "collections": {},
        "metadata": {}
    }
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Delete Database
### Method: DELETE
>```
>{{URL}}/api/databases/:database
>```
### Response: 200
```json
{
    "action": "deleteDatabase",
    "status": "success",
    "error": {}
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃
# 📁 Collection: Collections 


## End-point: Create Collection
### Method: POST
>```
>{{URL}}/api/databases/:database/collections
>```
### Body (**raw**)

```json
{
    "name": "test-collection",
    "schema": {
        "type": "object",
        "properties": {
            "name": {
                "type": "string"
            },
            "age": {
                "type": "integer"
            }
        },
        "required": [
            "name",
            "age"
        ]
    }
}
```

### Response: 200
```json
{
    "action": "createCollection",
    "status": "success",
    "error": {}
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Get Collection
### Method: GET
>```
>{{URL}}/api/databases/:database/collections/:collection
>```
### Response: 200
```json
{
    "action": "readCollection",
    "status": "success",
    "error": {},
    "data": {
        "data": {
            "035dd9b2-df2f-4e35-87cf-becb57e2baa9": {
                "data": {
                    "age": 28,
                    "name": "John"
                },
                "metadata": {},
                "id": "035dd9b2-df2f-4e35-87cf-becb57e2baa9"
            },
            "085e0a8d-1e5f-43f8-9373-673f27b38332": {
                "data": {
                    "age": 28,
                    "name": "John"
                },
                "metadata": {},
                "id": "085e0a8d-1e5f-43f8-9373-673f27b38332"
            },
            "c5593663-31e3-446f-9ecc-c9063d1bce74": {
                "data": {
                    "age": 40,
                    "name": "John"
                },
                "metadata": {},
                "id": "c5593663-31e3-446f-9ecc-c9063d1bce74"
            }
        },
        "metadata": {},
        "schema": {
            "properties": {
                "age": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            },
            "required": [
                "name",
                "age"
            ],
            "type": "object"
        }
    }
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Delete Collection
### Method: DELETE
>```
>{{URL}}/api/databases/:database/collections/:collection
>```
### Response: 200
```json
{
    "action": "deleteCollection",
    "status": "success",
    "error": {}
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃
# 📁 Collection: Data 


## End-point: Add Data
### Method: POST
>```
>{{URL}}/api/databases/:database/collections/:collection/data
>```
### Body (**raw**)

```json
{
  "name": "John",
  "age": 25
}
```

### Response: 200
```json
{
    "action": "addData",
    "status": "success",
    "error": {}
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Get Data
### Method: GET
>```
>{{URL}}/api/databases/:database/collections/:collection/data?key=&value=
>```
### Query Params

|Param|value|
|---|---|
|key||
|value||


### Response: 200
```json
{
    "action": "readData",
    "status": "success",
    "error": {},
    "data": {
        "035dd9b2-df2f-4e35-87cf-becb57e2baa9": {
            "data": {
                "age": 28,
                "name": "John"
            },
            "metadata": {},
            "id": "035dd9b2-df2f-4e35-87cf-becb57e2baa9"
        },
        "085e0a8d-1e5f-43f8-9373-673f27b38332": {
            "data": {
                "age": 28,
                "name": "John"
            },
            "metadata": {},
            "id": "085e0a8d-1e5f-43f8-9373-673f27b38332"
        }
    }
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Delete Data
### Method: DELETE
>```
>{{URL}}/api/databases/:database/collections/:collection/data?key=&value=
>```
### Query Params

|Param|value|
|---|---|
|key||
|value||


### Response: 200
```json
{
    "action": "deleteData",
    "status": "success",
    "error": {},
    "data": null
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Update Data
### Method: PATCH
>```
>{{URL}}/api/databases/:database/collections/:collection/data?key=&value=
>```
### Body (**raw**)

```json
{
    "field" : "updated value"
}
```

### Query Params

|Param|value|
|---|---|
|key||
|value||


### Response: 200
```json
{
    "action": "updateData",
    "status": "success",
    "error": {},
    "data": null
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Hello
### Method: GET
>```
>undefined
>```
### 🔑 Authentication noauth

|Param|value|Type|
|---|---|---|


### Response: 200
```json
{
    "action": "ping",
    "status": "success",
    "error": {},
    "data": {
        "message": "✋ hello"
    }
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃
_________________________________________________
Powered By: [postman-to-markdown](https://github.com/bautistaj/postman-to-markdown/)

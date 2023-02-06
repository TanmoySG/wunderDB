# wunderDb

wunderDb is a cross-platform JSON-based in-memory data store, inspired by mongoDb. wdb APIs are completely RESTful and easy to communicate to using HTTP requests.

## Getting Started

Being build with Go, the wdb-server is cross-platform and can run on windows, linux and mac. To start the wdb-server, download the platform-specific binary/executable from the latest release. Then run the binary - this starts the wdb-server with default configurations.

```shell
./wunderdb
```

To test if the instance is running fine, ping the `{URL}/api/hello` endpoint.

```sh
curl --location --request GET 'localhost:8086/api'
```

This should send back a 200 response status and a `✋ hello` message.

### wunderDB Container

To run wunderDB on docker use the official [wunderDB Image](https://github.com/TanmoySG/wunderDB/pkgs/container/wunderdb), that is just ~11MB in size!

Start the docker container by docker pulling the image and run the container with initial configurations.

```sh
docker run ghcr.io/tanmoysg/wunderdb:v1.0.11-test
```

To run wunderDB with configurations, use the `docker compose`. Update the values of the confugurations in the compose file and run as
```sh
docker compose up
```

Other compose files that can be used as `docker compose up -f <filename>`

- [`docker-compose.source.yml`](./../docker-compose.source.yml) - use this to build and run a container from the code in the repository. Great for development purposes.
<!-- - [`docker-compose.source.yml`](./../docker-compose.source.yml) - use this to debug/run container in debug mode (currently not ) -->

### `wdbctl` - CLI Tool for wunderDb

The `wdbctl` tool is a command-line tool to control the wdb-server. Use brew to install the binary (or download the `wdbctl` release binaries), and run the `wdbctl` command followed by `start` to spin up the wdb server with default configurations.

```shell
# install wdbctl
brew tap TanmoySG/TanmoySG
brew install wdbctl

# start wdb-server
wdbctl start
```

wdbctl currently supports following commands

```sh
USAGE:
   wdbctl [global options] command [command options] [arguments...]

COMMANDS:
    start    starts the wdb instance
    version  version of CLI and wunderDb
    help, h  Shows a list of commands or help for one command
```

To start a wdb instance with default configurations.

```sh
wdbctl start
```

To start wdb instance (for the first time) by passing custom configuration.

```sh
wdbctl start --port 8082 --storage '/path/to/wfs' --admin "user:pwd"

// or

wdbctl start -p 8082 -s '/path/to/wfs' -s "user:pwd"
```

Once configurations are set, using the configuration flags to pass custom values would not override the set values. To override the existing configurations use the `-o` flag followed by the config-flags to be overriden.

```sh
wdbctl start -o -p 8081
```


### Configuration

Some of the configurations that wunderDb uses are listed below. These configs can be set up using environemt variable or wdbctl flags.

| Configuration                                                 | Description                                                                             | Environment Variable     | wdbctl Flag                   | Type           | Default                |
|---------------------------------------------------------------|-----------------------------------------------------------------------------------------|--------------------------|-------------------------------|----------------|------------------------|
| Port                                                          | Port where instance should run                                                          | PORT                     | --port, -p  value             | number, int    | 8086                   |
| [Persistent Storage](README.md#persisting-data) Location/Path | Path value to directory to persist data after shutdown                                  | PERSISTANT_STORAGE_PATH  | --storage, -s value           | path, string   | $HOME/wdb/wfs (on mac) |
| Admin ID and Password                                         | Instance Admin Username and Password                                                    | ADMIN_ID, ADMIN_PASSWORD | --admin, -a username:password | string, string | admin, admin           |
| Override Flag                                                 | Once the other config are set, this flag is used to override value as and when required | OVERRIDE_CONFIG          | --overide, -o                 | boolean        | false                  |

### Persisting Data

wunderDb is completely in-memory, that is, all its data read, write operatio happen from/on the runtime memory of the server. But when the server is shutdown, the same data needs to be persisted, so that its not lost between startup and shutdown cycles.

Hence, the data is persisted as JSON Files on the file system. The data is loaded from the files when starting up and data in-memory is dumped while the wdb-server gracefully shuts down.

The Persistent Storage path can be defined by the user, if required, but when not set, data is persisted in the user's home directory, in the `wdb/wfs/` sub-directory.


## Users

Like most databases, wdb Users/User Profiles, `users` are the primary "agents" that commit operations, i.e. to perform most operations the requests would need to be requested by a user that exists in wdb. User profile-led operations also helps in access control, by allowing only certain operations to a user.

Each wdb instance has an **administrator** user, with WDB Super-Admin Role `wdb_super_admin_role`, that grants all available privileges on all entities (all databases and collections). The administrator can perform all operations on all entities.

While starting a wdb instance an `admin` user profile can be created by setting the required credentials, refer to the [configuration details](#configuration) for more. If no configuration is set for admin, the default admin credentials - username and password are set as `admin` and `admin`, respectively.

### Create User

Make POST request to the `/api/users` endpoint, passing username and password to create user.

```http
POST /api/users HTTP/1.1
Content-Type: application/json

{
    "username": "username",
    "password": "password"
}
```

### Login User

To login use the following route with GET request.

```http
GET /api/users/login HTTP/1.1
Accept: application/json
Authorization: Basic 
```

If right credentials are passed it returns `success`, otherwise returns `failure` status and details of error.

### Grant Role to User

To grant a user access to the role on a resource, query the following endpoint, passing the required details.

```http
POST /api/users/grant HTTP/1.1
Authorization: Basic 
Content-Type: application/json

{
    "username": "username",
    "permissions": {
        "role": "rolename",
        "on": {
            "databases": "database",
            "collections": "collection"
        }
    }
}
```

Passing wildcard (`*`) resource in databases or collections grants the user the role on any database or collection.

This action requires authentication, as well as autorization - the user commiting this action must have the `grantRole` privilege.

## Roles

A role grants [privileges](#privileges) to perform a specified actions on a [resource](). To ensure security and fine-grained access control, wdb uses [RBAC or Role-based Access Control](). A user is granted one or more roles that controls the user's access to a resource.

### Creating a Role

To create a `role`, query the following endpoint passing the role name, allowed and denied actions. To perform this action, the user must have the `createRole` privilege.

```http
POST /api/roles HTTP/1.1
Authorization: Basic 
Content-Type: application/json

{
    "role": "rolename",
    "allowed": [
        "createDatabase",
        "grantRole",
        "...",
    ],
    "denied": [
        "addData"
    ]
}
```

### List Roles

To list the roles available in wdb query the following endpoint. An authenticated user requires the `listRoles` privilege to run this action.

```http
GET /api/roles HTTP/1.1
Accept: application/json
Authorization: Basic 
```

## Database

A Database in wunderDb is a group of similar kind of collections.

### Create Database

To create a database in wdb, the user requires the `createDatabase` privilege.

```http
POST /api/databases HTTP/1.1
Authorization: Basic
Content-Type: application/json

{
    "name" : "name-of-db"
}
```

### Read Database

To read/fetch a database in wdb, the user requires the `readDatabase` privilege. It returns the list of collections in the database and other metadata information.

```http
GET /api/databases/{databse-name} HTTP/1.1
Accept: application/json
Authorization: Basic 
```

### Delete Database

To delete a database from wdb, the user requires the `deleteDatabase` privilege.

```http
DELETE /api/databases/{databse-name} HTTP/1.1
Accept: application/json
Authorization: Basic 
```

## Collections

A collection is a group of records/data of same modality (schema). Collections are the primary containers of data.

### Schema

Each collection has a schema that defines its modality and how data in that collection should be structured. In wunderDb schema for a collection is defined using [JSON Schema](https://json-schema.org/) and at the time when collections are created. JSON Schema defines the structure, type and various other standards of the data. Read more on how to define schema using JSON Schema [here](https://json-schema.org/learn/getting-started-step-by-step.html).

### Create Collection

To create a collection in a database, use the following endpoint, passing the schema of the data (in JSON Schema notations) in the body. User must have the `createCollection` access granted on the database where the collection is to be created.

```http
POST /api/databases/{database-name}/collections HTTP/1.1
Authorization: Basic 
Content-Type: application/json

{
    "name": "collection-name",
    "schema": {
        // JSON Schema
        "type": "object",
        "properties": {...},
        "required": [...]
    }
}
```

### Read Collection

To read a collection, the user can use the following request. The user must have `readCollection` access granted on the collection that needs to be read/fetched.

```http
GET /api/databases/{database-name}/collections/{collection-name} HTTP/1.1
Accept: application/json
Authorization: Basic 
```

### Delete Collection

To delete a collection from a database, the user must have `deleteCollection` access granted on the collection that needs to be deleted.

```http
DELETE /api/databases/{database-name}/collections/{collection-name} HTTP/1.1
Accept: application/json
Authorization: Basic 
```

## Data

Records that are complaint/in-line with the collection's schema are stored as data. While reading, data can be filtered using `filters`.

### Filters

Filters are used to (as the name suggests) filter or to create smaller buckets/views of data. Filters are extremely important while updating, deleting or reading specific records based on some conditions. Currently wdb only supports key and value match based filters.

To filter data while reading, updating or deleting, we need to pass the field name to the `key` and the value (of the field) that needs to be matched to the `value`.

Example, `.../data?key:name&value:John`, will filter all records with `name=John`.

### Insert/Add Data

User must have `addData` permission granted on the collection to add data to. Pass the data to add in the body as JSON object.

```http
POST /api/databases/{database}/collections/{collection}/data HTTP/1.1
Authorization: Basic 
Content-Type: application/json

{
  // data
}
```

If the data passes schema validation it is added otherwise returns error.

### Get Data

To fetch/read data user must have `readData` permission granted on the collection.

```http
GET /api/databases/{database}/collections/{collection}/data HTTP/1.1
Accept: application/json
Authorization: Basic 
```

Use filters to fetch specific records based on some condition.

```http
GET /api/databases/{database}/collections/{collection}/data?key={field-name}&value={field-value} HTTP/1.1
```

### Delete Data

Use filters to specify/select the data to be deleted based on the key-value condition. User must have `deleteData` permission granted on the collection to delete data from.

```http
DELETE /api/databases/{database}/collections/{collection}/data?key={field-name}&value={field-value} HTTP/1.1
Accept: application/json
Authorization: Basic 
```

### Update Data

Updating data requires the user to pass the `filters` to specify the data to update as well as the updated values of the fields to change in the body of the request. The user required `updateData` permission granted on the collection.

```http
PATCH /api/databases/{database}/collections/{collection}/data?key={field-name}&value={field-value} HTTP/1.1
Authorization: Basic 
Content-Type: application/json

{
    "field": "updated value"
}
```

## API Response

All requests made to wunderDb returns a response that has the same structure and consists of specific fields. The response consists of the request status, the data returned (if any) and errors (if any). The structure of the response returned by wunderDb is

```jsonl
{
    "action": "action-code",    // privilege performed, eg: addData, readCollection, etc.
    "status": "request status", // success or failure
    "error": {},                // errors returned
    "data": {}                  // data returned
}
```

The API Response also returns the appropriate HTTP status code.

#### Error

If any error is raised by the wunderDb server as reponse, the error returned has the error code and the error stack.

```jsonl
{
    "code": "invalidCredentials",
    "stack": [
        "username/password/token provided is not valid"
    ]
}
```

The `code` field contains the error code. While the `stack` contains the stack of error(s), currently only the latest error is returned.

#### Data

The `data` field contains the data/response returned by the particular action. Like the `getData` action would return the list of records in the `data` field.

Each action has its own format of returning data/messages in the `data` field. Read more about data returned in the API Documentation or Postman Collection examples.

## WunderDB Errors

wunderDb has a defined set of errors in the [`wdbError`](../internal/errors/errors.go) package. These standard set of errors are used through-out the actions for raising and returning any errors, if there is any issue while processing a request.

Read more about the error in the errors documentation.

## API Documentation

Refer to [API Documentation](https://documenter.getpostman.com/view/15618820/2s8Z6yXtAq#92ca810f-9f7d-4d65-8e2e-7941bb1990d0) for more details on the wunderDb API, examples, known errors, and API responses. The Postman Collection JSON, can be downloaded from the API Doc page and can be loaded onto Postman for ease of use.

## Privileges

A privilege is the right to commit a particular action on a wunderDb resource. There are multiple privileges that wdb uses to control access to the actions that can be performed. Multiple privileges are grouped together in a role. Privileges can be allowed or denied while defining a role.

### Privilege Category

In wunderDb privileges are categorized based on their scope.

- Global Privilege
  
  Some privileges don't need an associated resource, they have global scoped, that is, wdb doesn't check if the privilege is granted on a resource or not. Example: the `listRole` privilege is a global privilege, when a user runs the query for listing roles, wdb only checks if the associated privilege is granted on the user or not.
  
- Database Privilege
  
  A Database Privilege is scoped to specific databases. While checking if the user has the access to the action, wunderDb also checks if the privilege is granted on the target database. A role granted on a specific Database would only allow access to that database while blocking access for others. Example, if a user, A is granted a role, R with `readDatabase` privilege on the resource (database) DB1, then the user can only read data from DB1, if user A tried to read database B, it'll be blocked.

- Collection Privileges
  
  A Collection Privilege is scoped to specific collections of a specific databae. While checking if the user has the access to the action, wunderDb also checks if the privilege is granted on the target collection. A role granted on a specific collection would only allow access to that collecttion while blocking access for others.

### Resources

A resource is a database, collection, set of databases and collections, or more system specifc resources like users, roles and permissions.

Some of the Privileges available for use in wunderDb and associated actions.

| Privilege        | Category              | Action                             |
|------------------|-----------------------|------------------------------------|
| createUser       | global privilege      | create user                        |
| createRole       | global privilege      | create roles                       |
| listRole         | global privilege      | list roles in wdb                  |
| createDatabase   | global privilege      | create database                    |
| grantRole        | database privileges   | grant role to user                 |
| readDatabase     | database privileges   | read/fetch database                |
| updateDatabase   | database privileges   | update database                    |
| deleteDatabase   | database privileges   | delete existing database           |
| createCollection | database privileges   | create collection in database      |
| readCollection   | collection privileges | read/fetch collections in database |
| updateCollection | collection privileges | update collections in database     |
| deleteCollection | collection privileges | delete collection from database    |
| addData          | collection privileges | add/insert data in collection      |
| readData         | collection privileges | read/fetch data from collection    |
| updateData       | collection privileges | update data in collection          |
| deleteData       | collection privileges | delete data from collection        |

# wunderDb

wunderDb is a cross-platform JSON-based in-memory data store, inspired by mongoDb. wdb APIs are completely RESTful and eay to communicate to using HTTP requests.

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

#### `wdbctl` - CLI Tool for wunderDb

You may also use the `wdbctl` commandline tool to start the wdb-server. Use brew to install the binary (or download the `wdbctl` release binaries), and run the `wdbctl` command followed by `start` to spin up the wdb server with default configurations.

```shell

# install wdbctl
brew tap TanmoySG/TanmoySG
brew install wdbctl

# start wdb-server
wdbctl start
```

Find more about `wdbctl` here.

<!-- #### Docker -->

<!-- As mentioned in the [root README](../README.md#setup), wunderDb can -->

### Configuration

Some of the configurations that wunderDb uses are listed below. These configs can be set up using environemt variable or wdbctl flags.

| Configuration                                                 | Description                                                                             | Environment Variable     | wdbctl Flag                   | Type           | Default                |
| ------------------------------------------------------------- | --------------------------------------------------------------------------------------- | ------------------------ | ----------------------------- | -------------- | ---------------------- |
| Port                                                          | Port where instance should run                                                          | PORT                     | --port, -p  value             | number, int    | 8086                   |
| [Persistent Storage](README.md#persisting-data) Location/Path | Path value to directory to persist data after shutdown                                  | PERSISTANT_STORAGE_PATH  | --storage, -s value           | path, string   | $HOME/wdb/wfs (on mac) |
| Admin ID and Password                                         | Instance Admin Username and Password                                                    | ADMIN_ID, ADMIN_PASSWORD | --admin, -a username:password | string, string | admin, admin           |
| Override Flag                                                 | Once the other config are set, this flag is used to override value as and when required | OVERRIDE_CONFIG          | --overide, -o                 | boolean        | false                  |

### Persisting Data

wunderDb is completely in-memory, that is, all its data read, write operatio happen from/on the runtime memory of the server. But when the server is shutdown, the same data needs to be persisted, so that its not lost between startup and shutdown cycles.

Hence, the data is persisted as JSON Files on the file system. The data is loaded from the files when starting up and data in-memory is dumped while the wdb-server gracefully shuts down.

The Persistent Storage path can be defined by the user, if required, but when not set, data is persisted in the user's home directory, in the `wdb/wfs/` sub-directory.

<!-- ## wdb: Design

TBD -->

## Users

Like most databases, wdb Users/User Profiles, `users` are the primary "agents" that commit operations, i.e. to perform most operations the requests would need to be requested by a user that exists in wdb. User profile-led operations also helps in access control, by allowing only certain operations to a user.

Each wdb instance has an **administrator** user, with WDB Super-Admin Role `wdb_super_admin_role`, that grants all available privileges on all entities (all databases and collections). The administrator can perform all operations on all entities.

While starting a wdb instance an `admin` user profile can be created by setting the required credentials, refer to the [configuration details](#configuration) for more. If no configuration is set for admin, the default admin credentials - username and password are set as `admin` and `admin`, respectively.

<!-- In wdb users can be added/created and granted roles (with permissions) for access-control using the `user-API`s available.  -->

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

## Privileges

A privilege is the right to commit a particular action on a wunderDb resource. There are multiple privileges that wdb uses to control access to the actions that can be performed. Multiple privileges are grouped together in a role. Privileges can be allowed or denied while defining a role.

### Privilege Category

In wunderDb privileges are categorized based on their scope.

- Global Privilege : Some privileges don't need an associated resource, they have global scoped, that is, wdb doesn't check if the privilege is granted on a resource or not. Example: the `listRole` privilege is a global privilege, when a user runs the query for listing roles, wdb only checks if the associated privilege is granted on the user or not.
- Database Privilege : A Database Privilege is scoped to specific databases. While checking if the user has the access to the action, wunderDb also checks if the privilege is granted on the target database. A role granted on a specific Database would only allow access to that database while blocking access for others.
- Collection Privileges :  A Collection Privilege is scoped to specific collections of a specific databae. While checking if the user has the access to the action, wunderDb also checks if the privilege is granted on the target collection. A role granted on a specific collection would only allow access to that collecttion while blocking access for others.

### Resources

A resource is a database, collection, set of databases and collections, or more system specifc resources like users, roles and permissions.

Some of the Privileges available for use in wunderDb and associated actions.

| Privilege        | Category              | Action | Endpoint |
| ---------------- | --------------------- | ------ | -------- |
| createDatabase   | global privilege      |        |          |
| createRole       | global privilege      |        |          |
| listRole         | global privilege      |        |          |
| grantRole        | database privileges   |        |          |
| deleteDatabase   | database privileges   |        |          |
| readDatabase     | database privileges   |        |          |
| updateDatabase   | database privileges   |        |          |
| createCollection | database privileges   |        |          |
| readCollection   | collection privileges |        |          |
| updateCollection | collection privileges |        |          |
| deleteCollection | collection privileges |        |          |
| addData          | collection privileges |        |          |
| deleteData       | collection privileges |        |          |
| readData         | collection privileges |        |          |
| updateData       | collection privileges |        |          |

## Tools

Here are some of the tools built to help you run and use wunderDb.

### wdbctl

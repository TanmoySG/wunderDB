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

| Configuration                                                 | Description                                                                             | Environment Variable     | wdbctl Flag                     | Type           | Default                |
| ------------------------------------------------------------- | --------------------------------------------------------------------------------------- | ------------------------ | ------------------------------- | -------------- | ---------------------- |
| Port                                                          | Port where instance should run                                                          | PORT                     | --port, -p  value             | number, int    | 8086                   |
| [Persistent Storage](README.md#persisting-data) Location/Path | Path value to directory to persist data after shutdown                                  | PERSISTANT_STORAGE_PATH  | --storage, -s value           | path, string   | $HOME/wdb/wfs (on mac) |
| Admin ID and Password                                         | Instance Admin Username and Password                                                    | ADMIN_ID, ADMIN_PASSWORD | --admin, -a username:password | string, string | admin, admin           |
| Override Flag                                                 | Once the other config are set, this flag is used to override value as and when required | OVERRIDE_CONFIG          | --overide, -o                   | boolean        | false                  |

### Persisting Data

wunderDb is completely in-memory, that is, all its data read, write operatio happen from/on the runtime memory of the server. But when the server is shutdown, the same data needs to be persisted, so that its not lost between startup and shutdown cycles.

Hence, the data is persisted as JSON Files on the file system. The data is loaded from the files when starting up and data in-memory is dumped while the wdb-server gracefully shuts down.

The Persistent Storage path can be defined by the user, if required, but when not set, data is persisted in the user's home directory, in the `wdb/wfs/` sub-directory.

<!-- ## wdb: Design

TBD -->

## Users

Like most databases, wdb uses `users` as the primary "agents" that commit operations, i.e. to perform most operations the requests would need to be requested by a user that exists in wdb. User profile-led operations also helps in access control, by allowing only certain operations to a user.

Each wdb instance has an **administrator** user, with WDB Super-Admin Role `wdb_super_admin_role`, that grants all available privileges on all entities (all databases and collections). The administrator can perform all operations on all entities. 

While starting a wdb instance an `admin` user profile can be created by setting the required credentials, refer to the [configuration details](#configuration) for more. If no configuration is set for admin, the default admin credentials - username and password are set as `admin` and `admin`, respectively.

In wdb users can be added/created and granted roles (with permissions) for access-control using the `user-API`s available. 


## Tools

Here are some of the tools built to help you run and use wunderDb. 

### wdbctl
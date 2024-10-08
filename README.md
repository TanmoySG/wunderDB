# wunderDB

wunderDb is a JSON-based in-memory Data Store.
<!--  For persistent data storage wunderDb loads data from and dumps to filesystem at the start and end of it's lifecycle (startup and shutdown). -->

> [!NOTE]
> We've updated the models for collections, replacing data by records. Please use wunderDb `v1.7.0` with the environment variable `RUN_MODE=RUN_MODE_UPGRADE` to migrate the data to new model, otherwise there will be loss of data.

## Quickstart

To run wunderDb, download the `wunderdb` binary of the [latest release](https://github.com/TanmoySG/wunderDB/releases) based on your OS and Architecture. Once downloaded, run the binary to start wunderdb.

```shell
./wunderdb
```

This should start a wunderDb instance. For configuration documentation, check [this](./documentation/README.md#configuration).

### Running wunderDB Container

To run wunderDB on docker, use the [docker-compose](docker-compose.yml) to start wunderDB with basic configurations.
```shell
docker compose up
```

For more details refer to [this](./documentation/README.md#wunderdb-container).

### wdbctl - CLI Tool

For ease of use, we've also developed a command-line tool for wunderDb - `wdbctl`. 

```sh
# install wdbctl
brew tap TanmoySG/TanmoySG
brew install wdbctl

# starting wunderDB
wdbctl start
```

<!-- To specify configuration while starting an instance, use the flags available, eg: `wdbctl start -p <port>` will start the instance on the port value passed. For more flags and how to use then, run `wdbctl start --help`.

Once set, configurations cant be updated with the configuration flags. To override default or existing configurations, use the override flag `-o` `--override`, followed by the configuration flags, eg: `wdbctl start -o -p 5000` will override the existing/default port and run the instance on port 5000. -->

For more about `wdbctl`, refer to the [documentation](./documentation/README.md#wdbctl).

## Usage

<!-- Once wunderDb instance is running, use the [Admin]() credentials to perform any operations required. For additional security, we recommend creating delegate user(s) with coarse-grained access, to perform the actions. -->

wunderDb APIs are completely RESTful and all actions can be performed using simple HTTP Requests. Refer to the [documentation](./documentation/README.md) for usage instructions.

<!-- Here's an outline of some of the topics in the documentations. -->

## Client Libraries

- [`wdb-go`](https://github.com/TanmoySG/wdb-go) Go client library for wunderDb. [ [Documentation](https://github.com/TanmoySG/wdb-go#readme) ]

<!-- ## Migrations

Certain versions of wdb introduces model changes that need additional migrations from the wdb instance admin. See the [`MIGRATION.md`](./documentation/MIGRATION.md) documentation. -->

## wunderDB-Retro

The first version, based on Python Flask will not be phased out any time soon. To keep it accessible and so that the version 1 doesn't get lost in the version list, I have moved the v0 to a new repository here - [wdb-retro](https://github.com/TanmoySG/wdb-retro).

The wdb-retro repo has all the version 0 code as well as the Docker Image with the new name - so that when I publish the wunderDB v2 Docker Images are published the v1 Image doesnt get lost and anyone planning to use that still can. The version 1 will not be actively maintained.

Check out the final release of v1 <https://github.com/TanmoySG/wdb-retro/releases/tag/v1.1.0>

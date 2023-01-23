# wunderDB

wunderDb is a JSON-based in-memory Data Store. For persistent data storage wunderDb loads data from and dumps to filesystem at the start and end of it's lifecycle (startup and shutdown).

## Setup

To run wunderDb, download the `wunderdb` binary of the [latest release](https://github.com/TanmoySG/wunderDB/releases) based on your OS and Architecture. Once downloaded, run the binary to start wunderdb.

```shell
./wunderdb
```

This should start a wunderDb instance. For ease of used move the binaries to your bin directory.

Configurations can also be passed using environment variables. Refer [Documentation]() for details.

### wdbctl - CLI Tool

For ease of use, we've also developed a command-line tool for wunderDb - `wdbctl`. Install it using homebrew or download the `wdbctl` binary of the latest release.

```sh
brew tap TanmoySG/TanmoySG
brew install wdbctl
```

To start the wunderDb server using wdbctl, run the `start` command. It spins up an instance of wunderDb with default configurations.

```shell
wdbctl start
```

To specify configuration while starting an instance, use the flags available, eg: `wdbctl start -p <port>` will start the instance on the port value passed. For more flags and how to use then, run `wdbctl start --help`.

Once set, configurations cant be updated with the configuration flags. To override default or existing configurations, use the override flag `-o` `--override`, followed by the configuration flags, eg: `wdbctl start -o -p 5000` will override the existing/default port and run the instance on port 5000.

For more about `wdbctl`, refer to the [documentation]().
<!-- 
### Docker

TBD -->

## Usage

Once wunderDb instance is running, use the [Admin]() credentials to perform any operations required. For additional security, we recommend creating delegate user(s) with coarse-grained access, to perform the actions.

wunderDb APIs are completely RESTful and all actions can be performed using simple HTTP Requests. Read the API Documentation for more.

Feel free to browse the documentation for usage instructions and other important concepts used in wunderDb - like coarse-grained access using Role-based Access Control.

<!-- Here's an outline of some of the topics in the documentations. -->

- Users, Roles and Access Control
    - Users and Roles
    - Role-Based Access Control
- Databases, Collections and Data
    - Databases
    - Collections
    - Data
- 


## wunderDB-Retro

The first version, based on Python Flask will not be phased out any time soon. To keep it accessible and so that the version 1 doesn't get lost in the version list, I have moved the v1 to a new repository here - [wdb-retro](https://github.com/TanmoySG/wdb-retro).

The wdb-retro repo has all the version 1 code as well as the Docker Image with the new name - so that when I publish the wunderDB v2 Docker Images are published the v1 Image doesnt get lost and anyone planning to use that still can. The version 1 will not be actively maintained.

Check out the final release of v1 <https://github.com/TanmoySG/wdb-retro/releases/tag/v1.1.0>

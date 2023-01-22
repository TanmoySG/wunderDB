# wunderDB

wunderDb is a JSON-based in-memory Data Store. For persistent data storage wunderDb loads data from and dumps to filesystem at the start and end of it's lifecycle (startup and shutdown). 

## Setup

To run wunderDb, download the `wunderdb` binary of the [latest release](https://github.com/TanmoySG/wunderDB/releases) based on your OS and Architecture. Once downloaded, run the binary to start wunderdb.

```shell
./wunderdb
```

This should start a wunderDb instance. For ease of used move the binaries to your bin directory.

wunderDb has a few configurations that are required to run the instance - port (default to 8086), persitant file storage (default to ~/wdb/wfs). To pass the configurations use environment variables - read [this]() for more.

### wdbctl - CLI Tool

For ease of use, we've also developed a command-line tool for wunderDb - `wdbctl`. Install it using homebrew or download the `wdbctl` binary of the latest release.

```sh
brew tap TanmoySG/TanmoySG
brew install wdbctl
```

To start the wunderDb server using wdbctl, run the `start` command.

```shell
wdbctl start
```
To start wunderDb by overriding default and existing configuration values, use the override flag `-o` , followed by `-p <port>` and/or `-s <path>`.

### Docker
TBD

## Usage




## wunderDB-Retro

The first version, based on Python Flask will not be phased out any time soon. To keep it accessible and so that the version 1 doesn't get lost in the version list, I have moved the v1 to a new repository here - [wdb-retro](https://github.com/TanmoySG/wdb-retro).

The wdb-retro repo has all the version 1 code as well as the Docker Image with the new name - so that when I publish the wunderDB v2 Docker Images are published the v1 Image doesnt get lost and anyone planning to use that still can. The version 1 will not be actively maintained.

Check out the final release of v1 <https://github.com/TanmoySG/wdb-retro/releases/tag/v1.1.0>

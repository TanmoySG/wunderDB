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

##
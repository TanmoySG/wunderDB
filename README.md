<p align="center">
   <img width="50%" src="https://github.com/TanmoySG/wunderDB/blob/main/showcase/wdb-complete.png">
</p>

# wunderDB - a Document Database

***wunderDB*** is a JSON-based micro Document DB, inspired by MongoDB.

#### Deployed at [wdb.tanmoysg.com](http://wdb.tanmoysg.com)

## Documentation

Create a cluster and get started with wunderDB at [wdb.tanmoysg.com](http://wdb.tanmoysg.com).

The cluster can be accessed using the ```Unified Actions API```. To consume this API, use the following endpoint :
```
wdb.tanmoysg.com/connect?cluster=<cluster-id>&token=<one-of-the-three-tokens-generated>
```
The operations on this API are facilitated through ```Actions``` 

**Detailed documentation for wunderDB available [here](https://github.com/TanmoySG/wunderDB/blob/master/documentation/documentation.md)**

## Running WDB Using Docker

Download the docker-compose file and run Docker Compose Up to run wdb locally. But before that create a secrets folder and create `server-config.json`
```
docker-compose up
```


## Features

- Unified API Endpoint.
- Registration Portal for creating Cluster.
- ***Unified Actions API*** enables performing multiple operation through a similar structured API Call.
- **Remotely Hosted**, enabling on-the-go access of data.
- Reduces need of self-hosted servers/databases.
- Micro structure & Lightweight Database.
- Supports **multiple Databases** for a single cluster.
- Supports **multiple Collections** for each Database.
- **C**reate, **R**ead, **U**pdate, **D**elete operations on Data.
- **Schema Integrity Protection** for Data Creation & Updation reducing data-schema mismatch issues.
- **Markers** ensure easy pointing to specific data.
- **Summarised Reporting** of Databases & Collection.

## Progress

- [x] wunderCP, Dashboard for wunderDB -named wunderDash
- [ ] Query Language
- [ ] Mass Data Actions
- [ ] Advanced data Security & Protection
- [ ] Media Access Support
- [x] Registration Portal for creating Cluster.
- [x] Schema Integrity Protection
- [x] Summarised Reporting
- [x] Unified Actions API

#### Suggest/Request a feature that you want to see in wunderDB [here](mailto:tanmoysps@gmail.com)



Project by ***[Tanmoy Sen Gupta](https://www.tanmoysg.com)***



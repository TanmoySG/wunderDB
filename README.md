# wunderDB
wunderDB is a JSON-based micro Document DB, inspired by MongoDB.
version ***1.0 Beta***

New in this Beta Release - ***Unified Endpoint access for all actions***


## Documentation

wunderDB is a JSON-based micro Document DB hosted at [wdb.tanmoysg.com](https://wdb.tanmoysg.com). 

### Creating a Cluster

On registration, a cluster is created. A cluster can also be created by Posting a request to the API endpoint: <kbd>wdb.tanmoysg.com/register</kbd> from a REST Client like [Insomnia](https://insomnia.rest/) or [Postman](https://www.postman.com/) with the following JSON Data.

```
{  
    "name" : "name of the user",
    "email" : "email of the user",
    "password" : "password of user" 
}
```

On successfull registration, a Cluster ID and 3 Access Tokens are generated. These are used for accessing the API.

### Accessing the Cluster - Endpoint

The cluster can be accessed using the ```Unified Actions API```. To consume this API, use the following endpoint :
```
wdb.tanmoysg.com/connect?cluster=<cluster-id>&token=<one-of-the-three-tokens-generated>
```
The operations on this API are facilitated through ```Actions``` .

### Actions & Payloads

Actions and Payloads together form the backbone of the ```Unified Actions API```. While ```Actions``` facilitate operations on the Database, ```Payloads``` are used as specifications to specify data, selectors & configurations. 

- Create Operations
  * **Creating Databases**
  
    Action - <kbd>create-database</kbd> - Used for creating Databases.
  
    Payloads - <kbd>name</kbd> - Name of Database.
    
    ```
    {
        "action" : "create-database"
    }
    ```
  * <kbd>create-collection</kbd> - Used for creating Collections. 

- View Operations
  * <kbd>get-cluster</kbd> - Used for creating Databases.
  * <kbd>get-database</kbd> - Used for creating Databases.
  * <kbd>get-collection</kbd> - Used for creating Databases.
  
- Data Operations
  * <kbd>add-data</kbd> - Used for Adding new data to a Collection.
  * <kbd>update-data</kbd> - Used for Updating existing data in a Collection.
  * <kbd>delete-data</kbd> - Used for Deleting existing data in a Collection.
  * <kbd>get-data</kbd> - Used for Fetching existing data from a Collection.
     


Project by ***[Tanmoy Sen Gupta](https://www.tanmoysg.com)***

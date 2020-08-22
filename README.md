# wunderDB
wunderDB is a JSON-based micro Document DB, inspired by MongoDB.
version ***1.0 Beta***

New in this Beta Release - ***Unified Endpoint access for all actions***


## Documentation

wunderDB is a JSON-based micro Document DB hosted at [wdb.tanmoysg.com](https://wdb.tanmoysg.com). 

### Accessing the API of the database

The database can be accessed using a <kbd>Common Unified Endpoint</kbd> using the **Cluster ID** of the cluster created for the user and an **Access Token**.

The Cluster ID and Access Token are generated on user registration. A cluster can also be created by Posting a request to the API endpoint: <kbd>wdb.tanmoysg.com/register</kbd> from a REST Client like [Insomnia](https://insomnia.rest/) or [Postman](https://www.postman.com/) with the following JSON Data.

<code>
  {
  
    "name" : "name of the user",
  
    "email" : "email of the user",
    
    "password" : "password of user"
    
  }
</code>

Developed by ***[Tanmoy Sen Gupta](https://www.tanmoysg.com)***

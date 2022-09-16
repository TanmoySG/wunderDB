# Proposals for v2 of WunderDB

Tracked in Issue [#6](https://github.com/TanmoySG/wunderDB/issues/6)

## Development Maturity 

- Replace the "Unified Actions Endpoint/API" with entity (cluster, DB, collection, etc) differentiated API endpoints and Actions Routes - `v2/{entity}/{action}` . Example
    - Cluster EPs can be `v2/cluster/create`, `v2/cluster/delete`. 
    - Similarly for data, EPs can be `v2/data/add`, `v2/data/delete` etc.
-  EPs can also be differentiated Endpoints and action routes can replaced with HTTP Verbs. Like
    - GET for Read/Get Data (or any other entity value)
    - POST for adding Data (or any other entity value)
    - PATCH for updating Data (or any other entity value)
    - DELETE for deleting Data (or any other entity value)
    - [Reference for HTTP Verbs](https://www.restapitutorial.com/lessons/httpmethods.html#:~:text=The%20primary%20or%20most%2Dcommonly,but%20are%20utilized%20less%20frequently.)
    - Example - `POST v2/data/` to add data , `PATCH v2/data/` to update data, and more.


## Code Modularity

- Break and Modularize Existing Code into Smaller chunks.
- Club similar or of same "origin" into same "packages". Eg. Code for cluster actions should be clubbed in a "clusters" package.
- Packages should be separated into files for modularity.
- Common Reusable code - "helper" code should be put in separate packages accessible to any components.
- Utility Functions/Code should also be part of helpers

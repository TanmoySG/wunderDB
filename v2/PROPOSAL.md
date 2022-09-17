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


## Code Maturity 

- Use Design Patterns that hide implementation and can facilitate usage without worrying much about the underlying flow.
- One such design pattern is - [Composite Design Pattern](https://refactoring.guru/design-patterns/composite) that can help in designing independent code blocks that can have zero (or little to none) dependancy on other components. These smaller chunks then come together in a "flow" performing their own actions and forming the end-to-end flow for a broader task, instead of a monolithic single code block to perform every task itself.
- The Facade design pattern can be used to hide implementation of any task and only focus on taking a set of arguments and processing it and providing the output without the user worrying about the implementation internally. 
- Code Reusability is another factor to be taken care of. 


## Schema

Currently the schemas are defined as `fieldName : description` and doesn't handle data-type, constraints. In v2,

- Option for using Schema-less and schema-full Collections will be allowed.
- Schema-Less Collections will have an open format, no/low constraints data storage, providing users flexibility when required. Users will be able to put/store any non-structured or non-schema data.
- Schema-Full Collections will have Strict Schema Enforcing with schema defined using `JSONSchema`. Each Data block will have to be complaint with the schema and must strictly follow the schema defined. Any data that isn't in accordance to the JSONSchema will be considered "illegal" and will not be inserted into the Collection.
- Schema-Full Collections will be able to use any one of the fields as primary key or identifier instead of the `_id` auto-generated ID field used as identifier in `v1`
- Schema-less Collections will still have the `_id` as primary identifier.
- JSONSchema - https://json-schema.org/
- Also can use a "fork" of JSONSchema - https://github.com/TanmoySG/w-service-manager/blob/service-onboarding/schema/schema.template.json
- Tutorial for Schema https://www.mongodb.com/basics/json-schema-examples#:~:text=JSON%20Schema%20is%20an%20IETF,validity%20across%20similar%20JSON%20data.
- Schema Definition will be JSON Doc based.
- (From a UI PoV) Schema Definition will no longer be per field basic, and will have a Single Text Box to put JSON Schema in. 
- As additional UI feature (feature extension) UI can provide a way to generate JSON Schema from Example JSON Object and a Schema Builder UI to create JSON Schema using text boxes (field name), drop-downs (data types) and other UI elements for other Schema Definition Utilities (constraints, limit, etc)
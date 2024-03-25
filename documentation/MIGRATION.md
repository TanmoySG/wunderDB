# Migrations

With recent changes in the docker image we've discontinued the use of `wdb-tools`. Migration documentation and mechanism witll be updated soon.
<!-- 
This document contains background details and instructions to migrate to new models and breaking changes.

## Migrating existing roles to add `hidden` field

To update existing roles to add the `hidden` field, there are several methods.

### Option 1: Manually Update Persisted Roles File

Before starting wdb-instance, manually update the roles in the persisted roles file.

- Go to the Persisted Storage Path - where all wdb files are persisted.
- Generally its stored in the `~/wdb/wfs` directory
- Go to the `roles` sub-directory (`~/wdb/wfs/roles`) and open the `roles_persisted.json` file
- Now manually update the roles and add the `hidden` field with the required value - `true` or `false`.


### Option 2: Use wdb-tools to Update the Roles

wdb-tools are commandline tools built to make life easy! To use wdb-tools, 

- Clone the wdb-sidekicks repository - `git clone https://github.com/TanmoySG/wdb-sidekicks`
- Browse to `wdb-sidekicks/tools` directory and build the tools by running - `sh ./build.sh`
- Using the `roles_hidden_field_update` - created in the `wdb-sidekicks/tools/bin` subdirectory, you can update the roles.
- Use the `roles_hidden_field_update` as
  - export the `USE_TOOL` variable as true, `export USE_TOOL=true`
  - export the `TOOL_INSTRUCTION` variable as, `export TOOL_INSTRUCTION=/path/to/roles_hidden_field_update -f path/to/wfs/roles/roles_persisted.json true`
    - Here the value of `-f` flag is the path to the persisted roles file
    - The last argument is the value to be set for the `hidden` field. It can be set as either true or false.
    - All roles will be set to the value
  - Run the `sh ./scripts/start.sh` to run the tool and start the wdb-server

### Option 3: Use wdb-tools in Docker Containers

Simillar to Option-2, to use wdb-tools in wdb-container, you just need to export the `TOOL_INSTRUCTION` and `USE_TOOL` variables. The wdb-image comes pre-packed with the wdb-tools so you dont need to clone it. The tools are stored in the `/tools` directory in wdb-image/container.

- The docker-compose files have the required values set, to use the tool, you just need to change the value of `USE_TOOL` to true (set as false in the compose file)
- This will run the tool inside the container and start the wdb-server -->

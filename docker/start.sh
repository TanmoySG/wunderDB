# Script to run wdb-server
# Usage:
#       sh ./docker/start.sh <optional path to wdb_server binary>

BASEDIR=$(dirname "$0")

PATH_TO_WDB_SERVER=$1
if [[ -z "$PATH_TO_WDB_SERVER" ]]; then
    PATH_TO_WDB_SERVER=bin
fi

# Placeholder for server pre-start jobs/tasks
#
# add command here...

# block to use and run wdb-tools, before starting server
# to use tools,
#       export USE_TOOL=true
#       export TOOL_INSTRUCTION="shell commands to run"
if [[ ! -z "$USE_TOOL" ]]; then
    if [[ "$USE_TOOL" == "true" ]]; then
        if [[ ! -z "$TOOL_INSTRUCTION" ]]; then
            $TOOL_INSTRUCTION
        fi
    fi
fi

# add command here...
#
# Placeholder for server pre-start jobs/tasks

# start wdb-server binary
$PATH_TO_WDB_SERVER/wdb_server

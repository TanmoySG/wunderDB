# block to use and run wdb-tools, before starting server
# to use tools,
#       export USE_TOOL=true
#       export TOOL_INSTRUCTION="shell commands to run"
if [[ ! -z "$USE_TOOL" ]]; then
    if $USE_TOOL; then 
        if [[ ! -z "$TOOL_INSTRUCTION" ]]; then
            eval $TOOL_INSTRUCTION
        fi
    fi
fi
package runmode

type RUN_MODE_TYPE string

const (
	RUN_MODE_MAINTENANCE RUN_MODE_TYPE = "RUN_MODE_MAINTENANCE"
	RUN_MODE_NORMAL      RUN_MODE_TYPE = "RUN_MODE_NORMAL"
	// RUN_MODE_UPGRADE     RUN_MODE_TYPE = "RUN_MODE_UPGRADE"
)

type INSTRUCTION_TYPE string

const (
	INSTRUCTION_TYPE_CMD INSTRUCTION_TYPE = "INSTRUCTION_TYPE_CMD"
)

type RunInstructions struct {
	ToolName        *string `json:"tool_name"`
	ToolPath        *string `json:"tool_path"`
	ToolArgs        *string `json:"tool_args"`
	InstructionType *string `json:"instruction_type"`
	Instructions    *string `json:"instructions"`
}

// convert to interface and client
func ShouldEnterRunMode(rm RUN_MODE_TYPE) bool {
	switch rm {
	case RUN_MODE_MAINTENANCE:
		return true
	// case RUN_MODE_UPGRADE:
	// 	return true
	case RUN_MODE_NORMAL:
		return false
	default:
		return false
	}
}

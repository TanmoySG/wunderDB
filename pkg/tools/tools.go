package tools

import (
	"fmt"
	"os/exec"
	"strings"
)

type tool struct {
	name string
	path string
	args int
}

func Use(toolPath string) (*tool, error) {
	c := exec.Command(toolPath, "usage")
	usageGetterOutput, err := c.Output()
	if err != nil {
		return nil, err
	}

	usageGetterOutputArray := strings.Split(string(usageGetterOutput), " ")
	toolName, toolArgsCount := usageGetterOutputArray[0], len(usageGetterOutputArray)-1

	return &tool{
		name: toolName,
		path: toolPath,
		args: toolArgsCount,
	}, nil
}

func (t *tool) Execute(args ...string) (interface{}, error) {
	if len(args) != t.args {
		return nil, fmt.Errorf("argument count mismatch")
	}

	c := exec.Command(t.path, args...)
	toolExecOutput, err := c.Output()
	if err != nil {
		return nil, err
	}

	return toolExecOutput, nil
}

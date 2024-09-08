package upgrades

import (
	"fmt"

	"github.com/TanmoySG/wunderDB/internal/config"
	"github.com/TanmoySG/wunderDB/internal/upgrades/steps"
)

func Upgrade(c config.Config) error {
	err := steps.Run(c)
	if err != nil {
		return err
	}

	fmt.Println("Upgrade completed successfully...")
	return nil
}

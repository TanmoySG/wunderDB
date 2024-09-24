package steps

import (
	"fmt"

	"github.com/TanmoySG/wunderDB/internal/config"
	v "github.com/TanmoySG/wunderDB/internal/version"
	"github.com/hashicorp/go-version"
)

func Run(c config.Config) error {
	currentVersion, err := version.NewVersion(v.WDB_VERSION)
	if err != nil {
		return err
	}

	constraints, err := version.NewConstraint(">= 1.6")
	if err != nil {
		return err
	}

	if constraints.Check(currentVersion) {
		fmt.Println("Running Upgrades...")
		return MigrateDataToRecords(c)
	} else {
		fmt.Println("No upgrades to run...")
	}

	return nil
}

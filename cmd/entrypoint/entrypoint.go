package entrypoint

import (
	"fmt"

	"github.com/TanmoySG/wunderDB/internal/config"
	"github.com/TanmoySG/wunderDB/internal/runmode"
	"github.com/TanmoySG/wunderDB/internal/server/lifecycle/shutdown"
	"github.com/TanmoySG/wunderDB/internal/server/lifecycle/startup"
	"github.com/TanmoySG/wunderDB/internal/upgrades"
	"github.com/charmbracelet/lipgloss"
)

func EntryPoint() error {
	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#7D56F4")).PaddingTop(1)

	fmt.Println(style.Render("Starting wunderDb..."))

	c, err := config.Load()
	if err != nil {
		return fmt.Errorf("error loading configurations: %s", err)
	}

	if runmode.ShouldEnterRunMode(runmode.RUN_MODE_TYPE(c.RunMode)) {
		if c.RunMode == string(runmode.RUN_MODE_MAINTENANCE) {
			mm, err := runmode.NewMaintenanceMode(*c)
			if err != nil {
				return fmt.Errorf("error initializing Maintenance Mode: %s", err)
			}

			err = mm.EnterMaintenanceMode()
			if err != nil {
				return fmt.Errorf("error entering Maintenance Mode: %s", err)
			}
		} else if c.RunMode == string(runmode.RUN_MODE_UPGRADE) {
			err := upgrades.Upgrade(*c)
			if err != nil {
				fmt.Printf("error initializing Upgrade Mode: %s", err)
			}

			c.RunMode = string(runmode.RUN_MODE_NORMAL)

			err = config.Unload(c)
			if err != nil {
				fmt.Printf("error unloading configurations: %s", err)
			}
		}
	}

	w, err := startup.Prepare(*c)
	if err != nil {
		return fmt.Errorf("error starting wdb server: %s", err)
	}

	shutdown.Listen(*w, *c) // listens to shutdown signals
	startup.Start(w, c)     // starts server and initial setup

	return nil
}

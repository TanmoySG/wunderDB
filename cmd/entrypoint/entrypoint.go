package entrypoint

import (
	"fmt"

	"github.com/TanmoySG/wunderDB/internal/config"
	"github.com/TanmoySG/wunderDB/internal/runmode"
	"github.com/TanmoySG/wunderDB/internal/server/lifecycle/shutdown"
	"github.com/TanmoySG/wunderDB/internal/server/lifecycle/startup"
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
		// add logic to enter run mode
		// runmode.IsValidUpgradeInstruction()

		mm, err := runmode.NewMaintenanceMode(*c)
		if err != nil {
			return fmt.Errorf("error initializing Maintenance Mode: %s", err)
		}

		err = mm.EnterMaintenanceMode()
		if err != nil {
			return fmt.Errorf("error entering Maintenance Mode: %s", err)
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

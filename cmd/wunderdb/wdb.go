package main

import (
	"github.com/TanmoySG/wunderDB/internal/config"
	"github.com/TanmoySG/wunderDB/internal/server/lifecycle/shutdown"
	"github.com/TanmoySG/wunderDB/internal/server/lifecycle/startup"
	"github.com/charmbracelet/lipgloss"
	log "github.com/sirupsen/logrus"
)

func main() {
	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#7D56F4")).PaddingTop(1)

	log.Println(style.Render("Starting wunderDb..."))

	c, err := config.Load()
	if err != nil {
		log.Fatalf("error loading configurations: %s", err)
	}

	w, err := startup.Prepare(*c)
	if err != nil {
		log.Fatalf("error starting wdb server: %s", err)
	}

	shutdown.Listen(*w, *c) // listens to shutdown signals
	startup.Start(w, c)     // starts server and initial setup
}

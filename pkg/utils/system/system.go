package system

import (
	"fmt"
	"os"
	"os/user"
	"runtime"
)

func GetUserHome(operatingSystem string) string {
	switch operatingSystem {
	case WINDOWS:
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	case LINUX:
		home := os.Getenv("XDG_CONFIG_HOME")
		if home != "" {
			return home
		}
	case DARWIN:
		return os.Getenv("HOME")
	}
	return os.Getenv("HOME")
}

func GetHostOS() (string, error) {
	currentOs := runtime.GOOS
	switch currentOs {
	case WINDOWS:
		return WINDOWS, nil
	case DARWIN:
		return DARWIN, nil
	case LINUX:
		return LINUX, nil
	case UNIX:
		return UNIX, nil
	}
	return "", fmt.Errorf("unsupported os: %s", currentOs)
}

func GetCurrentUser() (string, string, error) {
	user, err := user.Current()
	if err != nil {
		return "", "", fmt.Errorf("error fetching %s", err)
	}
	return user.Username, user.HomeDir, nil
}

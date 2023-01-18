package lifecycle

import (
	"fmt"

	"github.com/TanmoySG/wunderDB/pkg/fs"
	"github.com/TanmoySG/wunderDB/pkg/utils/system"
)

const (
	WDB_ROOT_PATH_FORMAT = "%s/wdb"

	WDB_LOGS_DIR_PATH_FORMAT    = "%s/logs"
	WDB_CONFIG_DIR_PATH_FORMAT  = "%s/configs"
	WDB_CONFIG_FILE_PATH_FORMAT = "%s/conf.json"

)

func FirstLaunch() error {
	hostOS, err := system.GetHostOS()
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	homeDir := system.GetUserHome(hostOS)
	wdbRootDirectory := fmt.Sprintf(WDB_ROOT_PATH_FORMAT, homeDir)
	err = fs.CreateDirectory(wdbRootDirectory)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	wdbConfigDirectory := fmt.Sprintf(WDB_CONFIG_DIR_PATH_FORMAT, wdbRootDirectory)
	err = fs.CreateDirectory(wdbConfigDirectory)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	wdbConfigFilePath := fmt.Sprintf(WDB_CONFIG_FILE_PATH_FORMAT, wdbConfigDirectory)
	err = handleFile(wdbConfigFilePath)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	return nil
}

func handleFile(filePath string) error {
	if !fs.CheckFileExists(filePath) {
		err := fs.CreateFile(filePath)
		if err != nil {
			return err
		}

		err = fs.WriteToFile(filePath, []byte("{}"))
		if err != nil {
			return err
		}
	}
	return nil
}

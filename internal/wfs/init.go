package wfs

import (
	"fmt"

	"github.com/TanmoySG/wunderDB/internal/config"
	"github.com/TanmoySG/wunderDB/pkg/fs"
	"github.com/TanmoySG/wunderDB/pkg/utils/system"
)

func (wfs WFileSystem) InitializeFS() error {
	hostOS, err := system.GetHostOS()
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	homeDir := system.GetUserHome(hostOS)
	wdbRootDirectory := fmt.Sprintf(config.WDB_ROOT_PATH_FORMAT, homeDir)

	wdbConfigDirectory := fmt.Sprintf(config.WDB_CONFIG_DIR_PATH_FORMAT, wdbRootDirectory)
	err = fs.CreateDirectory(wdbConfigDirectory)
	if err != nil {
		return err
	}

	wdbConfigFilePath := fmt.Sprintf(config.WDB_CONFIG_FILE_PATH_FORMAT, wdbConfigDirectory)
	err = handleFile(wdbConfigFilePath)
	if err != nil {
		return err
	}

	wfsDirectoryPath := wfs.wfsBasePath
	for _, entity := range config.WFS_DIRECTORIES {
		wdbEntityDirectory := fmt.Sprintf(config.WFS_PERSISTANT_ENTITY_DIR_PATH_FORMAT, wfsDirectoryPath, entity)
		err = fs.CreateDirectory(wdbEntityDirectory)
		if err != nil {
			return err
		}

		wdbEntityFile := fmt.Sprintf(config.WFS_PERSISTANT_ENTITY_FILE_PATH_FORMAT, wdbEntityDirectory, entity)
		err = handleFile(wdbEntityFile)
		if err != nil {
			return err
		}
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

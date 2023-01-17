package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/TanmoySG/wunderDB/pkg/fs"
	"github.com/TanmoySG/wunderDB/pkg/utils/system"
)

var WFS_DIRECTORIES = []string{"users", "namespaces", "databases", "roles"}

const (
	WDB_ROOT_PATH_FORMAT = "%s/wdb"

	WDB_LOGS__DIR_PATH_FORMAT   = "%s/logs"
	WDB_CONFIG_DIR_PATH_FORMAT  = "%s/configs"
	WDB_CONFIG_FILE_PATH_FORMAT = "%s/conf.json"

	WDB_PERSISTANT_STORAGE_DIR_PATH_FORMAT = "%s/wfs"

	WFS_PERSISTANT_ENTITY_DIR_PATH_FORMAT  = "%s/%s"
	WFS_PERSISTANT_ENTITY_FILE_PATH_FORMAT = "%s/%s_persisted.json"
)

type Config struct {
	AdminID               string `json:"-"`
	AdminPassword         string `json:"-"`
	Port                  string `json:"PORT"`
	PersistantStoragePath string `json:"PERSISTANT_STORAGE_PATH"`
}

type ConfigMap map[string]string

func Load() (*Config, error) {
	hostOS, err := system.GetHostOS()
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	homeDir := system.GetUserHome(hostOS)
	wdbRootDirectory := fmt.Sprintf(WDB_ROOT_PATH_FORMAT, homeDir)

	wdbConfigDirectory := fmt.Sprintf(WDB_CONFIG_DIR_PATH_FORMAT, wdbRootDirectory)
	wdbConfigFilePath := fmt.Sprintf(WDB_CONFIG_FILE_PATH_FORMAT, wdbConfigDirectory)

	configFileBytes, err := fs.ReadFile(wdbConfigFilePath)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	var configMap ConfigMap
	err = json.Unmarshal(configFileBytes, &configMap)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	c := &Config{
		AdminID:               configMap.getValue(ADMIN_ID),
		AdminPassword:         configMap.getValue(ADMIN_PASSWORD),
		Port:                  configMap.getValue(PORT),
		PersistantStoragePath: configMap.getValue(PERSISTANT_STORAGE_PATH),
	}

	if c.PersistantStoragePath == "" {
		c.PersistantStoragePath = fmt.Sprintf(WDB_PERSISTANT_STORAGE_DIR_PATH_FORMAT, wdbRootDirectory)
	}

	return c, nil

}

func (c ConfigMap) getValue(key string) string {
	val, exists := c[key]
	if !exists {
		envVal := os.Getenv(key)
		if envVal == "" {
			defaultVal, defaultvalExists := defaultValues[key]
			if !defaultvalExists {
				return ""
			} else {
				return defaultVal
			}
		}
		return envVal
	}
	return val
}

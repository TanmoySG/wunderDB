package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

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
	RootDirectoryPath     string `json:"ROOT_DIR_PATH"`
	AdminID               string `json:"ADMIN_ID"`
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

	override, _ := strconv.ParseBool(os.Getenv(OVERRIDE_CONFIG))

	homeDir := system.GetUserHome(hostOS)
	wdbRootDirectory := fmt.Sprintf(WDB_ROOT_PATH_FORMAT, homeDir)

	wdbConfigDirectory := fmt.Sprintf(WDB_CONFIG_DIR_PATH_FORMAT, wdbRootDirectory)
	err = fs.CreateDirectory(wdbConfigDirectory)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	wdbConfigFilePath := fmt.Sprintf(WDB_CONFIG_FILE_PATH_FORMAT, wdbConfigDirectory)
	err = handleFile(wdbConfigFilePath)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

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
		RootDirectoryPath:     wdbRootDirectory,
		AdminID:               configMap.getValue(ADMIN_ID, override),
		AdminPassword:         configMap.getValue(ADMIN_PASSWORD, override),
		Port:                  configMap.getValue(PORT, override),
		PersistantStoragePath: configMap.getValue(PERSISTANT_STORAGE_PATH, override),
	}

	if c.PersistantStoragePath == "" {
		c.PersistantStoragePath = fmt.Sprintf(WDB_PERSISTANT_STORAGE_DIR_PATH_FORMAT, wdbRootDirectory)
	}

	err = writeConfigFile(wdbConfigFilePath, *c)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return c, nil

}

func (c ConfigMap) getValue(key string, override bool) string {
	val, exists := c[key]

	// invoke handleConfig if config doesn't exist
	if !exists {
		return getConfigurationValue(key)
	}

	// invoke handleConfig if override is enabled
	if override {
		return getConfigurationValue(key)
	}

	// invoke handleConfig if value is empty
	if val == "" {
		return getConfigurationValue(key)
	}

	return val
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

func writeConfigFile(configFilePath string, config Config) error {
	configMapOverrideBytes, err := json.Marshal(config)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	err = fs.WriteToFile(configFilePath, configMapOverrideBytes)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	return nil
}

func getConfigurationValue(key string) string {
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

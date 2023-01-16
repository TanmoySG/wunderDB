package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/TanmoySG/wunderDB/pkg/fs"
	"github.com/TanmoySG/wunderDB/pkg/utils/system"
)

var 	WFS_DIRECTORIES = []string{"users", "namespaces", "databases", "roles"}


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

// func LoadConfig() (*Config, error) {
// 	hostOS, err := system.GetHostOS()
// 	if err != nil {
// 		return nil, err
// 	}

// 	homeDir := system.GetUserHome(hostOS)

// 	wdbRootDirectory := fmt.Sprintf(WDB_ROOT_PATH_FORMAT, homeDir)

// 	configFilePath, err := configPreChecks(wdbRootDirectory)
// 	if err != nil {
// 		return nil, err
// 	}

// 	c, err := LoadConfigFromFile(configFilePath)
// 	if err != nil {
// 		return nil, err
// 	}

// 	fmt.Print(c)

// 	return c, nil
// }

// func LoadConfigFromEnv() (*Config, error) {
// 	config := Config{}
// 	if err := env.Parse(&config); err != nil {
// 		return nil, err
// 	}
// 	return &config, nil
// }

// func LoadConfigFromFile(configFilePath string) (*Config, error) {
// 	config := Config{}

// 	configBytes, err := fs.ReadFile(configFilePath)
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = json.Unmarshal(configBytes, &config)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &config, nil
// }

// func configPreChecks(rootDir string) (string, error) {

// 	wdbConfigDirectory := fmt.Sprintf(WDB_CONFIG_DIR_PATH_FORMAT, rootDir)
// 	err := fs.CreateDirectory(wdbConfigDirectory)
// 	if err != nil {
// 		return "", err
// 	}

// 	wdbConfigFilePath := fmt.Sprintf(WDB_CONFIG_FILE_PATH_FORMAT, wdbConfigDirectory)

// 	configFileExists := fs.CheckFileExists(wdbConfigFilePath)
// 	if !configFileExists {
// 		err := fs.CreateFile(wdbConfigFilePath)
// 		if err != nil {
// 			return "", err
// 		}

// 		c, err := LoadConfigFromEnv()
// 		if err != nil {
// 			return "", err
// 		}

// 		if c.PersistantStoragePath == "" {
// 			c.PersistantStoragePath = fmt.Sprintf(WDB_PERSISTANT_STORAGE_DIR_PATH_FORMAT, rootDir)
// 		}

// 		if c.Port == "" {
// 			c.Port = DEFAULT_PORT
// 		}

// 		configMap, err := json.Marshal(c)
// 		if err != nil {
// 			return "", err
// 		}

// 		err = fs.WriteToFile(wdbConfigFilePath, configMap)
// 		if err != nil {
// 			return "", err
// 		}
// 	}

// 	return wdbConfigFilePath, nil
// }

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

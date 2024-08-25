package config

import "github.com/TanmoySG/wunderDB/internal/users/admin"

const (
	ADMIN_ID                = "ADMIN_ID"
	ADMIN_PASSWORD          = "ADMIN_PASSWORD"
	PORT                    = "PORT"
	PERSISTANT_STORAGE_PATH = "PERSISTANT_STORAGE_PATH"
	OVERRIDE_CONFIG         = "OVERRIDE_CONFIG"
	ROOT_DIR_PATH           = "ROOT_DIR_PATH"
	RUN_MODE                = "RUN_MODE"
)

type RUN_MODE_TYPE string

const (
	RUN_MODE_MAINTENANCE RUN_MODE_TYPE = "RUN_MODE_MAINTENANCE"
	RUN_MODE_UPGRADE     RUN_MODE_TYPE = "RUN_MODE_UPGRADE"
	RUN_MODE_NORMAL      RUN_MODE_TYPE = "RUN_MODE_NORMAL"
)

var defaultValues = map[string]string{
	PORT:           "8086",
	ADMIN_ID:       admin.DEFAULT_ADMIN_USERID,
	ADMIN_PASSWORD: admin.DEFAULT_ADMIN_PASSWORD,
	RUN_MODE:       string(RUN_MODE_NORMAL),
}

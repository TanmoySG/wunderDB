package config

import "github.com/TanmoySG/wunderDB/internal/users/admin"

const (
	ADMIN_ID                = "ADMIN_ID"
	ADMIN_PASSWORD          = "ADMIN_PASSWORD"
	PORT                    = "PORT"
	PERSISTANT_STORAGE_PATH = "PERSISTANT_STORAGE_PATH"
	OVERRIDE_CONFIG         = "OVERRIDE_CONFIG"
)

var defaultValues = map[string]string{
	PORT:           "8086",
	ADMIN_ID:       admin.DEFAULT_ADMIN_USERID,
	ADMIN_PASSWORD: admin.DEFAULT_ADMIN_PASSWORD,
}

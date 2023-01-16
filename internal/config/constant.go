package config

const (
	ADMIN_ID                = "ADMIN_ID"
	ADMIN_PASSWORD          = "ADMIN_PASSWORD"
	PORT                    = "PORT"
	PERSISTANT_STORAGE_PATH = "PERSISTANT_STORAGE_PATH"
)

var defaultValues = map[string]string{
	PORT: "8086",
}

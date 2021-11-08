package config

import "main/utils"

var (
	SERVICE_PORT  = utils.Env("SERVICE_PORT", "8889")
	DUMP_DURATION = utils.Env("RECORD_FREQ", "60")
	LOG_PATH      = utils.Env("LOG_PATH", "/app/goys/log")
)

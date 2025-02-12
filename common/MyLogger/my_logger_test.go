package MyLogger

import (
	"testing"

	"govenv/pkg/common/Config"
)

func Test_1_check_only_logger_basic(t *testing.T) {
	config := Config.NewConfig("basic logger test", "./config.cfg", "", true)
	logger := NewMyLogger(config.Name, config, false, false, false, false, false, StrDEBUG, "./logs", "")
	logger.Debug(StrDEBUG)
	logger.Info(StrINFO)
	logger.Warning(StrWARNING)
	logger.Error(StrERROR)
	logger.Critical(StrCRITICAL)
}

func Test_2_check_notif_sending(t *testing.T) {
	config := Config.NewConfig("basic logger test with notifs", "./config.cfg", "", true)
	logger := NewMyLogger(config.Name, config, false, false, false, false, false, StrDEBUG, "./logs", "")
	logger.Debug(StrDEBUG)
	logger.Info(StrINFO)
	logger.Warning(StrWARNING)
	logger.Error(StrERROR)
	logger.Critical(StrCRITICAL)
}

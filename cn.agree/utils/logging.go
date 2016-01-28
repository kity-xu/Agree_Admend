package utils

import (
	log "github.com/cihub/seelog"
)

func init() {
}

//1:代表debug
//2:代表Info
//3:代表WARNING
//4:代表ERROR
func InitializeLogging(level int) {
	ChangeLogEnv(level)
}

func Trace(format string, values ...interface{}) {
	log.Tracef(format, values...)
}

func Debug(format string, values ...interface{}) {
	log.Debugf(format, values...)
}

func Info(format string, values ...interface{}) {
	log.Infof(format, values...)
}

func Warn(format string, values ...interface{}) {
	log.Warnf(format, values...)
}

func Error(format string, values ...interface{}) {
	log.Errorf(format, values...)
}

func Critical(format string, values ...interface{}) {
	log.Criticalf(format, values...)
}

func Flush() {
	log.Flush()
}

func ChangeLogEnv(kind int) {
	//开发环境
	if kind == 1 {
		logger, err := log.LoggerFromConfigAsFile("config/logconfig_debug.xml")
		if err != nil {
			log.Errorf("create log error,error is %s ", err.Error())
		}
		log.UseLogger(logger)

	} else if kind == 2 { //生产环境
		logger, err := log.LoggerFromConfigAsFile("config/logconfig_production.xml")
		if err != nil {
			log.Errorf("create log error,error is %s ", err.Error())
		}
		log.UseLogger(logger)

	} else {
		Error("can't recognize log environment")
	}

}

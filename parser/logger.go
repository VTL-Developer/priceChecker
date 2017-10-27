package main

import (
	"fmt"
	"log"
	"os"
)

type logLevel int

const (
	fatalLevel   logLevel = iota
	errorLevel   logLevel = iota
	warningLevel logLevel = iota
	infoLevel    logLevel = iota
	debugLevel   logLevel = iota
)

var logger = log.New(os.Stdout, "", log.LUTC|log.LstdFlags)

func logFatal(unformatted string, data ...interface{}) {
	if configuration.LogLevel >= fatalLevel {
		logger.Print(formattedLog(fatalLevel, formatLogInput(unformatted, data...)))
	}
}

func logError(unformatted string, data ...interface{}) {
	if configuration.LogLevel >= errorLevel {
		logger.Print(formattedLog(errorLevel, formatLogInput(unformatted, data...)))
	}
}

func logWarning(unformatted string, data ...interface{}) {
	if configuration.LogLevel >= warningLevel {
		logger.Print(formattedLog(warningLevel, formatLogInput(unformatted, data...)))
	}
}

func logInfo(unformatted string, data ...interface{}) {
	if configuration.LogLevel >= infoLevel {
		logger.Print(formattedLog(infoLevel, formatLogInput(unformatted, data...)))
	}
}

func logDebug(unformatted string, data ...interface{}) {
	if configuration.LogLevel >= debugLevel {
		logger.Print(formattedLog(debugLevel, formatLogInput(unformatted, data...)))
	}
}

func formatLogInput(unformatted string, data ...interface{}) string {
	return fmt.Sprintf(unformatted, data...)
}

func formattedLog(level logLevel, log string) string {
	return fmt.Sprintf("[%v] [%v]: %v\n", convertLogLevelToString(level), os.Getpid(), log)
}

func convertLogLevelToString(level logLevel) string {
	switch level {
	case fatalLevel:
		return "FATAL"
	case errorLevel:
		return "ERROR"
	case warningLevel:
		return "WARNING"
	case infoLevel:
		return "INFO"
	case debugLevel:
		return "DEBUG"
	default:
		return ""
	}
}

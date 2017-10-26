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

func logFatal(log string) {
	if configuration.LogLevel >= fatalLevel {
		logger.Print(formattedLog(fatalLevel, log))
	}
}

func logError(log string) {
	if configuration.LogLevel >= errorLevel {
		logger.Print(formattedLog(errorLevel, log))
	}
}

func logWarning(log string) {
	if configuration.LogLevel >= warningLevel {
		logger.Print(formattedLog(warningLevel, log))
	}
}

func logInfo(log string) {
	if configuration.LogLevel >= infoLevel {
		logger.Print(formattedLog(infoLevel, log))
	}
}

func logDebug(log string) {
	if configuration.LogLevel >= debugLevel {
		logger.Print(formattedLog(debugLevel, log))
	}
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

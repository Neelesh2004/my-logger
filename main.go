package mylogger

import "log"

func LogInfo(message string) {
	log.Println("INFO - ",message)
}

func LogWarn(message string) {
	log.Println("WARN - ",message)
}

func LogError(message string) {
	log.Println("ERROR - ",message)
}

func LogDebug(message string) {
	log.Println("DEBUG - ",message)
}
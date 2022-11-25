package nutil

import "testing"

func allPrintLog(log *BasicLogger) {
	log.Debug("Log Level Debug")
	log.Info("Log Level Info")
	log.Warn("Log Level Warn")
	log.Error("Log Level Error")
}

func TestLogger(t *testing.T) {

	t.Run("DefaultLog Test", func(t *testing.T) {
		log := CreateBasicLogger()
		allPrintLog(log)
	})

	t.Run("Change Log Level Test", func(t *testing.T) {
		log := CreateBasicLogger()
		log.SetLogLevel(LogPrintError)
		allPrintLog(log)
	})
}

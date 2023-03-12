package StrategyPattern

import "testing"

func TestNewLogManager(t *testing.T) {
	fileLogging := &FileLogging{}
	logmanager := NewLogManager(fileLogging)

	logmanager.Info()
	logmanager.Error()

	dbLogging := &DBLogging{}
	logmanager = NewLogManager(dbLogging)

	logmanager.Info()
	logmanager.Error()
}

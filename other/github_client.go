package main

import (
	"log/slog"
	"net/http"
	"sync"
)

var lock = &sync.Mutex{}

var _singleInstance *http.Client

func getGithubClient() *http.Client {
	if _singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if _singleInstance == nil {
			slog.Debug("Creating single instance now.")
			_singleInstance = &http.Client{}
		} else {
			slog.Debug("Single instance already created.")
		}
	} else {
		slog.Debug("Single instance already created.")
	}
	return _singleInstance
}

package main

type SessionMessage struct {
	SessionID int
	Attempts  int
}

var sessionQueue = make(chan SessionMessage, 10)
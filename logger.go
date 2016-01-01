package main

import (
	"fmt"
	"log"
	"os"
	"io/ioutil"
)

var errorlog *os.File
var logger *log.Logger
var logfile = "debug.log"

/*
	This file enables the usage of the `logger`.
	Based on this answer from SO: http://stackoverflow.com/a/32826322
*/

func init() {
	errorlog, err := os.OpenFile(logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("error opening logging file: %v", err)
		os.Exit(1)
	}
	logger = log.New(errorlog, "debug: ", log.Lshortfile|log.LstdFlags)

	// DEV flag helps to stop overhead from logging in production
	if DEVELOPMENT == false {
		logger.SetFlags(-1)
		logger.SetOutput(ioutil.Discard)
	}
}
/*
*	Copyright © 2017 Joe Xu <joe.0x01@gmail.com>
*	This work is free. You can redistribute it and/or modify it under the
*	terms of the Do What The Fuck You Want To Public License, Version 2,
*	as published by Sam Hocevar. See the COPYING file for more details.
*
 */

//Package logger provides logger in different level
package logger

import (
	"io"
	"log"
	"os"
)

var (
	//InfoPrefix info-logger prefix
	InfoPrefix = "[INFO]"
	//WarnPrefix warn-logger prefix
	WarnPrefix = "[WARN]"
	//ErrPrefix error-logger prefix
	ErrPrefix = "[ERROR]"
	//DebugPrefix debug-logger prefix
	DebugPrefix = "[DEBUG]"
)

var (
	//Info do info-level logging
	Info *log.Logger
	//Warn do warning-level logging
	Warn *log.Logger
	//Error do error-level logging
	Error *log.Logger
	//Debug do debug-level logging
	Debug *log.Logger
)

func init() {
	Info = log.New(os.Stdout, InfoPrefix, log.LstdFlags)
	Warn = log.New(os.Stderr, WarnPrefix, log.LstdFlags)
	Error = log.New(os.Stderr, ErrPrefix, log.LstdFlags)
	Debug = log.New(os.Stdout, DebugPrefix, log.LstdFlags)
}

//Init func init loggers
func Init(infoHandle, warnHandle, errHandle, debugHandle io.Writer) {
	Info = log.New(infoHandle, InfoPrefix, log.LstdFlags)
	Warn = log.New(warnHandle, WarnPrefix, log.LstdFlags)
	Error = log.New(errHandle, ErrPrefix, log.LstdFlags)
	Debug = log.New(debugHandle, DebugPrefix, log.LstdFlags)
}

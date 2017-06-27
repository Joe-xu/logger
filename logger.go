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

const (
	//InfoPrefix info-logger prefix
	InfoPrefix = "[INFO]"
	//WarnPrefix warn-logger prefix
	WarnPrefix = "[WARN]"
	//ErrPrefix error-logger prefix
	ErrPrefix = "[ERROR]"
)

var (
	//Info do info-level logging
	Info *log.Logger
	//Warn do warning-level logging
	Warn *log.Logger
	//Error do error-level logging
	Error *log.Logger
)

func init() {
	Info = log.New(os.Stdout, InfoPrefix, log.LstdFlags)
	Warn = log.New(os.Stderr, WarnPrefix, log.LstdFlags)
	Error = log.New(os.Stderr, ErrPrefix, log.LstdFlags)
}

//Init func init loggers
func Init(infoHandle, warnHandle, errHandle io.Writer) {
	Info = log.New(infoHandle, InfoPrefix, log.LstdFlags)
	Warn = log.New(warnHandle, WarnPrefix, log.LstdFlags)
	Error = log.New(errHandle, ErrPrefix, log.LstdFlags)
}

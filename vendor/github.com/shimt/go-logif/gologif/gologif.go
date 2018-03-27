// Copyright 2016 Shinichi MOTOKI. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package gologif is logif wrapper for log package
package gologif

import (
	"io"
	"log"
	"sync/atomic"

	"github.com/shimt/go-logif"
)

var levelString []string
var levelStringWithSpace []string

func init() {
	var l logif.LogLevel
	levelString = make([]string, logif.MAXLEVEL+1)
	levelStringWithSpace = make([]string, logif.MAXLEVEL+1)
	for l = logif.MINLEVEL; l <= logif.MAXLEVEL; l++ {
		levelString[l] = "[" + l.String() + "]"
		levelStringWithSpace[l] = "[" + l.String() + "] "
	}
}

type logger struct {
	entity      *log.Logger
	outputLevel int32
}

// SetFlags sets the output flags for the logger.
func (l *logger) SetFlags(flag int) {
	l.entity.SetFlags(flag)
}

// Flags returns the output flags for the logger.
func (l *logger) Flags() int {
	return l.entity.Flags()
}

// SetPrefix sets the output prefix for the logger.
func (l *logger) SetPrefix(prefix string) {
	l.entity.SetPrefix(prefix)
}

// Prefix returns the output prefix for the logger.
func (l *logger) Prefix() string {
	return l.entity.Prefix()
}

// Output writes the output for a logging event.
// The string s contains the text to print after the prefix specified
// by the flags of the Logger. A newline is appended if the last character
// of s is not already a newline.
// Calldepth is used to recover the PC and is provided for generality,
// although at the moment on all pre-defined paths it will be 2
func (l *logger) Output(calldepth int, s string) error {
	return l.entity.Output(calldepth, s)
}

// SetOutput sets the output destination for the logger.
func (l *logger) SetOutput(w io.Writer) {
	l.entity.SetOutput(w)
}

// Print calls l.Output to print to the logger. Arguments are handled in the manner of fmt.Print.
func (l *logger) Print(v ...interface{}) {
	l.entity.Print(v...)
}

// Printf calls l.Output to print to the logger. Arguments are handled in the manner of fmt.Printf.
func (l *logger) Printf(format string, v ...interface{}) {
	l.entity.Printf(format, v...)
}

// Println calls l.Output to print to the logger. Arguments are handled in the manner of fmt.Println.
func (l *logger) Println(v ...interface{}) {
	l.entity.Println(v...)
}

// Fatal write message(level=ERROR) to the logger followed by a call to os.Exit(1).
// Arguments are handled in the manner of fmt.Print.
func (l *logger) Fatal(v ...interface{}) {
	l.entity.Fatal(v...)
}

// Fatalf write message(level=ERROR) to the logger followed by a call to os.Exit(1).
// Arguments are handled in the manner of fmt.Printf.
func (l *logger) Fatalf(format string, v ...interface{}) {
	l.entity.Fatalf(format, v...)
}

// Fatalln iwrite message(level=ERROR) to the logger followed by a call to os.Exit(1).
// Arguments are handled in the manner of fmt.Println.
func (l *logger) Fatalln(v ...interface{}) {
	l.entity.Fatalln(v...)
}

// Panic write message(level=PANIC) to the logger followed by a call to panic().
// Arguments are handled in the manner of fmt.Print.
func (l *logger) Panic(v ...interface{}) {
	l.entity.Panic(v...)
}

// Panicf write message(level=PANIC) to the logger followed by a call to panic().
// Arguments are handled in the manner of fmt.Printf.
func (l *logger) Panicf(format string, v ...interface{}) {
	l.entity.Panicf(format, v...)
}

// Panic write message(level=PANIC) to the logger followed by a call to panic().
// Arguments are handled in the manner of fmt.Print.
func (l *logger) Panicln(v ...interface{}) {
	l.entity.Panicln(v...)
}

func (l *logger) p(level logif.LogLevel, v []interface{}) {
	v2 := make([]interface{}, 0, len(v)+1)
	v2 = append(v2, levelStringWithSpace[level])
	v2 = append(v2, v...)
	l.Print(v2...)
}

func (l *logger) pf(level logif.LogLevel, format string, v []interface{}) {
	l.Printf(levelStringWithSpace[level]+format, v...)
}

func (l *logger) pln(level logif.LogLevel, v []interface{}) {
	v2 := make([]interface{}, 0, len(v)+1)
	v2 = append(v2, levelString[level])
	v2 = append(v2, v...)
	l.Println(v2...)
}

// Debug write message(level=DEBUG) to the logger.
// Arguments are handled in the manner of fmt.Print.
func (l *logger) Debug(v ...interface{}) {
	if l.OutputLevel() > logif.DEBUG {
		return
	}

	l.p(logif.DEBUG, v)
}

// Debugf write message(level=DEBUG) to the logger.
// Arguments are handled in the manner of fmt.Printf.
func (l *logger) Debugf(format string, v ...interface{}) {
	if l.OutputLevel() > logif.DEBUG {
		return
	}

	l.pf(logif.DEBUG, format, v)
}

// Debugln write message(level=DEBUG) to the logger.
// Arguments are handled in the manner of fmt.Println.
func (l *logger) Debugln(v ...interface{}) {
	if l.OutputLevel() > logif.DEBUG {
		return
	}

	l.pln(logif.DEBUG, v)
}

// Info write message(level=INFO) to the logger.
// Arguments are handled in the manner of fmt.Print.
func (l *logger) Info(v ...interface{}) {
	if l.OutputLevel() > logif.INFO {
		return
	}

	l.p(logif.INFO, v)
}

// Infof write message(level=INFO) to the logger.
// Arguments are handled in the manner of fmt.Printf.
func (l *logger) Infof(format string, v ...interface{}) {
	if l.OutputLevel() > logif.INFO {
		return
	}

	l.pf(logif.INFO, format, v)
}

// Infoln write message(level=INFO) to the logger.
// Arguments are handled in the manner of fmt.Println.
func (l *logger) Infoln(v ...interface{}) {
	if l.OutputLevel() > logif.INFO {
		return
	}

	l.pln(logif.INFO, v)
}

// Warn write message(level=WARN) to the logger.
// Arguments are handled in the manner of fmt.Print.
func (l *logger) Warn(v ...interface{}) {
	if l.OutputLevel() > logif.WARN {
		return
	}

	l.p(logif.WARN, v)
}

// Warnf write message(level=WARN) to the logger.
// Arguments are handled in the manner of fmt.Printf.
func (l *logger) Warnf(format string, v ...interface{}) {
	if l.OutputLevel() > logif.WARN {
		return
	}

	l.pf(logif.WARN, format, v)
}

// Warnln write message(level=WARN) to the logger.
// Arguments are handled in the manner of fmt.Println.
func (l *logger) Warnln(v ...interface{}) {
	if l.OutputLevel() > logif.WARN {
		return
	}

	l.pln(logif.WARN, v)
}

// Error write message(level=ERROR) to the logger.
// Arguments are handled in the manner of fmt.Print.
func (l *logger) Error(v ...interface{}) {
	if l.OutputLevel() > logif.ERROR {
		return
	}

	l.p(logif.ERROR, v)
}

// Errorf write message(level=ERROR) to the logger.
// Arguments are handled in the manner of fmt.Printf.
func (l *logger) Errorf(format string, v ...interface{}) {
	if l.OutputLevel() > logif.ERROR {
		return
	}

	l.pf(logif.ERROR, format, v)
}

// Errorln write message(level=ERROR) to the logger.
// Arguments are handled in the manner of fmt.Println.
func (l *logger) Errorln(v ...interface{}) {
	if l.OutputLevel() > logif.ERROR {
		return
	}

	l.pln(logif.ERROR, v)
}

// SetOutputLevel set output level
func (l *logger) SetOutputLevel(level logif.LogLevel) {
	atomic.StoreInt32(&l.outputLevel, int32(level))
}

// OutputLevel set output level
func (l *logger) OutputLevel() logif.LogLevel {
	return logif.LogLevel(atomic.LoadInt32(&l.outputLevel))
}

// ToGoLogger cast to logif.GoLogger
func (l *logger) ToGoLogger() logif.GoLogger {
	return l.entity
}

// New create new logger instance.
func New(out io.Writer, prefix string, flag int) (l logif.Logger) {
	return &logger{
		entity:      log.New(out, prefix, flag),
		outputLevel: int32(logif.WARN),
	}
}

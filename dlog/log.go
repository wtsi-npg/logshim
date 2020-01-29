/*
 * Copyright (C) 2019, 2020. Genome Research Ltd. All rights reserved.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License,
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 * @file log.go
 * @author Keith James <kdj@sanger.ac.uk>
 */

package dlog

import (
	"fmt"
	"io"
	"log"
	"strings"
	"time"

	"github.com/kjsanger/logshim"
)

type levelName string

const (
	errorLevel levelName = "ERROR"
	warnLevel  levelName = "WARN"
	// noticeLevel levelName = "NOTICE"
	infoLevel  levelName = "INFO"
	debugLevel levelName = "DEBUG"
)

func translateLevel(level logshim.Level) (levelName, error) {
	var (
		lvn levelName
		err error
	)

	switch level {
	case logshim.ErrorLevel:
		lvn = errorLevel
	case logshim.WarnLevel:
		lvn = warnLevel
	case logshim.NoticeLevel:
		fallthrough
	case logshim.InfoLevel:
		lvn = infoLevel
	case logshim.DebugLevel:
		lvn = debugLevel
	default:
		lvn = warnLevel
		err = fmt.Errorf("invalid log level %d, defaulting to "+
			"WARN level", level)
	}

	return lvn, err
}

type StdLogger struct {
	name  string
	Level logshim.Level
	*log.Logger
}

func New(writer io.Writer, level logshim.Level) *StdLogger {
	lg := log.New(writer, "", log.LstdFlags|log.Lshortfile)

	_, err := translateLevel(level)
	if err != nil {
		log.Print(errorLevel, "log configuration error", err)
		level = logshim.WarnLevel
	}

	return &StdLogger{"StdLog", level, lg}
}

func (log *StdLogger) Name() string {
	return log.name
}

func (log *StdLogger) Err(err error) logshim.Message {
	effectiveLevel := logshim.InfoLevel
	if err != nil {
		effectiveLevel = logshim.ErrorLevel
	}

	active := log.Level >= effectiveLevel
	msg := &stdMessage{active, effectiveLevel, &strings.Builder{}}
	msg.Err(err)
	return msg
}

func (log *StdLogger) Error() logshim.Message {
	active := log.Level >= logshim.ErrorLevel
	msg := &stdMessage{active, logshim.ErrorLevel, &strings.Builder{}}
	return msg
}

func (log *StdLogger) Warn() logshim.Message {
	active := log.Level >= logshim.WarnLevel
	msg := &stdMessage{active, logshim.WarnLevel, &strings.Builder{}}
	return msg
}

func (log *StdLogger) Notice() logshim.Message {
	active := log.Level >= logshim.NoticeLevel
	msg := &stdMessage{active, logshim.InfoLevel, &strings.Builder{}}
	return msg
}

func (log *StdLogger) Info() logshim.Message {
	active := log.Level >= logshim.InfoLevel
	msg := &stdMessage{active, logshim.InfoLevel, &strings.Builder{}}
	return msg
}

func (log *StdLogger) Debug() logshim.Message {
	active := log.Level >= logshim.DebugLevel
	msg := &stdMessage{active, logshim.DebugLevel, &strings.Builder{}}
	return msg
}

type stdMessage struct {
	active  bool
	level   logshim.Level
	builder *strings.Builder
}

func (msg *stdMessage) Err(err error) logshim.Message {
	if msg.active {
		msg.builder.WriteString(fmt.Sprintf(" error: %v", err))
	}
	return msg
}

func (msg *stdMessage) Bool(key string, val bool) logshim.Message {
	if msg.active {
		msg.builder.WriteString(fmt.Sprintf(" %s: %v", key, val))
	}
	return msg
}

func (msg *stdMessage) Dur(key string, val time.Duration) logshim.Message {
	if msg.active {
		msg.builder.WriteString(fmt.Sprintf(" %s: %v", key, val))
	}
	return msg
}

func (msg *stdMessage) Int(key string, val int) logshim.Message {
	if msg.active {
		msg.builder.WriteString(fmt.Sprintf(" %s: %d", key, val))
	}
	return msg
}

func (msg *stdMessage) Int64(key string, val int64) logshim.Message {
	if msg.active {
		msg.builder.WriteString(fmt.Sprintf(" %s: %d", key, val))
	}
	return msg
}

func (msg *stdMessage) Uint64(key string, val uint64) logshim.Message {
	if msg.active {
		msg.builder.WriteString(fmt.Sprintf(" %s: %d", key, val))
	}
	return msg
}

func (msg *stdMessage) Str(key string, val string) logshim.Message {
	if msg.active {
		msg.builder.WriteString(fmt.Sprintf(" %s: %s", key, val))
	}
	return msg
}

func (msg *stdMessage) Time(key string, val time.Time) logshim.Message {
	if msg.active {
		msg.builder.WriteString(fmt.Sprintf(" %s: %v", key, val))
	}
	return msg
}

func (msg *stdMessage) Msg(val string) {
	if msg.active {
		lvn, err := translateLevel(msg.level)
		if err != nil {
			// This should never happen because the Logger constructor corrects
			// invalid level values.
			log.Print(errorLevel, "log configuration error", err)
		}

		msg.builder.WriteString(" ")
		msg.builder.WriteString(val)
		log.Print(lvn, msg.builder.String())
		// Once this method is called, deactivate for all future calls
		msg.active = false
	}
}

func (msg *stdMessage) Msgf(format string, a ...interface{}) {
	msg.Msg(fmt.Sprintf(format, a...))
}

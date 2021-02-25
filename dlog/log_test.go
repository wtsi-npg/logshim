/*
 * Copyright (C) 2019, 2021. Genome Research Ltd. All rights reserved.
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
 * @file log_test.go
 * @author Keith James <kdj@sanger.ac.uk>
 */

package dlog

import (
	"os"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"github.com/wtsi-npg/logshim"
)

func TestNew(t *testing.T) {
	for _, level := range logshim.Levels() {
		logger := New(os.Stderr, level)
		assert.NotNil(t, logger, "StdLogger level %d was nil", level)
	}
}

func TestStdLogger_Err(t *testing.T) {
	log := New(os.Stderr, logshim.ErrorLevel)
	assert.NotNil(t, log.Err(errors.New("test error")),
		"Err() should return a Message")
}

func TestStdLogger_Error(t *testing.T) {
	log := New(os.Stderr, logshim.ErrorLevel)
	assert.NotNil(t, log.Error(), "Error() should return a Message")
}

func TestStdLogger_Warn(t *testing.T) {
	log := New(os.Stderr, logshim.ErrorLevel)
	assert.NotNil(t, log.Warn(), "Warn() should return a Message")
}

func TestStdLogger_Notice(t *testing.T) {
	log := New(os.Stderr, logshim.ErrorLevel)
	assert.NotNil(t, log.Notice(), "Notice() should return a Message")
}

func TestStdLogger_Info(t *testing.T) {
	log := New(os.Stderr, logshim.ErrorLevel)
	assert.NotNil(t, log.Info(), "Info() should return a Message")
}

func TestStdLogger_Debug(t *testing.T) {
	log := New(os.Stderr, logshim.ErrorLevel)
	assert.NotNil(t, log.Debug(), "Debug() should return a Message")
}

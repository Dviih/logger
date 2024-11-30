/*
 *     A colored handler for slog.
 *     Copyright (C) 2024  Dviih
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU Affero General Public License as published
 *     by the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU Affero General Public License for more details.
 *
 *     You should have received a copy of the GNU Affero General Public License
 *     along with this program.  If not, see <https://www.gnu.org/licenses/>.
 *
 */
package logger

import (
	"log/slog"
	"os"
	"testing"
	"unsafe"
)

const message = "a test message"

var test = slog.New(New(os.Stdout, Time, slog.LevelDebug))

func TestLoggerDebug(t *testing.T) {
	test.Debug(message)
}

func TestLoggerInfo(t *testing.T) {
	test.Info(message)
}

func TestLoggerWarn(t *testing.T) {
	test.Warn(message)
}

func TestLoggerError(t *testing.T) {
	test.Error(message)
}

func TestLoggerUintptr(t *testing.T) {
	i := 13
	test.Info("a uintptr", slog.Any("pointer", unsafe.Pointer(&i)))
}

func TestLoggerComplex(t *testing.T) {
	test.Debug("a complex number", slog.Any("number", complex(13, 256)))
}

func TestLoggerSlice(t *testing.T) {
	test.Warn("a slice of uint64", slog.Any("numbers", []uint64{128, 256, 512, 1024}))
}

func TestLoggerChan(t *testing.T) {
	test.Info("a channel", slog.Any("chan", make(chan int, 64)))
}


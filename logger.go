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

const (
	Time = "02 Jan 06 15:04:05 MST"
)

var (
	m sync.Mutex

	debug  = []byte("\033[1;34mDEBUG \033[0m")
	info   = []byte("\033[1;36mINFO \033[0m")
	warn   = []byte("\033[1;31mWARN \033[0m")
	_error = []byte("\033[1;91mERROR \033[0m")

)

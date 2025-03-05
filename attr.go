/*
 *     A colored handler for slog.
 *     Copyright (C) 2025  Dviih
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
	"reflect"
	"strconv"
	"time"
)

func (logger *Logger) attrs(s string, attribute slog.Attr) error {
	switch attribute.Value.Kind() {
	case slog.KindAny, slog.KindBool, slog.KindFloat64, slog.KindInt64, slog.KindString, slog.KindUint64:
		if err := logger.attr2(s, attribute); err != nil {
			return err
		}

		return nil
	case slog.KindDuration:
		if err := logger.attr1(s, attribute); err != nil {
			return err
		}

		if err := logger.write([]byte("\u001B[33m" + attribute.Value.Duration().String())); err != nil {
			return err
		}

		return nil
	case slog.KindTime:
		if err := logger.attr1(s, attribute); err != nil {
			return err
		}

		if err := logger.write([]byte("\u001B[33m" + attribute.Value.Time().Format(time.RFC822))); err != nil {
			return err
		}

		return nil
	case slog.KindGroup:
		for _, attr := range attribute.Value.Group() {
			if err := logger.attrs(prefix(s, attribute.Key), attr); err != nil {
				return err
			}
		}

		return nil
	case slog.KindLogValuer:
		if err := logger.attrs(s, slog.Any(attribute.Key, attribute.Value.LogValuer().LogValue())); err != nil {
			return err
		}

		return nil
	}

	return nil
}

func (logger *Logger) attr1(s string, attribute slog.Attr) error {
	if err := logger.write(" \u001B[0;32m"); err != nil {
		return err
	}

	if logger.group != nil {
		if err := logger.write(append(logger.group, '.')); err != nil {
			return err
		}
	}

	if s != "" {
		if err := logger.write(s + "."); err != nil {
			return err
		}
	}

	if err := logger.write(attribute.Key + "\u001B[0m"); err != nil {
		return err
	}

	if err := logger.write("->"); err != nil {
		return err
	}

	return nil
}

func (logger *Logger) attr2(s string, attribute slog.Attr) error {
	value := reflect.ValueOf(attribute.Value.Any())
	for value.Kind() == reflect.Pointer || (value.Kind() == reflect.Interface && !value.IsZero()) {
		value = value.Elem()
	}

	attribute.Value = slog.AnyValue(value.Interface())

	switch value.Kind() {
	case reflect.Invalid:
		return nil
	case reflect.Bool:
		if err := logger.attr1(s, attribute); err != nil {
			return err
		}

		if attribute.Value.Bool() {
			if err := logger.write(_true); err != nil {
				return err
			}

			return nil
		}

		if err := logger.write(_false); err != nil {
			return err
		}

		return nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64, reflect.String:
		if err := logger.attr1(s, attribute); err != nil {
			return err
		}

		if err := logger.write("\u001B[0;93m" + attribute.Value.String() + "\u001B[0m"); err != nil {
			return err
		}

		return nil
	case reflect.Uintptr, reflect.UnsafePointer:
		if err := logger.attrs(s, slog.String(attribute.Key, "0x"+strconv.FormatUint(uint64(value.Pointer()), 16))); err != nil {
			return err
		}

		return nil
	case reflect.Complex64, reflect.Complex128:
		c := value.Complex()

		if err := logger.attrs(prefix(s, attribute.Key), slog.Float64("real", real(c))); err != nil {
			return err
		}

		if err := logger.attrs(prefix(s, attribute.Key), slog.Float64("imag", imag(c))); err != nil {
			return err
		}

		return nil
	case reflect.Array, reflect.Slice:
		for i := 0; i < value.Len(); i++ {
			if err := logger.attrs(s, slog.Any(attribute.Key+"["+strconv.Itoa(i)+"]", value.Index(i).Interface())); err != nil {
				return err
			}
		}
	case reflect.Chan:
		if err := logger.attrs(s, slog.String("kind", chanDir(value.Type().ChanDir()))); err != nil {
			return err
		}

		if err := logger.attrs(s, slog.String("type", value.Type().Elem().String())); err != nil {
			return err
		}

		if err := logger.attrs(s, slog.String("pointer", "0x"+strconv.FormatUint(uint64(value.Pointer()), 16))); err != nil {
			return err
		}

		if err := logger.attrs(s, slog.Int("len", value.Len())); err != nil {
			return err
		}

		if err := logger.attrs(s, slog.Int("cap", value.Cap())); err != nil {
			return err
		}

		return nil
	case reflect.Func:
		if err := logger.attrs(s, slog.String("signature", value.Type().String())); err != nil {
			return err
		}

		if err := logger.attrs(s, slog.String("pointer", "0x"+strconv.FormatUint(uint64(value.Pointer()), 16))); err != nil {
			return err
		}

		return nil
	case reflect.Interface, reflect.Pointer:
		// both cases should be impossible.
		return nil
	case reflect.Map:
		m := value.MapRange()

		for m.Next() {
			if err := logger.attrs(s, slog.Any(attribute.Key+"["+slog.AnyValue(m.Key().Interface()).String()+"]", m.Value().Interface())); err != nil {
				return err
			}
		}
	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			ft := value.Type().Field(i)

			if !ft.IsExported() {
				continue
			}

			if err := logger.attrs(prefix(s, attribute.Key), slog.Any(ft.Name, value.Field(i).Interface())); err != nil {
				return err
			}
		}
	}

	return nil
}

func chanDir(dir reflect.ChanDir) string {
	switch dir {
	case reflect.BothDir:
		return "bidirectional"
	case reflect.RecvDir:
		return "receiver"
	case reflect.SendDir:
		return "sender"
	}

	return ""
}

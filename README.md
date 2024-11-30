# Logger
### A colored handler for Go's `slog`, this project was also created to help with [golinux](https://github.com/Dviih/golinux)

---

## Constants and Variables
- **Time**: default time formating for logger based on `time.RFC822`.
- **debug**: text for `Debug` message.
- **info**: text for `Info` message.
- **warn**: text for `Warn` message.
- **_error**: text for `Error` message.
- **_true**: text for true boolean.
- **_false**: text for false boolean.

## Functions
- **prefix**: group and attributes names for children of them.

---

## Logger
### A structure that contains options required for printing and implements `slog.Handler` interface

### Properties
- **writer**: stores an `io.Writer` interface.
- **time**: time formatting.
- **level**: minimum level required for printing.
- **attributes**: array of additional attributes.
- **group**: slog group name.

### Methods
- **Enabled**: returns if requested level can be printed compared to `logger.level`.
#### Made for Gophers by @Dviih

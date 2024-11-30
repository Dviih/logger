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
- **chanDir**: returns a string with the direction of the channel.
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
- **Handle**: handles the printing of a log based on provided `slog.Record`.
- **WithAttrs**: pushes attributes into `logger.attributes`
- **WithGroup**: groups a slog but rather add a prefix to it.
- **write**: writes string, byte(s) to writer and handles errors and size.
- **attrs**: main attribute writer, for slog kinds it is as it is but for Go's standard it uses `reflect`.
- **attr1**: used to parse name and arrows (used to not repeat code at all).
- **attr2**: real writing bound with `reflect`.

---

## Example

```go
package main

import (
	"github.com/Dviih/logger"
	"log/slog"
	"os"
)

func main() {
	s := slog.New(logger.New(os.Stdout, logger.Time, slog.LevelInfo))
	
	s.Info("Hello, World!")
}
```

---

#### Made for Gophers by @Dviih

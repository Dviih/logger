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

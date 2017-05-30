package database

import "os"

func ensureDBPath(path string) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		_ = os.Mkdir(path, 0777)
	}
}

package file

import "os"

func Exists(path string) bool {
	_, err := os.Stat(path)

	return err == nil
}

func Delete(path string) error {
	return os.Remove(path)
}

func DeleteAll(path string) error {
	return os.RemoveAll(path)
}

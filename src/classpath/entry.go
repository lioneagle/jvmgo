package classpath

import (
	"os"
)

const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	ReadClass(classname string) ([]byte, Entry, error)
	String() string
}

func newEntry(path string) Entry {
	return Entry{}
}

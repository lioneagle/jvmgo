package classpath

import (
	"errors"
	//"fmt"
	"strings"
)

type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	compositeEntry := []Entry{}
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}

	return compositeEntry
}

func (self CompositeEntry) ReadClass(classname string) ([]byte, Entry, error) {
	//fmt.Printf("CompositeEntry.ReadClass: %s, class = %s\n", self.String(), classname)
	for _, entry := range self {
		//fmt.Printf("entry: %s\n", entry.String())
		data, from, err := entry.ReadClass(classname)
		if err == nil {
			return data, from, nil
		}
	}
	return nil, nil, errors.New("class not found: " + classname)
}

func (self CompositeEntry) String() string {
	strs := make([]string, len(self))
	for i, entry := range self {
		strs[i] = entry.String()
	}
	return strings.Join(strs, pathListSeparator)
}

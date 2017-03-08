package classpath

import (
	"archive/zip"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}

func (self *ZipEntry) ReadClass(classname string) ([]byte, Entry, error) {
	//fmt.Printf("ZipEntry.ReadClass: absPath=%s\nclass = %s\n", self.String(), classname)
	reader, err := zip.OpenReader(self.absPath)
	if err != nil {
		return nil, nil, err
	}
	defer reader.Close()

	for _, file := range reader.File {
		//fmt.Println("file.Name =", file.Name, "classname =", classname)
		if file.Name == classname {
			rc, err := file.Open()
			if err != nil {
				return nil, nil, err
			}

			fmt.Println("find at ", self.String())

			data, err := ioutil.ReadAll(rc)
			rc.Close()
			if err != nil {
				return nil, nil, err
			}
			//fmt.Println("data =", data)
			return data, self, nil
		}

	}
	return nil, nil, errors.New("Can not find: " + classname)
}

func (self *ZipEntry) String() string {
	return self.absPath
}

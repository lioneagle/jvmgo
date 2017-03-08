package classpath

import (
	//"fmt"
	"os"
	"path/filepath"
)

type ClassPath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

func Parse(jreOption, cpOption string) *ClassPath {
	cp := &ClassPath{}
	cp.parseBootAndExtClassPath(jreOption)
	cp.parseUserClassPath(cpOption)
	return cp
}

func (self *ClassPath) parseBootAndExtClassPath(jreOption string) {
	jreDir := getJreDir(jreOption)
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClasspath = newWildcardEntry(jreLibPath)
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.extClasspath = newWildcardEntry(jreExtPath)
}

func (self *ClassPath) parseUserClassPath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}

	self.userClasspath = newEntry(cpOption)

	//fmt.Println("self.userClasspath =", self.userClasspath)
}

func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	if exists("./jre") {
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("Can not find jre folder!")
}

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func (self *ClassPath) ReadClass(classname string) ([]byte, Entry, error) {
	classname = classname + ".class"

	if data, entry, err := self.bootClasspath.ReadClass(classname); err == nil {
		//fmt.Println("find in self.bootClasspath:", self.bootClasspath)
		return data, entry, nil
	}
	if data, entry, err := self.extClasspath.ReadClass(classname); err == nil {
		//fmt.Println("find in self.extClasspath:", self.extClasspath)
		return data, entry, nil
	}

	data, entry, err := self.userClasspath.ReadClass(classname)
	if err == nil {
		//fmt.Println("find in self.userClasspath:", self.userClasspath)
		return data, entry, nil
	}
	return nil, nil, err
}

func (self *ClassPath) String() string {
	return self.userClasspath.String()
}

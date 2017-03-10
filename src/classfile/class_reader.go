package classfile

import (
	"encoding/binary"
)

type ClassReader struct {
	data []byte
}

func (self *ClassReader) ReadUint8() uint8 {
	val := self.data[0]
	self.data = self.data[1:]
	return val
}

func (self *ClassReader) ReadUint16() uint16 {
	val := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:]
	return val
}

func (self *ClassReader) ReadUint32() uint32 {
	val := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:]
	return val
}

func (self *ClassReader) ReadUint64() uint64 {
	val := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:]
	return val
}

func (self *ClassReader) ReadUint16s() []uint16 {
	n := self.ReadUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = self.ReadUint16()
	}
	return s
}

func (self *ClassReader) ReadBytes(n uint32) []byte {
	bytes := self.data[:n]
	self.data = self.data[n:]
	return bytes
}
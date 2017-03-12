package conversions

import (
	"instructions/base"
	"rtda"
)

type D2F struct{ base.NoOperandsInstruction }
type D2I struct{ base.NoOperandsInstruction }
type D2L struct{ base.NoOperandsInstruction }

// Convert double to float
func (self *D2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	x := stack.PopDouble()
	y := float32(x)
	stack.PushFloat(y)
}

// Convert double to int
func (self *D2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	x := stack.PopDouble()
	y := int32(x)
	stack.PushInt(y)
}

// Convert double to long
func (self *D2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	x := stack.PopDouble()
	y := int64(x)
	stack.PushLong(y)
}

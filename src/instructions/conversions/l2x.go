package conversions

import (
	"instructions/base"
	"rtda"
)

type L2D struct{ base.NoOperandsInstruction }
type L2F struct{ base.NoOperandsInstruction }
type L2I struct{ base.NoOperandsInstruction }

// Convert long to double
func (self *L2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	x := stack.PopLong()
	y := float64(x)
	stack.PushDouble(y)
}

// Convert long to float
func (self *L2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	x := stack.PopLong()
	y := float32(x)
	stack.PushFloat(y)
}

// Convert long to int
func (self *L2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	x := stack.PopLong()
	y := int32(x)
	stack.PushInt(y)
}

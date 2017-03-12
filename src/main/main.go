package main

import (
	"classfile"
	"classpath"
	"cmds"
	"fmt"
	//"instructions/base"
	"instructions/constants"
	"rtda"
	"strings"
)

func main() {

	cmd := cmds.ParseCmd()
	if cmd.Version {
		fmt.Println("version 0.0.1")
	} else if cmd.Help || cmd.Class == "" {
		cmds.PrintUsage()
	} else {
		startJVM(cmd)
	}

	//fmt.Println("cmd =", cmd)

}

func startJVM(cmd *cmds.Cmd) {

	frame := rtda.NewFrame(100, 100)
	testLocalVars(frame.LocalVars())
	testOperandStack(frame.OperandStack())

	x1 := &constants.ACONST_NULL{}
	x1.Execute(frame)

	classPath := classpath.Parse(cmd.Xjre, cmd.ClassPath)
	fmt.Printf("classpath: %s, class: %v, args: %v\n", classPath, cmd.Class, cmd.Args)
	className := strings.Replace(cmd.Class, ".", "/", -1)
	classFile := loadClass(className, classPath)
	fmt.Println(cmd.Class)
	printClassInfo(classFile)
	/*
		classData, _, err := classPath.ReadClass(className)
		if err != nil {
			fmt.Printf("Can not find or load main class %s\n", cmd.Class)
			return
		}
		fmt.Printf("class data: %v\n", classData)*/
}

func loadClass(className string, cp *classpath.ClassPath) *classfile.ClassFile {
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}
	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}
	return cf
}

func printClassInfo(cf *classfile.ClassFile) {
	fmt.Printf("version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("constants count: %v\n", len(cf.ConstantPool()))
	fmt.Printf("access flags: 0x%x\n", cf.AccessFlags())
	fmt.Printf("this class: %v\n", cf.ClassName())
	fmt.Printf("super class: %v\n", cf.SuperClassName())
	fmt.Printf("interfaces: %v\n", cf.InterfaceNames())
	fmt.Printf("fields count: %v\n", len(cf.Fields()))
	for _, f := range cf.Fields() {
		fmt.Printf(" %s\n", f.Name())
	}
	fmt.Printf("methods count: %v\n", len(cf.Methods()))
	for _, m := range cf.Methods() {
		fmt.Printf(" %s\n", m.Name())
	}
}

func testLocalVars(vars rtda.LocalVars) {
	vars.SetInt(0, 100)
	vars.SetInt(1, -100)
	vars.SetLong(2, 2997924580)
	vars.SetLong(4, -2997924580)
	vars.SetFloat(6, 3.1415926)
	vars.SetDouble(7, 2.71828182845)
	vars.SetRef(9, nil)
	println(vars.GetInt(0))
	println(vars.GetInt(1))
	println(vars.GetLong(2))
	println(vars.GetLong(4))
	println(vars.GetFloat(6))
	println(vars.GetDouble(7))
	println(vars.GetRef(9))
}

func testOperandStack(ops *rtda.OperandStack) {
	ops.PushInt(100)
	ops.PushInt(-100)
	ops.PushLong(2997924580)
	ops.PushLong(-2997924580)
	ops.PushFloat(3.1415926)
	ops.PushDouble(2.71828182845)
	ops.PushRef(nil)
	println(ops.PopRef())
	println(ops.PopDouble())
	println(ops.PopFloat())
	println(ops.PopLong())
	println(ops.PopLong())
	println(ops.PopInt())
	println(ops.PopInt())
}

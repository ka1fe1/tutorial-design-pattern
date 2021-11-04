package template_method

import (
	"fmt"
	"testing"
)

// 模板接口
// 定义了一系列模板方法
type ITemplate interface {
	MethodA()
	MethodB()
	MethodC()
}

// 父类：用于继承，复用代码
type Base struct {
}

func (i *Base) MethodA()  {
	fmt.Println("exec method a")
}

func (i *Base) MethodB()  {
	fmt.Println("exec method b")
}

// 子类-A：重写了 C 方法
type A struct {
	Base
}

func (i *A) MethodC()  {
	fmt.Println("a exec method c")
}

// 子类-B：重写了 A, C 方法
type B struct {
	Base
}

func (i *B) MethodA()  {
	fmt.Println("b exec method a")
}

func (i *B) MethodC()  {
	fmt.Println("b exec method c")
}

func TestExample(t *testing.T)  {
	a := A{}
	a.MethodA()
	a.MethodB()
	a.MethodC()

	fmt.Println("---")

	b := B{}
	b.MethodA()
	b.MethodB()
	b.MethodC()
}


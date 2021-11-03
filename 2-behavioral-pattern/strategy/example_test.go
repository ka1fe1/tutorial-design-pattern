package strategy

import (
	"fmt"
	"testing"
)

// 1
// 策略接口
type Strategy interface {
	Exec() string
}

// 2
// 具体策略实现-A
type A struct {
}

func (i *A) Exec() string {
	return "object is a"
}

// 2
// 具体策略实现-B
type B struct {
}

func (i *B) Exec() string {
	return "object is b"
}

// 3
// 上下文定义了客户端关注的接口；
// 上下文维护了某个具体策略接口的引用，并不知晓具体的策略类，通过策略接口与该对象进行交互
type StrategyContext struct {
	s Strategy
}

func (i *StrategyContext) SetStrategy(p Strategy) {
	i.s = p
}

func (i *StrategyContext) DoSomething() string {
	if i.s != nil {
		return i.s.Exec()
	}

	return ""
}

func TestExample(t *testing.T) {
	var a float64
	var c StrategyContext

	a = 0.3
	if a >= 0.5 {
		c.SetStrategy(&A{})
	} else {
		c.SetStrategy(&B{})
	}

	fmt.Println(c.DoSomething())
}

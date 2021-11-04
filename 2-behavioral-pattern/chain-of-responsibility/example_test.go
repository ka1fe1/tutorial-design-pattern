package chain_of_responsibility

import (
	"math/rand"
	"time"
	"fmt"
	"testing"
)

// 处理者：声明了具体处理者所有的通用接口
type IHandler interface {
	Handle(params *string)
	SetNext(next IHandler) IHandler

	Name() string
}

// 父类：用于继承，封装公共代码
type BaseHandler struct {
	next IHandler
}

func (i *BaseHandler) SetNext(h IHandler) IHandler {
	i.next = h
	return i.next
}

func (i *BaseHandler) Handle(currentHandler IHandler, params *string) {
	if i.next != nil {
		fmt.Printf("%s passed on the request to %s, params: %s\n", currentHandler.Name(), i.next.Name(), *params)
		i.next.Handle(params)
	} else {
		fmt.Printf("%s is last handler, final params: %s\n", currentHandler.Name(), *params)
	}
}

// 处理者-A
// 包含处理请求的实际代码，每个处理者接收到请求后，都必须决定是否处理以及是否沿着链传递请求
type A struct {
	BaseHandler
}

func (i *A) Handle(params *string) {
	rand.Seed(time.Now().Unix())
	f := rand.Float64()
	currentHandler := i.Name()
	if f >= 0.6 {
		fmt.Printf("%s break handler chain\n", currentHandler)
		i.next = nil
	} else {
		np := fmt.Sprintf("%s-%s", *params, currentHandler)
		params = &np
	}
	i.BaseHandler.Handle(i, params)
}

func (i *A) Name() string {
	return "a"
}

// 处理者-B
type B struct {
	BaseHandler
}

func (i *B) Handle(params *string) {
	rand.Seed(time.Now().Unix())
	f := rand.Float64()
	currentHandler := i.Name()
	if f >= 0.7 {
		fmt.Printf("%s break handler chain\n", currentHandler)
		i.next = nil
	} else {
		np := fmt.Sprintf("%s-%s", *params, currentHandler)
		params = &np
	}
	i.BaseHandler.Handle(i, params)
}

func (i *B) Name() string {
	return "b"
}

// 处理者-C
type C struct {
	BaseHandler
}

func (i *C) Handle(params *string) {
	rand.Seed(time.Now().Unix())
	f := rand.Float64()
	currentHandler := i.Name()
	if f >= 0.8 {
		fmt.Printf("%s break handler chain\n", currentHandler)
		i.next = nil
	} else {
		np := fmt.Sprintf("%s-%s", *params, currentHandler)
		params = &np
	}
	i.BaseHandler.Handle(i, params)
}

func (i *C) Name() string {
	return "c"
}

func TestExample(t *testing.T) {
	a := &A{}
	b := &B{}
	c := &C{}

	a.SetNext(c).SetNext(b)
	params := "test params"
	a.Handle(&params)

	time.Sleep(2 * time.Second)
	fmt.Println("---")

	a.SetNext(nil)
	b.SetNext(nil)
	c.SetNext(nil)

	b.SetNext(a).SetNext(c)

	a.Handle(&params)
}

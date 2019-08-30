package main

import (
	"fmt"
	"sync"
)

func main() {
	dic := newDIContainer()
	a := dic.serviceA()
	fmt.Println("from A:")
	a.serviceB.initB()
	c := dic.serviceC()
	fmt.Println("from C:")
	c.serviceB.initB()
}

type diContainter struct {
	env string

	serviceA func() *serviceA
	serviceB func() *serviceB
	serviceC func() *serviceC
}

func newDIContainer() *diContainter {
	dic := &diContainter{
		env: "testing",
	}

	dic.serviceA = newServiceADIProvider(dic)
	dic.serviceB = newServiceBDIProvider(dic)
	dic.serviceC = newServiceCDIProvider(dic)
	return dic

}

type serviceA struct {
	x        int
	y        int
	serviceB interface {
		initB()
	}
}

func newServiceA(dic *diContainter) *serviceA {
	fmt.Println("** Init serviceA **")
	return &serviceA{
		x:        1,
		y:        2,
		serviceB: dic.serviceB(),
	}
}

func newServiceADIProvider(dic *diContainter) func() *serviceA {
	fmt.Println("** Init DI Provider A **")
	var a *serviceA
	var mu sync.Mutex
	return func() *serviceA {
		mu.Lock()
		defer mu.Unlock()
		if a == nil {
			a = newServiceA(dic)
		}
		return a
	}
}

type serviceC struct {
	x        int
	y        int
	serviceB interface {
		initB()
	}
}

func newServiceC(dic *diContainter) *serviceC {
	fmt.Println("** Init serviceC **")
	return &serviceC{
		x:        1,
		y:        2,
		serviceB: dic.serviceB(),
	}
}

func newServiceCDIProvider(dic *diContainter) func() *serviceC {
	fmt.Println("** Init DI Provider C **")
	var c *serviceC
	var mu sync.Mutex
	return func() *serviceC {
		mu.Lock()
		defer mu.Unlock()
		if c == nil {
			c = newServiceC(dic)
		}
		return c
	}
}

type serviceB struct {
	env string
}

func (b *serviceB) initB() {
	fmt.Println("hey there:", b.env)
}

func newServiceB(dic *diContainter) *serviceB {
	fmt.Println("** Init serviceB **")
	return &serviceB{
		env: dic.env,
	}
}

func newServiceBDIProvider(dic *diContainter) func() *serviceB {
	fmt.Println("** Init DI Provider B **")
	var b *serviceB
	var mu sync.Mutex
	return func() *serviceB {
		mu.Lock()
		defer mu.Unlock()
		if b == nil {
			b = newServiceB(dic)
		} else {
			fmt.Println("** Using already init service B **")
		}
		return b
	}
}

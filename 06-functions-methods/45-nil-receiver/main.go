package main

import (
	"errors"
	"log"
	"strings"
)

type MultiError struct {
	errs []string
}

func (m *MultiError) Add(err error) {
	m.errs = append(m.errs, err.Error())
}

func (m *MultiError) Error() string {
	return strings.Join(m.errs, ";")
}

type Customer struct {
	Age  int
	Name string
}

func (c Customer) Validate1() error {
	// nil で初期化されるので、nilのポインタを返す
	// 検査結果によってnilポインタか構造体へのポインタを返す
	var m *MultiError

	if c.Age < 0 {
		m = &MultiError{}
		m.Add(errors.New("age is negative"))
	}
	if c.Name == "" {
		if m == nil {
			m = &MultiError{}
		}
		m.Add(errors.New("name is nil"))
	}

	return m
}

func (c Customer) Validate2() error {
	var m *MultiError

	if c.Age < 0 {
		m = &MultiError{}
		m.Add(errors.New("age is negative"))
	}
	if c.Name == "" {
		if m == nil {
			m = &MultiError{}
		}
		m.Add(errors.New("name is nil"))
	}

	// nilでない時だけ、mを返す
	if m != nil {
		return m
	}
	return nil
}

func main() {
	customer := Customer{Age: 33, Name: "John"}
	//この場合はnilぽインタなので、trueになってしまう
	//if err := customer.Validate1(); err != nil {
	// nilポインタの時だけnilになる
	if err := customer.Validate2(); err != nil {
		log.Fatalf("customer is invalid: %v", err)
	}
}

package main

import "testing"

func TestSum(t *testing.T) {
	a := 2
	b := 32
	expected := int(a+b)
	actual, err := sum(a,b)

	if err!= nil{
		t.Errorf("Ошибки нет")
	}

	if expected != actual {
		t.Errorf("результат не целочисленный")
	}

	
}
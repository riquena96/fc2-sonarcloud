package main

import "testing"

func TestSum(t *testing.T) {

	result := sum(2, 3)

	if result != 5 {
		t.Error(("O resultado é diferente de 5"))
	}

}

func TestSub(t *testing.T) {

	result := sub(5, 2)

	if result != 3 {
		t.Error(("O resultado é diferente de 3"))
	}

}

func TestMult(t *testing.T) {

	result := mult(5, 2)

	if result != 10 {
		t.Error(("O resultado é diferente de 10"))
	}

}

func TestDiv(t *testing.T) {

	result := div(10, 2)

	if result != 5 {
		t.Error(("O resultado é diferente de 5"))
	}

}

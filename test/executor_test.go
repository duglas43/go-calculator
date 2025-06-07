package executor_test

import (
	"testing"

	"github.com/duglas43/calculator/internal/executor"
)

func TestSimpleAdd(t *testing.T) {
	e := executor.NewExecutor()
	input := []executor.Instruction{
		{Type: "calc", Op: "+", Var: "x", Left: 2, Right: 3},
		{Type: "print", Var: "x"},
	}
	res, err := e.Execute(input)
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 1 || res[0].Value != 5 {
		t.Fatalf("Expected x = 5, got %+v", res)
	}
}
